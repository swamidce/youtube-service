package add_key

import (
	"context"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/youtube-service/internal/configs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var collection *mongo.Collection

func SetCollection(client *mongo.Client) {
	collection = client.Database("cmd").Collection("add_key")
}

// Make a single read call to YouTube API to check if the key is valid
func IsKeyValid(key string) bool {
	ctx := context.Background()
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(key))
	if err != nil {
		log.Errorf("IsKeyValid: Error creating YouTube service: %v", err)
		return false
	}

	call := youtubeService.Channels.List([]string{"id"}).ForUsername("Youtube")
	_, err = call.Do()
	if err != nil {
		log.Errorf("IsKeyValid: Error making YouTube API call: %v", err)
		return false
	}
	return true
}

// Inserts a new key into the database
func InsertKey(key string) error {
	_, err := collection.InsertOne(context.TODO(), bson.M{"key": key, "isExpired": false, "lastUpdated": time.Now()})
	if err != nil {
		log.Errorf("InsertKey: Error inserting key: %v", err)
		return err
	}
	log.Info("InsertKey: Inserted new API key")
	return nil
}

func GetValidKey() (string, error) {
	if configs.GetValidApiKey() != "" {
		return configs.GetValidApiKey(), nil
	}

	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{"isExpired": false})
	if err != nil {
		log.Errorf("GetValidKey: Error finding valid key: %v", err)
		return "", errors.New("error finding valid key")
	}
	defer cursor.Close(ctx)

	var key bson.M

	for cursor.Next(ctx) {
		err := cursor.Decode(&key)
		if err != nil {
			log.Errorf("GetValidKey: Error decoding key: %v", err)
			continue
		}
		return key["key"].(string), nil
	}

	return "", errors.New("no valid key found")
}

// Expires an API key
func SetKeyToExpired(key string) error {
	ctx := context.Background()
	_, err := collection.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": bson.M{"isExpired": true, "lastUpdated": time.Now()}})
	if err != nil {
		log.Errorf("SetKeyToExpired: Error setting key to expired: %v", err)
		return err
	}
	return nil
}

// Sets an API key to not expired
func SetKeyToNotExpired(key string) error {
	ctx := context.Background()
	_, err := collection.UpdateOne(ctx, bson.M{"key": key}, bson.M{"$set": bson.M{"isExpired": false, "lastUpdated": time.Now()}})
	if err != nil {
		log.Errorf("SetKeyToNotExpired: Error setting key to not expired: %v", err)
		return err
	}
	return nil
}

// Sets keys whose last updated time is more than 24 hours to not expired
func UpdateExpirationOfExpiredKeys() {
	ctx := context.Background()
	cursor, err := collection.Find(ctx, bson.M{"isExpired": true})
	if err != nil {
		log.Errorf("UpdateExpirationOfExpiredKeys: Error finding expired keys: %v", err)
		return
	}
	defer cursor.Close(ctx)

	var key bson.M

	for cursor.Next(ctx) {
		err := cursor.Decode(&key)
		if err != nil {
			log.Errorf("UpdateExpirationOfExpiredKeys: Error decoding key: %v", err)
			continue
		}
		// Check if last updated is greater than 24 hours and if key is valid
		lastUpdated := key["lastUpdated"].(time.Time)
		if time.Since(lastUpdated) > 24*time.Hour && IsKeyValid(key["key"].(string)) {
			// Set key to not expired
			SetKeyToNotExpired(key["key"].(string))
		}
	}
}
