package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-rest/controllers"
)

func SongRoute(router *gin.Engine) {
	router.GET("/songs", controller.GetSongs)
	router.POST("/songs", controller.CreateSong)
	router.GET("/songs/:id", controller.GetSongById)
	router.DELETE("/songs/:id", controller.DeleteSong)
	router.PUT("/songs/:id", controller.UpdateSong)
}