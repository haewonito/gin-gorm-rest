package routes

import( "github.com/gin-gonic/gin"
		"github.com/haewonito/gin-gorm-rest/controllers"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users", controller.GetUsers)
	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserById)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.PUT("/users/:id", controller.UpdateUser)
}