package configs

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/youtube-service/pkg/utils"
)

type Config struct {
	Port                       string
	Etag                           string
	MaxVideosFetched               int64
	PerPageLimit                   int64
	FetchLatestVideosSeconds       int64
	UpdateApiKeysExpirationMinutes int64
	Query                          string
	MongoDbURI                     string
	ValidApiKey                    string
}

const (
	DEFAULT_MAX_TOKENS                         = 5
	DEFAULT_PER_PAGE_LIMIT                     = 5
	DEFAULT_FETCH_LATEST_VIDEOS_SECONDS        = 10
	DEFAULT_UPDATE_API_KEYS_EXPIRATION_MINUTES = 120
)

var configs Config

// Initializes configs with environment variables or default values if not specified
func InitConfig() {
	flag.StringVar(&configs.Port, "port", os.Getenv("PORT"), "Port where server is running")
	if configs.Port == "" {
		log.Fatalf("Config: Environment variable PORT not found. Please refer to README to find how to set it.")
	}

	flag.StringVar(&configs.MongoDbURI, "mongodburi", os.Getenv("MONGODB_URI"), "MongoDB URI for connection")
	if configs.MongoDbURI == "" {
		log.Fatalf("Config: Environment variable MONGODB_URI not found. Please refer to README to find how to set it.")
	}

	flag.StringVar(&configs.Query, "query", os.Getenv("QUERY"), "Predefined search query")
	if configs.Query == "" {
		log.Fatalf("Config: Environment variable QUERY not found. Please refer to README to find how to set it.")
	}

	flag.Int64Var(&configs.MaxVideosFetched, "maxvideosfetched", utils.GetEnvInt("MAX_VIDEOS_FETCHED", DEFAULT_MAX_TOKENS), "Max videos that can be fetched in a single API call")
	if configs.MaxVideosFetched > 50 || configs.MaxVideosFetched < 1 {
		log.Infof("Config: Environment variable MAX_VIDEOS_FETCHED should be between 1 and 50. Please refer to README. Setting it to default value: %d", DEFAULT_MAX_TOKENS)
		configs.MaxVideosFetched = DEFAULT_MAX_TOKENS
	}

	flag.Int64Var(&configs.PerPageLimit, "perpagelimit", utils.GetEnvInt("PER_PAGE_LIMIT", DEFAULT_PER_PAGE_LIMIT), "Number of videos to be displayed per page")
	if configs.PerPageLimit < 1 {
		log.Infof("Config: Environment variable PER_PAGE_LIMIT should be greater than 0. Please refer to README. Setting it to default value: %d", DEFAULT_PER_PAGE_LIMIT)
		configs.PerPageLimit = DEFAULT_PER_PAGE_LIMIT
	}

	flag.Int64Var(&configs.FetchLatestVideosSeconds, "fetchlatestvideosseconds", utils.GetEnvInt("FETCH_LATEST_VIDEOS_SECONDS", DEFAULT_FETCH_LATEST_VIDEOS_SECONDS), "Number of seconds after which latest videos are fetched from youtube and database is updated")
	if configs.FetchLatestVideosSeconds < 1 {
		log.Infof("Config: Environment variable FETCH_LATEST_VIDEOS_SECONDS should be greater than 0. Please refer to README. Setting it to default value: %d", DEFAULT_FETCH_LATEST_VIDEOS_SECONDS)
		configs.FetchLatestVideosSeconds = DEFAULT_FETCH_LATEST_VIDEOS_SECONDS
	}

	flag.Int64Var(&configs.UpdateApiKeysExpirationMinutes, "updateapikeysexpirationminutes", utils.GetEnvInt("UPDATE_API_KEYS_EXPIRATION_MINUTES", DEFAULT_UPDATE_API_KEYS_EXPIRATION_MINUTES), "Number of minutes after which expired api keys whose quota has exceeded are checked for validity and updated")
	if configs.UpdateApiKeysExpirationMinutes < 1 {
		log.Infof("Config: Environment variable UPDATE_API_KEYS_EXPIRATION_MINUTES should be greater than 0. Please refer to README. Setting it to default value: %d", DEFAULT_UPDATE_API_KEYS_EXPIRATION_MINUTES)
		configs.UpdateApiKeysExpirationMinutes = DEFAULT_UPDATE_API_KEYS_EXPIRATION_MINUTES
	}

	flag.Parse()

	// Etag will be empty for the first API call
	configs.Etag = ""
}

func GetPort() string {
	return configs.Port
}

func GetQuery() string {
	return configs.Query
}

func GetMaxVideosFetched() int64 {
	return configs.MaxVideosFetched
}

func GetPerPageLimit() int64 {
	return configs.PerPageLimit
}

func GetFetchLatestVideosSeconds() int64 {
	return configs.FetchLatestVideosSeconds
}

func GetUpdateApiKeysExpirationMinutes() int64 {
	return configs.UpdateApiKeysExpirationMinutes
}

func GetMongoDbURI() string {
	return configs.MongoDbURI
}

func GetEtag() string {
	return configs.Etag
}

func GetValidApiKey() string {
	return configs.ValidApiKey
}

func SetValidApiKey(apiKey string) {
	configs.ValidApiKey = apiKey
}

func SetEtag(etag string) {
	configs.Etag = etag
}
