package main

import (
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/youtube-service/internal/router"
	"github.com/youtube-service/internal/db/mongo"
	"github.com/youtube-service/internal/handlers"
	"github.com/youtube-service/internal/models-services"
	"github.com/youtube-service/internal/configs"
	"github.com/youtube-service/pkg/logger"
	log "github.com/sirupsen/logrus"
	"github.com/go-playground/validator/v10"
)

func main() {
	logger.InitLogger()

	// Load environment variables from configs file
	if err := godotenv.Load(); err != nil {
		log.Fatal("main: failed to load environment variables")
	}

	config.InitConfig()
	db.ConnectionDb()

	// Start a goroutine to fetch videos from youtube periodically
	go func() {
		ticker := time.NewTicker(time.Duration(config.GetFetchLatestVideosSeconds()) * time.Second)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				err := get_video-search_video.FetchNewVideosAndUpdateDb()
				if err != nil {
					log.Errorf("main: error fetching new videos and updating db: %v", err)
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}

	}()

	// Start a goroutine to update expiation of API keys in the database periodically
	go func() {
		ticker := time.NewTicker(time.Duration(config.GetUpdateApiKeysExpirationMinutes()) * time.Minute)
		quit := make(chan struct{})
		for {
			select {
			case <-ticker.C:
				apikeys.UpdateExpirationOfExpiredKeys()
			case <-quit:
				ticker.Stop()
				return
			}
		}

	}()

	app := fiber.New()
	router.SetRoutes(app)

	log.Infof("main: Starting server on port %v", config.GetPort())
	app.Listen(config.GetPort())
}
