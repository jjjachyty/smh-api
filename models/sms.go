package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SMS struct {
	ID       string
	Phone    string
	Code     string
	CreateAt time.Time
}

func sms() *mongo.Collection {
	return db.GetCollection("sms")
}

func (m *SMS) Insert() error {
	m.CreateAt = time.Now()
	if _, err := sms().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}
func (m *SMS) Update() error {
	if _, err := sms().UpdateOne(context.TODO(), bson.M{"phone": m.Phone}, bson.M{"$set": bson.M{"code": m.Code, "createat": time.Now()}}); err != nil {
		return err
	}
	return nil
}

func (m *SMS) Delete() error {
	if _, err := sms().DeleteOne(context.TODO(), bson.M{"phone": m.Phone}); err != nil {
		return err
	}
	return nil
}

func (m *SMS) Get() (err error) {
	var result *mongo.SingleResult
	if result = sms().FindOne(context.TODO(), bson.M{"phone": m.Phone}); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}
