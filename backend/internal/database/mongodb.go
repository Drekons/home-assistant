package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	URI    string
	Client *mongo.Client
}

func NewMongoDB(ctx context.Context, uri string) (*MongoDB, error) {
	var err error
	db := &MongoDB{
		URI: uri,
	}

	db.Client, err = db.connect(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (m *MongoDB) connect(ctx context.Context) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.URI))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
