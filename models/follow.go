package models

import (
	"context"
	"fmt"
	"smh-api/db"
	"time"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Follow struct {
	ID           string `bson:"_id" binding:"-"` //
	UserID       string
	FollowID     string
	FollowName   string
	FollowAvatar string
	CreateAt     time.Time
}

func follow() *mongo.Collection {
	return db.GetCollection("follow")
}

func (m *Follow) Insert() error {
	m.ID = xid.New().String()
	if _, err := follow().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *Follow) Get(where bson.M) (err error) {
	var result *mongo.SingleResult
	if result = follow().FindOne(context.TODO(), where); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}

func (m *Follow) Delete() error {
	if _, err := follow().DeleteOne(context.TODO(), bson.M{"userid": m.UserID, "followid": m.FollowID}); err != nil {
		return err
	}
	return nil
}

func Follows(offset int64, limit int64, userid string) ([]*Follow, error) {
	var results []*Follow
	var err error
	var cursor *mongo.Cursor
	q := []bson.M{
		{"$match": bson.M{"userid": userid}}, //在15秒内
		{"$lookup": bson.M{
			"from":         "user",
			"localField":   "followid",
			"foreignField": "_id",
			"as":           "user",
		}}, {"$sort": bson.M{"createat": -1}},
		{"$limit": 15},
	}

	if cursor, err = follow().Aggregate(context.TODO(), q); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = make(map[string]interface{}, 0)

		fmt.Println(cursor.Current.String())
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		as := elem["user"].(primitive.A)
		user := as[0].(map[string]interface{})
		results = append(results, &Follow{FollowID: elem["followid"].(string), FollowName: user["nickname"].(string), FollowAvatar: user["avatar"].(string), CreateAt: elem["createat"].(primitive.DateTime).Time()})
	}
	return results, err
}
