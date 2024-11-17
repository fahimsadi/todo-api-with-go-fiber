package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var redisClient *redis.Client

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisAddr := redisHost + ":" + redisPort

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis!")
}

func Set(key string, value string) error {
	err := redisClient.Set(ctx, key, value, 0).Err()
	return err
}

func Get(key string) (string, error) {
	val, err := redisClient.Get(ctx, key).Result()
	return val, err
}

func Delete(key string) error {
	err := redisClient.Del(ctx, key).Err()
	return err
}

func Increment(key string) (int, error) {
	val, err := redisClient.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

func Close() {
	if redisClient != nil {
		redisClient.Close()
	}
}

func Keys(pattern string) ([]string, error) {
	keys, err := redisClient.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}
	return keys, nil
}
