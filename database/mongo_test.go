package database

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	mongoHost     = "localhost"
	mongoPort     = 27017
	mongoUser     = "root"
	mongoPassword = "password"
	mongoDatabase = "test_go_kit"
)

func TestConn(t *testing.T) {
	mongoDB := NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)
	client, err := mongoDB.Client()
	if err != nil {
		t.Fatal(err)
	}

	defer mongoDB.Close(context.Background(), client)
}

func TestInsertOne(t *testing.T) {
	mongoDB := NewMongoDB(mongoHost, mongoPort, mongoUser, mongoPassword, mongoDatabase)
	db, client, err := mongoDB.Database()
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = db.Collection("testing").InsertOne(ctx, bson.D{{"hello", "world"}})
	if err != nil {
		t.Fatal(err)
	}

	defer mongoDB.Close(context.Background(), client)
}
