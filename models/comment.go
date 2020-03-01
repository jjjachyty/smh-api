package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Comment struct {
	ID          string `bson:"_id" binding:"-"` //
	MovieID     string
	Content     string
	Sender      string //发送者
	Receiver    string //接收者
	RefID       string //
	LikeCount   int64
	UnLikeCount int64
	Likes       []string //点赞
	UnLikes     []string //不点赞
	At          []string //@
	CreateAt    time.Time
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
	if _, err := comment().UpdateMany(context.TODO(), bson.M{"_id": m.ID}, set); err != nil {
		return err
	}
	return nil
}

func FindComments(where bson.M, offset int64, limit int64, sort bson.M) ([]*Comment, error) {
	var results []*Comment
	var err error
	var cursor *mongo.Cursor
	if cursor, err = comment().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(Comment)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}
