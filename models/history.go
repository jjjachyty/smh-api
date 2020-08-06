package models

import (
	"context"
	"fmt"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Watching struct {
	MovieID        string
	MovieThumbnail string
	Count          int32
}

type WatchingHistory struct {
	UserID         int64
	VideoID        string
	VideoName      string
	ResourcesID    string
	ResourcesName  string
	VideoThumbnail string
	MovieDuration  float64
	Progress       float64
	Finish         bool
	CreateAt       time.Time
	UpdateAt       time.Time
}

func watchingHistoryCollection() *mongo.Collection {
	return db.GetCollection("history")
}

func (m *WatchingHistory) Insert() error {
	fmt.Println(m.VideoThumbnail)
	if _, err := watchingHistoryCollection().InsertOne(context.TODO(), m); err != nil {
		return err
	}
	return nil
}

func (m *WatchingHistory) Remove(where bson.M) (err error) {
	_, err = watchingHistoryCollection().DeleteMany(context.TODO(), where)
	return err
}

func (m *WatchingHistory) Update() error {
	if _, err := watchingHistoryCollection().UpdateOne(context.TODO(), bson.M{"userid": m.UserID, "movieid": m.VideoID, "resourcesid": m.ResourcesID}, bson.M{"$set": bson.M{"movieduration": m.MovieDuration, "progress": m.Progress, "updateat": time.Now()}}); err != nil {
		return err
	}
	return nil
}

func (m *WatchingHistory) Get(where bson.M) (err error) {
	var result *mongo.SingleResult
	if result = watchingHistoryCollection().FindOne(context.TODO(), where); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil
		}
		return result.Err()
	}
	err = result.Decode(m)

	return err
}

func FindWatchHistorys(where bson.M, offset int64, limit int64, sort bson.M) ([]*WatchingHistory, error) {
	var results []*WatchingHistory
	var err error
	var cursor *mongo.Cursor
	if cursor, err = watchingHistoryCollection().Find(context.TODO(), where, options.Find().SetSort(sort).SetSkip(offset).SetLimit(limit)); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = new(WatchingHistory)
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		results = append(results, elem)
	}
	return results, err
}

func FindWatching(offset int64, limit int64) ([]*Watching, error) {
	var results []*Watching
	var err error
	var cursor *mongo.Cursor
	q := []bson.M{

		{"$match": bson.M{"updateat": bson.M{"$gte": time.Now().Add(time.Minute * -1)}}}, //在15秒内
		{"$project": bson.M{
			"_id":            0,
			"movieid":        1,
			"moviethumbnail": 1,
		}},
		{"$group": bson.M{"_id": bson.M{"movieid": "$movieid", "moviethumbnail": "$moviethumbnail"}, "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": 15},
	}

	if cursor, err = watchingHistoryCollection().Aggregate(context.TODO(), q); err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var elem = make(map[string]interface{}, 0)

		fmt.Println(cursor.Current.String())
		if err = cursor.Decode(&elem); err != nil {
			return nil, err
		}
		fmt.Println(elem)
		movies := elem["_id"].(map[string]interface{})
		results = append(results, &Watching{MovieID: movies["movieid"].(string), MovieThumbnail: movies["moviethumbnail"].(string), Count: elem["count"].(int32)})
	}
	return results, err
}
