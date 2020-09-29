package mongodb

import (
	"context"
	"time"

	"github.com/qiniu/qmgo"

	"ginner/config"
	"ginner/logger"
)

var (
	ctx    context.Context
	client *qmgo.Client
	cancel func()
)

// Init MongoDB connect initialization
func Init() {
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

	var err error
	address := config.Conf().Mongo.Address
	username := config.Conf().Mongo.Username
	password := config.Conf().Mongo.Password

	uri := "mongodb://" + address
	if username != "" && password != "" {
		uri = "mongodb://" + username + ":" + password + "@" + address
	}

	client, err = qmgo.NewClient(
		ctx,
		&qmgo.Config{Uri: uri},
	)
	if err != nil {
		logger.Log.Errorf("mongodb initialization failed: %v", err)
		return
	}

}

// GetCollection get a collection(table) of the default database
func GetCollection(collectionName string) *qmgo.Collection {
	dbname := config.Conf().Mongo.DBName
	return client.Database(dbname).Collection(collectionName)
}

// GetDBCollection get a collection(table) of the specific database
func GetDBCollection(dbName, collectionName string) *qmgo.Collection {
	return client.Database(dbName).Collection(collectionName)
}

// Close close MongoDB connection when the server stopped
func Close() {
	if err := client.Close(ctx); err != nil {
		logger.Log.Errorf("mongodb connection closed failed: %v", err)
	}

	cancel()
}
