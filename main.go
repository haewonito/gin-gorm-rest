//tutorial: https://www.youtube.com/watch?v=ZI6HaPKHYsg

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
	routes.SongRoute(router)
	routes.AlbumRoute(router)
	routes.AlbumSongRoute(router)

	router.Run(":8080")
}