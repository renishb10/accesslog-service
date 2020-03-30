package services

import "github.com/renishb10/foxg-accesslog-service/models"

type IAccesslogService interface {
	CheckUsersUnique(userIDs *models.UserIDs) (bool, error)
	SeedData(n int) error
	PurgeData() error
}
