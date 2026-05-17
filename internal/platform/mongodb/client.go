package mongodb

import (
	"context"
	"time"

	"github.com/artarazdev/API-REST-con-Gin-Golang/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Client struct {
	client   *mongo.Client
	database *mongo.Database
}

func Connect(ctx context.Context, cfg config.MongoConfig) (*Client, error) {
	timeout := time.Duration(cfg.TimeoutSeconds) * time.Second

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &Client{
		client:   client,
		database: client.Database(cfg.Database),
	}, nil
}

func (c *Client) Ping(ctx context.Context) error {
	return c.client.Ping(ctx, readpref.Primary())
}

func (c *Client) Database() *mongo.Database {
	return c.database
}

func (c *Client) Disconnect(ctx context.Context) error {
	return c.client.Disconnect(ctx)
}
