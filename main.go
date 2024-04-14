//tutorial: https://www.youtube.com/watch?v=ZI6HaPKHYsg

// todo haewon - try to make many to many?  wahat would be many to many for album and songs?
//artists and albums

// todo haewon - try to put some validations.  use that panic thing? 
// also there must be something like "require" tag when I make the gorm models.
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