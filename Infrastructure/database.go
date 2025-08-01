package Infrastructure

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Database represents the MongoDB database connection
type Database struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewDatabase creates a new database connection
func NewDatabase(uri, databaseName string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected to MongoDB")

	return &Database{
		Client:   client,
		Database: client.Database(databaseName),
	}, nil
}

// Close closes the database connection
func (db *Database) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}

// GetCollection returns a MongoDB collection
func (db *Database) GetCollection(collectionName string) *mongo.Collection {
	return db.Database.Collection(collectionName)
}

// Collection names
const (
	CollectionUsers            = "users"
	CollectionBlogPosts        = "blog_posts"
	CollectionComments         = "comments"
	CollectionUserInteractions = "user_interactions"
	CollectionAuthTokens       = "auth_tokens"
	CollectionTags             = "tags"
	CollectionCategories       = "categories"
	CollectionAISuggestions    = "ai_suggestions"
	CollectionUserSessions     = "user_sessions"
)
