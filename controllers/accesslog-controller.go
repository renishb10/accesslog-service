package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/renishb10/foxg-accesslog-service/models"
	"github.com/renishb10/foxg-accesslog-service/services"
)

var (
	accesslogService services.IAccesslogService = services.AccesslogService()
)

// Checks given users (ids) has 2 unique duplicate Ips
// @Params - :userId1/:userId2
// Returns - JSON
func CheckUsersUnique(c *gin.Context) {
	var userIDs models.UserIDs
	if paramErr := c.ShouldBindUri(&userIDs); paramErr != nil {
		c.JSON(400, gin.H{"message": paramErr})
		return
	}

	isUnique, err := accesslogService.CheckUsersUnique(&userIDs)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"dupes": isUnique})
}

// Seed sample data for testing
// @Params - :count
// Returns - JSON
func SeedData(c *gin.Context) {
	count, cErr := strconv.Atoi(c.Param("count"))
	if cErr != nil {
		c.JSON(400, gin.H{"message": cErr})
	}

	err := accesslogService.SeedData(count)
	if err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"message": "Seeded records:" + c.Param("count")})
}

// Truncates data
// Returns - JSON
func PurgeData(c *gin.Context) {
	err := accesslogService.PurgeData()
	if err != nil {
		c.JSON(400, gin.H{"message": err})
	}
	c.JSON(200, gin.H{"message": "Purged all the Data"})
}
