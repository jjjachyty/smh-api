package models

import (
	"context"
	"fmt"
	"smh-api/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Player struct {
	ID   string `bson:"_id"`
	Name string
	URL  string
}

func player() *mongo.Collection {
	return db.GetCollection("player")
}

func (m *Player) Get() (err error) {
	fmt.Println(m.ID)
	var result *mongo.SingleResult
	if result = player().FindOne(context.TODO(), bson.M{"_id": m.ID}); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}
