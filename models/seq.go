package models

import (
	"context"
	"smh-api/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Seq struct {
	ID            string `bson:"_id"`
	SequenceValue int64  `bson:"sequence_value"`
}

//GetNextSeqID 获取递增序列
func GetNextSeqID() (int64, error) {
	var err error
	var result *mongo.SingleResult
	var after = options.After

	result = db.GetCollection("seq").FindOneAndUpdate(context.TODO(), bson.M{"_id": "seqid"}, bson.M{"$inc": bson.M{"sequence_value": 1}}, &options.FindOneAndUpdateOptions{ReturnDocument: &after})
	var seq = &Seq{}
	if err = result.Decode(seq); err != nil {
		return 0, err
	}

	return seq.SequenceValue, nil
}
