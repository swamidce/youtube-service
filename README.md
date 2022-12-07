# Youtube Service
This service fetch & search video and add Api Key so that if quota is exhausted on one then it automatically uses the next available key.

## Pre-Requisites

1. Install [Golang](https://golang.org/doc/install).
2. Install and run [MongoDB](https://docs.mongodb.com/manual/installation/).

## Rest APIs

### Get Video

```
curl -X GET -H "Content-Type: application/json" http://localhost:3500/get_video?page=2
```

### Search Video

```
curl -X GET -H "Content-Type: application/json" http://localhost:3500/search_video?query=ind+live&page=2
```

### Add API Key

```
curl -X POST -H "Content-Type: application/json" http://localhost:3500/add_key?key=<API_KEY>
```
