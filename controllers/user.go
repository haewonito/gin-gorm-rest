package controller
//https://gorm.io/docs/query.html
//https://gorm.io/docs/create.html
import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-rest/models"
	"github.com/haewonito/gin-gorm-rest/config"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func GetUserById(c *gin.Context) {
	var user models.User
	config.DB.First(&user, c.Param("id"))
	c.JSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}