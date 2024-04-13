package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-rest/controllers"
)

func AlbumSongRoute(router *gin.Engine) {
	router.GET("/albums/:id/songs", controller.GetSongByAlbumId)
}