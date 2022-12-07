package entities

import "time"

type Video struct {
	Id          string `json:"_id,omitempty" bson:"_id,omitempty"`
	UniqueId    string `json:"uniqueId" bson:"uniqueId"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	PublishedAt string `json:"publishedAt" bson:"publishedAt"`
}

type ApiKey struct {
	Id          string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Key         string    `json:"key" bson:"key"`
	IsExpired   bool      `json:"isExpired" bson:"isExpired"`
	LastUpdated time.Time `json:"lastUpdated" bson:"lastUpdated"`
}
