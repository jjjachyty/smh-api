package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func InitDB() {
	var err error
	var mongoURI = "mongodb://smh:smh~2019@127.0.0.1:27117"
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	//    "mongodb+srv://<username>:<password>@<cluster-address>/test?w=majority"
	// ))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI)); err != nil {
		panic(err)
	}

	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}

func GetDB() *mongo.Database {
	return client.Database("smh")
}

func GetCollection(cname string) *mongo.Collection {
	return client.Database("smh").Collection(cname)
}
