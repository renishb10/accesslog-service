package utils

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Test Data (For Reference)
// 1, 127.0.0.1, 17:51:59
// 2, 127.0.0.1, 17:52:59
// 1, 127.0.0.2, 17:53:59
// 2, 127.0.0.2, 17:54:59
// 2, 127.0.0.3, 17:55:59
// 3, 127.0.0.3, 17:55:59
// 3, 127.0.0.1, 17:56:59
// 4, 127.0.0.1, 17:57:59

func GetTestData() []interface{} {
	ts := time.Now()

	r := []interface{}{
		bson.M{
			"userId":    1,
			"ipAddress": "127.0.0.1",
			"timestamp": ts,
		},
		bson.M{
			"userId":    2,
			"ipAddress": "127.0.0.1",
			"timestamp": ts,
		},
		bson.M{
			"userId":    1,
			"ipAddress": "127.0.0.2",
			"timestamp": ts,
		},
		bson.M{
			"userId":    2,
			"ipAddress": "127.0.0.2",
			"timestamp": ts,
		},
		bson.M{
			"userId":    2,
			"ipAddress": "127.0.0.3",
			"timestamp": ts,
		},
		bson.M{
			"userId":    3,
			"ipAddress": "127.0.0.3",
			"timestamp": ts,
		},
		bson.M{
			"userId":    3,
			"ipAddress": "127.0.0.1",
			"timestamp": ts,
		},
		bson.M{
			"userId":    4,
			"ipAddress": "127.0.0.1",
			"timestamp": ts,
		},
	}
	return r
}
