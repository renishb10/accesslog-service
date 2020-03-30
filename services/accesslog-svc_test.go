package services_test

import (
	"testing"

	"github.com/renishb10/foxg-accesslog-service/models"
	"github.com/renishb10/foxg-accesslog-service/services"
)

var (
	accesslogService services.IAccesslogService = services.AccesslogService()
)

func TestCheckUsersUnique(t *testing.T) {
	testStub1 := models.UserIDs{
		UserID1: 1,
		UserID2: 1,
	}

	_, err1 := accesslogService.CheckUsersUnique(&testStub1)
	if err1 != nil {
		t.Error("CheckUsersUnique test failed")
	}

	testStub2 := models.UserIDs{
		UserID1: 1,
		UserID2: 1,
	}

	_, err2 := accesslogService.CheckUsersUnique(&testStub2)
	if err2 != nil {
		t.Log("CheckUsersUnique [ Parameter validation ] passed")
	} else {
		t.Error("CheckUsersUnique [ Parameter validation ] test failed")
	}

	t.Log("CheckUsersUnique test passed")
}
