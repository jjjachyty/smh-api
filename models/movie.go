package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID       string    `bson:"_id" binding:"-" ` //binding:"-"
	Name     string    `bson:"name"`
	Cover    string    `bson:"cover"`
	Years    string    `bson:"years"`
	Region   string    `bson:"region"`
	Genre    string    `bson:"genre"`
	ScoreDB  int       `bson:"scoreDB"`
	Director string    `bson:"director"`
	Actor    string    `bson:"actor"`
	CreateAt time.Time `bson:"createAt"`
	UpdateAt time.Time `bson:"updateAt"`
}

func c() *mongo.Collection {
	return db.GetCollection("movie")
}

func (m *Movie) Insert() error {
	if _, err := c().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *Movie) Get(where bson.M) (err error) {
	var result *mongo.SingleResult
	if result = c().FindOne(context.TODO(), where); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}

func FindMovie(where bson.M, offset int64, limit int64, sort bson.M) ([]*Movie, error) {
	var results []*Movie
	var err error
	var cursor *mongo.Cursor
	if cursor, err = c().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(Movie)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
