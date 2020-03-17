package models

import (
	"context"
	"fmt"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	ID             string `bson:"_id" binding:"-"` //
	MovieID        string
	MovieName      string
	MovieThumbnail string
	Content        string
	Sender         string //发送者
	SenderAvatar   string `binding:"-"`
	Receiver       string //接收者
	RefID          string //
	LikeCount      int64
	UnLikeCount    int64
	Likes          []string //点赞
	UnLikes        []string //不点赞
	At             []string //@
	CreateAt       time.Time
}

func comment() *mongo.Collection {
	return db.GetCollection("comment")
}

func (m *Comment) Insert() error {
	if _, err := comment().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *Comment) Get(where bson.M) (err error) {
	var result *mongo.SingleResult
	if result = comment().FindOne(context.TODO(), where); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}

func (m Comment) Update(set interface{}) error {
	fmt.Println(set)
	if _, err := comment().UpdateOne(context.TODO(), bson.M{"_id": m.ID}, set); err != nil {
		return err
	}
	return nil
}

func FindComments(where bson.M, offset int64, limit int64, sort bson.M) ([]*Comment, error) {
	var results []*Comment
	var err error
	var cursor *mongo.Cursor
	q := []bson.M{
		{"$match": where},
		{"$lookup": bson.M{
			"from":         "user",
			"localField":   "sender",
			"foreignField": "_id",
			"as":           "user",
		}},
		{"$sort": sort},
		{"$skip": offset},
		{"$limit": limit},
	}
	if cursor, err = comment().Aggregate(context.TODO(), q); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = make(map[string]interface{}, 0)

		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		as := elem["user"].(primitive.A)
		user := as[0].(map[string]interface{})

		likesDB := elem["likes"].(primitive.A)
		var likes []string
		for _, like := range likesDB {
			likes = append(likes, like.(string))
		}

		unlikesDB := elem["unlikes"].(primitive.A)
		var unlikes []string
		for _, like := range unlikesDB {
			unlikes = append(unlikes, like.(string))
		}

		results = append(results, &Comment{ID: elem["_id"].(string), LikeCount: elem["likecount"].(int64), UnLikeCount: elem["unlikecount"].(int64), MovieID: elem["movieid"].(string), MovieName: elem["moviename"].(string), MovieThumbnail: elem["moviethumbnail"].(string), Likes: likes, UnLikes: unlikes, Sender: elem["sender"].(string), Content: elem["content"].(string), SenderAvatar: user["avatar"].(string), CreateAt: elem["createat"].(primitive.DateTime).Time()})
	}
	return results, err
}
