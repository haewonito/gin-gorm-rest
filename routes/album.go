package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-rest/controllers"
)

func AlbumRoute(router *gin.Engine) {
	router.GET("/albums", controller.GetAlbums)
	router.POST("/albums", controller.CreateAlbum)
	router.GET("/albums/:id", controller.GetAlbumById)
	router.DELETE("/albums/:id", controller.DeleteAlbum)
	router.PUT("/albums/:id", controller.UpdateAlbum)
}