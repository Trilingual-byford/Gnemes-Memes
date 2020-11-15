package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Gneme struct {
	ID          primitive.ObjectID `bson:"_id"`
	CreatedTime time.Time          `bson:"createdTime"`
	Difficulty  int                `bson:"difficulty"`
	Dir         string             `bson:"dir"`
	Likes       int                `bson:"likes"`
	Viewer      int                `bson:"viewer"`
	Tag         []string           `bson:"tag"`
	PicHash     string             `bson:"picHash"`
	OLSentences []string           `bson:"oLSentences"`
	SLSentences []string           `bson:"sLSentences"`
	Phrase      []string           `bson:"phrase"`
}
