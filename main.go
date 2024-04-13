package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haewonito/gin-gorm-rest/config"
	"github.com/haewonito/gin-gorm-rest/routes"
)

func main() {
	router := gin.New()
	
		//connect database
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}