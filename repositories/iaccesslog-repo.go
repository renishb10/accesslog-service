package repositories

import "github.com/renishb10/foxg-accesslog-service/models"

type IAccessRepository interface {
	GetUsersIps(userIDs *models.UserIDs) ([]models.UserIpAddresses, error)
	SeedData(accesslogs []interface{}) error
	PurgeData() error
}
