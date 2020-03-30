package models

type UserIDs struct {
	UserID1 int `uri:"userID1" binding:"required"`
	UserID2 int `uri:"userID2" binding:"required"`
}

type UserIpAddresses struct {
	UserId int      `json:"userId" bson:"userId"`
	IPs    []string `json:"iPs" bson:"iPs"`
}
