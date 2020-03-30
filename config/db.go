package config

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

var DB *mgo.Database

var Accesslogs *mgo.Collection

func init() {
	dbConn := fmt.Sprintf("mongodb://%s/%s", os.Getenv("MONGO_URL"), os.Getenv("MONGO_DB_NAME"))
	session, err := mgo.Dial(dbConn)
	if err != nil {
		panic(err)
	}

	if err = session.Ping(); err != nil {
		panic(err)
	}

	DB = session.DB(os.Getenv("MONGO_DB_NAME"))
	Accesslogs = DB.C(os.Getenv("MONGO_DB_COL_ACCESSLOG"))

	// Indexing AccessLog
	Accesslogs.EnsureIndexKey("userId", "ipAddress")

	fmt.Println("You connected to your mongo database.")
}
