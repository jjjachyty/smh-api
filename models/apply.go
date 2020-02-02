package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Apply struct {
	ID       string `json:"id" form:"id" query:"id" bson:"_id" binding:"-" ` //binding:"-"
	Name     string
	Describe string
	Users    []string
	State    bool
	CreateAt time.Time
	UpdateAt time.Time
}

func apply() *mongo.Collection {
	return db.GetCollection("apply")
}

func (m *Apply) Insert() error {
	if _, err := apply().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func FindMovieApply(where bson.M, offset int64, limit int64, sort bson.M) ([]*Apply, error) {
	var results []*Apply
	var err error
	var cursor *mongo.Cursor
	if cursor, err = apply().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(Apply)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
