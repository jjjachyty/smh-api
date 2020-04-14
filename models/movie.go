package models

import (
	"context"
	"errors"
	"fmt"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Movie struct {
	ID        string    `bson:"_id" binding:"-" ` //binding:"-"
	Name      string    `bson:"name"`
	Cover     string    `bson:"cover"`
	Years     string    `bson:"years"`
	Region    string    `bson:"region"`
	Genre     string    `bson:"genre"`
	ScoreDB   int       `bson:"scoreDB"`
	Director  string    `bson:"director"`
	Actor     string    `bson:"actor"`
	DetailURL string    `bson:"detailURL"`
	CreateAt  time.Time `bson:"createAt"`
	CreateBy  string    `bson:"createBy"`
	UpdateAt  time.Time `bson:"updateAt"`
}

func c() *mongo.Collection {
	return db.GetCollection("movie")
}

func (m *Movie) Insert() error {
	var err error
	var count int64
	if count, err = c().CountDocuments(context.TODO(), bson.M{"name": m.Name, "actor": m.Actor}); err == nil && count == 0 {
		if _, err = c().InsertOne(context.TODO(), m); err != nil {
			return err
		}
	}
	if count > 0 {
		err = errors.New("已存在,重复添加")
	}
	fmt.Println(count)
	return err
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

func (m *Movie) Remove(where bson.M) (err error) {
	_, err = c().DeleteMany(context.TODO(), where)
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

func FindMovieGenre() ([]interface{}, error) {
	var err error
	var genres []interface{}
	if genres, err = c().Distinct(context.TODO(), "genre", bson.M{}); err != nil {
		return nil, err
	}
	for _, genre := range genres {
		fmt.Println(genre)
	}

	return genres, err
}
