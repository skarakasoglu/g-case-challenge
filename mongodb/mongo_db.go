// Package mongodb manages mongodb connection operations.
package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// Connection represents a mongo db client connection.
type Connection struct{
	*mongo.Client

	connectionString string
	databaseName string

	timeoutDuration int
}

// NewConnection creates a mongo db connection object.
func NewConnection(connectionString, databaseName string) *Connection {
	return &Connection{
		connectionString: connectionString,
		databaseName: databaseName,
		timeoutDuration: 10,
	}
}

// Connect creates a mongodb client and connects to it by using the connection string.
func (c *Connection) Connect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeoutDuration) * time.Second)
	defer cancel()

	var err error
	c.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(c.connectionString))
	if err != nil {
		log.Fatalf("error on connecting to the database: %v", err)
	}
}

// Disconnect disconnects from the connected mongodb host.
func (c *Connection) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(c.timeoutDuration) * time.Second)
	defer cancel()

	if err := c.Client.Disconnect(ctx); err != nil {
		log.Fatalf("error on disconnecting from the database: %v", err)
	}
}

// DatabaseName returns the default database name.
func (c *Connection) DatabaseName() string {
	return c.databaseName
}