package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-rest/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)
}