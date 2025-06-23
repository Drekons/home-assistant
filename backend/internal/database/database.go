package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	Client *MongoDB
	dbname string
}

func NewDB(ctx context.Context, uri string, dbname string) (*DB, error) {
	conn, err := NewMongoDB(ctx, uri)
	if err != nil {
		return nil, err
	}

	return &DB{Client: conn, dbname: dbname}, nil
}

func (d *DB) Conn() *mongo.Database {
	return d.Client.Client.Database(d.dbname)
}

func (d *DB) Close(ctx context.Context) error {
	return d.Client.Client.Disconnect(ctx)
}
