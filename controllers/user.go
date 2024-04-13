package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-rest/models"
	"github.com/haewonito/gin-gorm-rest/config"
)

func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}