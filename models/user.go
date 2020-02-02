package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID             string `bson:"_id" binding:"-"` //
	Mail           string
	NickName       string
	Avatar         string
	Introduce      string
	Sex            int
	Phone          string
	PassWord       string
	DeviceID       string
	DevicePlatform string
	IP             string
	Coin           float64
	Longitude      float64
	Latitude       float64
	CreateAt       time.Time
	LastLogin      time.Time
	State          bool
}

func userCollection() *mongo.Collection {
	return db.GetCollection("user")
}

func (m *User) Insert() error {
	if _, err := userCollection().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *User) Get(where bson.M) (err error) {
	var result *mongo.SingleResult
	if result = userCollection().FindOne(context.TODO(), where); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}

func (m *User) Update(set bson.M) error {
	if _, err := userCollection().UpdateOne(context.TODO(), bson.M{"_id": m.ID}, set); err != nil {
		return err
	}
	return nil
}

func FindUsers(where bson.M, offset int64, limit int64, sort bson.M) ([]*User, error) {
	var results []*User
	var err error
	var cursor *mongo.Cursor
	if cursor, err = userCollection().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(User)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
