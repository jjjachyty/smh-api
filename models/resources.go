package models

import (
	"context"
	"smh-api/db"
	"time"

	"github.com/rs/xid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resources struct {
	ID       string `bson:"_id"`
	Platform string
	Name     string
	MovieID  string
	URL      string
	CreateBy int64
	ReportBy int64
	CreateAt time.Time
	UpdateAt time.Time
	State    bool
}

func resources() *mongo.Collection {
	return db.GetCollection("resources")
}

func (m *Resources) Insert() error {
	m.ID = xid.New().String()
	if _, err := resources().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *Resources) Remove(where bson.M) (err error) {
	_, err = resources().DeleteMany(context.TODO(), where)
	return err
}

func FindMovieResources(where bson.M, offset int64, limit int64, sort bson.M) ([]*Resources, error) {
	var results []*Resources
	var err error
	var cursor *mongo.Cursor
	if cursor, err = resources().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(Resources)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
