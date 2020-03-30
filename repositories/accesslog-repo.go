package repositories

import (
	"log"

	"github.com/renishb10/foxg-accesslog-service/config"
	"github.com/renishb10/foxg-accesslog-service/models"
	"gopkg.in/mgo.v2/bson"
)

type accesslogRepository struct{}

func AccesslogRepository() IAccessRepository {
	return &accesslogRepository{}
}

func (*accesslogRepository) GetUsersIps(userIDs *models.UserIDs) ([]models.UserIpAddresses, error) {
	acl := []models.UserIpAddresses{}
	ids := []int{userIDs.UserID1, userIDs.UserID2}

	selector := []bson.M{
		{
			"$match": bson.M{
				"userId": bson.M{"$in": ids},
			},
		},
		{
			"$sort": bson.M{
				"ipAddress": -1,
			},
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"userId": "$userId",
				},
				"iPs": bson.M{"$addToSet": "$ipAddress"},
			},
		},
		{
			"$project": bson.M{
				"_id":    0,
				"userId": "$_id.userId",
				"iPs":    "$iPs",
			},
		},
	}

	err := config.Accesslogs.Pipe(selector).All(&acl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return acl, nil
}

func (*accesslogRepository) SeedData(accesslogs []interface{}) error {
	bulk := config.Accesslogs.Bulk()

	bulk.Insert(accesslogs...)
	bulk.Unordered()

	_, bulkInsertErr := bulk.Run()
	if bulkInsertErr != nil {
		log.Fatal(bulkInsertErr)
		return bulkInsertErr
	}

	return nil
}

func (*accesslogRepository) PurgeData() error {
	_, err := config.Accesslogs.RemoveAll(nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
