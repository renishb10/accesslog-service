package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/renishb10/foxg-accesslog-service/models"
	"github.com/renishb10/foxg-accesslog-service/repositories"
	"github.com/renishb10/foxg-accesslog-service/utils"
	"gopkg.in/mgo.v2/bson"
)

type accesslogService struct{}

var (
	accesslogRepo repositories.IAccessRepository = repositories.AccesslogRepository()
)

func AccesslogService() IAccesslogService {
	return &accesslogService{}
}

// Checks user's IPs for unique matches
// @Params - UserIDs Model
// Returns - bool, error
func (*accesslogService) CheckUsersUnique(userIDs *models.UserIDs) (bool, error) {
	if userIDs.UserID1 == userIDs.UserID2 {
		return true, nil
	}

	// Gets Each given user's unique IP addresses
	logData, _ := accesslogRepo.GetUsersIps(userIDs)
	var matchIps []string
	count := 0

	// Used Hashmap to finding intersection
	if logData != nil && len(logData) > 1 {
		userIps1 := logData[0].IPs
		userIps2 := logData[1].IPs
		m := make(map[string]bool)

		for _, item := range userIps1 {
			m[item] = true
		}

		for _, item := range userIps2 {
			if count >= 2 {
				break
			}

			if m[item] == true {
				matchIps = append(matchIps, item)
				count++
			}
		}
	}

	if count >= 2 {
		return true, nil
	}

	return false, nil
}

func (*accesslogService) SeedData(n int) error {
	var accesslogArray = make([]interface{}, n)
	ts := time.Now()
	for i := range accesslogArray {
		userId := rand.Intn(9) + 1
		ipAddress := fmt.Sprintf("%d:%d:%d:%d", rand.Intn(254)+1, rand.Intn(254)+1, rand.Intn(254)+1, rand.Intn(9)+1)
		accesslogArray[i] = bson.M{
			"userId":    userId,
			"ipAddress": ipAddress,
			"timestamp": ts,
		}
	}

	// Adds additional test data in spite of n of dummy mock data
	actualTestData := utils.GetTestData()
	accesslogArray = append(accesslogArray, actualTestData...)
	err := accesslogRepo.SeedData(accesslogArray)
	if err != nil {
		return err
	}

	return nil
}

func (*accesslogService) PurgeData() error {
	return accesslogRepo.PurgeData()
}
