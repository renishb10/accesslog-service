package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type AccessLog struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    int           `json:"userId" bson:"userId"`
	IpAddress string        `json:"ipAddress" bson:"ipAddress"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}
