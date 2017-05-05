package router

import (
	"github.com/coderockr/APIVagaFrontend/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	r.GET("/", controllers.APIEndpoints)

	api := r.Group("")
	{

		api.GET("/friends", controllers.GetFriends)
		api.GET("/friends/:id", controllers.GetFriend)
		api.POST("/friends", controllers.CreateFriend)
		api.PUT("/friends/:id", controllers.UpdateFriend)
		api.DELETE("/friends/:id", controllers.DeleteFriend)

		api.GET("/statuses", controllers.GetStatuses)
		api.GET("/statuses/:id", controllers.GetStatus)
		api.POST("/statuses", controllers.CreateStatus)
		api.PUT("/statuses/:id", controllers.UpdateStatus)
		api.DELETE("/statuses/:id", controllers.DeleteStatus)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)

	}
}
