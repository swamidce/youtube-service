package db

import (
	"context"
	"log"
	"time"

	"github.com/youtube-service/models-services/add_key"
	"github.com/youtube-service/models-services/get_video-search_video"
	"github.com/youtube-service/internal/configs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Establishes connection to the mongo database, sets collections and creates indexes
func ConnectionDb() {
	client := ConnectToMongoDb()
	get_video-search_video.SetCollection(client)
	get_video-search_video.CreateTitleAndDescriptionIndex()
	apikeys.SetCollection(client)
}

func ConnectToMongoDb() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI((config.GetMongoDbURI())).SetServerAPIOptions(serverAPIOptions))
	if err != nil {
		log.Fatalf("ConnectMongo: Error connecting to mongo db: %v", err)
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatalf("ConnectMongo: Error pinging mongo db: %v", err)
	}
	return client
}
