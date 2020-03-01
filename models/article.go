package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
	ID       string `bson:"_id" binding:"-"`
	Title    string
	Context  string
	CreateAt time.Time
	CreateBy string
}

func article() *mongo.Collection {
	return db.GetCollection("article")
}

func (m *Article) Insert() error {
	if _, err := article().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *Article) Remove(where bson.M) (err error) {
	_, err = article().DeleteMany(context.TODO(), where)
	return err
}

func FindArticles(where bson.M, offset int64, limit int64, sort bson.M) ([]*Article, error) {
	var results []*Article
	var err error
	var cursor *mongo.Cursor
	if cursor, err = article().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(Article)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
