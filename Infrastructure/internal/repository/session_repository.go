package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"Blog-API/internal/domain"
	"Blog-API/internal/infrastructure/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SessionRepository struct {
	db         *database.MongoDB
	collection *mongo.Collection
}

func NewSessionRepository(db *database.MongoDB) domain.SessionRepository {
	collection := db.GetCollection("sessions")
	
	indexModels := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
		},
		{
			Keys:    bson.D{{Key: "expires_at", Value: 1}},
		},
	}
	
	_, err := collection.Indexes().CreateMany(context.Background(), indexModels)
	if err != nil {
		// Log error but don't fail - indexes might already exist
		fmt.Printf("DEBUG: Failed to create session indexes: %v\n", err)
	} else {
		fmt.Printf("DEBUG: Session indexes created successfully\n")
	}
	
	return &SessionRepository{
		db:         db,
		collection: collection,
	}
}

func (r *SessionRepository) Create(session *domain.Session) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	session.CreatedAt = now
	session.LastActivity = now

	fmt.Printf("DEBUG: Inserting session into database: %+v\n", session)
	result, err := r.collection.InsertOne(ctx, session)
	if err != nil {
		fmt.Printf("DEBUG: Session insertion failed: %v\n", err)
		// check if user already has a session
		var mongoErr mongo.WriteException
		if errors.As(err, &mongoErr) {
			for _, writeErr := range mongoErr.WriteErrors {
				if writeErr.Code == 11000 { // duplicate key error
					fmt.Printf("DEBUG: Duplicate key error, updating session\n")
					return r.Update(session)
				}
			}
		}
		return err
	}

	fmt.Printf("DEBUG: Session inserted successfully with ID: %v\n", result.InsertedID)
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		session.ID = oid
	}

	return nil
}

func (r *SessionRepository) GetByID(id primitive.ObjectID) (*domain.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session domain.Session
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	return &session, nil
}

func (r *SessionRepository) GetByUserID(userID primitive.ObjectID) (*domain.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session domain.Session
	err := r.collection.FindOne(ctx, bson.M{"user_id": userID}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("session not found")
		}
		return nil, err
	}
	return &session, nil
}

func (r *SessionRepository) GetByUsername(username string) (*domain.Session, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var session domain.Session
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("session not found")
		}
		return nil, err
	}

	return &session, nil
}

func (r *SessionRepository) Update(session *domain.Session) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	session.LastActivity = time.Now()

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"user_id": session.UserID},
		bson.M{"$set": session},
	)
	return err
}

func (r *SessionRepository) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *SessionRepository) DeleteByUserID(userID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"user_id": userID})
	return err
}

func (r *SessionRepository) DeleteExpired() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteMany(ctx, bson.M{"expires_at": bson.M{"$lt": time.Now()}})
	return err
}

func (r *SessionRepository) UpdateLastActivity(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"last_activity": time.Now()}},
	)
	return err
} 