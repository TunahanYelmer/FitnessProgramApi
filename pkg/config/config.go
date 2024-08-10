package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
    OpenAIAPIKey string
}

type MongoConfig struct {
    URI string
}

// LoadConfig loads environment variables from the .env file and returns a Config struct
func LoadConfig() *Config {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    apiKey := os.Getenv("OPENAI_API_KEY")
    if apiKey == "" {
        log.Fatal("OPENAI_API_KEY environment variable is not set")
    }

    return &Config{
        OpenAIAPIKey: apiKey,
    }
}
func LoadMongoConfig() string {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}
	return uri
}
// GetConnection establishes a connection to MongoDB using the provided URI
func GetConnection() (*mongo.Client, error) {
    clientOptions := options.Client().ApplyURI(LoadMongoConfig())
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        return nil, err
    }
    return client, nil
}