package models

import (
	"context"
	"smh-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Version struct {
	VersionCode string
	ReleaseTime time.Time
	ReleaseNote string
	DownloadURL string
}

func versionCollection() *mongo.Collection {
	return db.GetCollection("version")
}

func GetVsersion(platform string) (*Version, error) {
	var err error
	var result *mongo.SingleResult
	var version = new(Version)
	if result = versionCollection().FindOne(context.TODO(), bson.M{"platform": platform}, options.FindOne().SetSort(bson.M{"versioncode": -1})); result.Err() != nil {
		if result.Err().Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, result.Err()
	}
	err = result.Decode(version)

	return version, err
}
