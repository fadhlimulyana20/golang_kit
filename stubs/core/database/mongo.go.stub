package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoConn struct {
	host     string
	port     int
	user     string
	password string
	database string
}

type MongoDB interface {
	Close(ctx context.Context, client *mongo.Client)
	Client() (*mongo.Client, error)
	Database() (*mongo.Database, *mongo.Client, error)
}

func NewMongoDB(host string, port int, user, password, database string) MongoDB {
	return &mongoConn{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		database: database,
	}
}

func (m *mongoConn) Close(ctx context.Context, client *mongo.Client) {
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func (m *mongoConn) Client() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", m.user, m.password, m.host, m.port)))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (m *mongoConn) Database() (*mongo.Database, *mongo.Client, error) {
	client, err := m.Client()
	if err != nil {
		return nil, nil, err
	}

	db := client.Database(m.database)
	return db, client, nil
}
