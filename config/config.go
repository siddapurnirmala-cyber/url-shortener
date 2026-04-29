package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Collection {

	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo not reachable:", err)
	}

	log.Println("✅ MongoDB connected")

	return client.Database("url_shortener").Collection("urls")
}

func ConnectRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("⚠️ Redis not available, continuing without cache")
		return nil
	}

	log.Println("✅ Redis connected")
	return rdb
}
