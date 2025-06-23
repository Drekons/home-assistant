package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/Drekons/home-assistant/backend/internal/database"
	"github.com/Drekons/home-assistant/backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const usersCollection = "users"

type UserRepository struct {
	db *database.DB
}

func NewUserRepository(ctx context.Context, db *database.DB) (*UserRepository, error) {
	repo := &UserRepository{db: db}

	err := repo.ensureIndexes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ensure indexes: %w", err)
	}

	return repo, nil
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	collection := r.db.Conn().Collection(usersCollection)

	// Set creation and update times
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Insert the user into the database
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		// Check for duplicate key error (e.g., duplicate email)
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("user with this email already exists")
		}
		return fmt.Errorf("failed to create user: %w", err)
	}

	user.ID = res.InsertedID.(primitive.ObjectID)

	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	collection := r.db.Conn().Collection(usersCollection)

	var user entity.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil // User not found
		}
		return nil, err // Other error occurred
	}

	return &user, nil
}

func (r *UserRepository) ensureIndexes(ctx context.Context) error {
	collection := r.db.Conn().Collection(usersCollection)

	_, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	})

	if err != nil {
		return fmt.Errorf("failed to create unique index on email: %w", err)
	}

	return nil
}
