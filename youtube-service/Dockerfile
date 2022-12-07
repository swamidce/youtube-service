FROM golang:1.19-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o build/server cmd/*

FROM alpine

COPY --from=build /app/build/server /server

EXPOSE 3500

ENTRYPOINT [ "/server" ]