package router

import (
	"github.com/coderockr/APIVagaFrontend/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(r *gin.Engine) {
	r.GET("/", controllers.APIEndpoints)

	api := r.Group("")
	{

		api.GET("/friends", cors, controllers.GetFriends)
		api.GET("/friends/:id", cors, controllers.GetFriend)
		api.POST("/friends", cors, controllers.CreateFriend)
		api.PUT("/friends/:id", cors, controllers.UpdateFriend)
		api.DELETE("/friends/:id", cors, controllers.DeleteFriend)
		api.OPTIONS("/friends", cors)
		api.OPTIONS("/friends/:id", cors)

		api.GET("/statuses", cors, controllers.GetStatuses)
		api.GET("/statuses/:id", cors, controllers.GetStatus)
		api.POST("/statuses", cors, controllers.CreateStatus)
		api.PUT("/statuses/:id", cors, controllers.UpdateStatus)
		api.DELETE("/statuses/:id", cors, controllers.DeleteStatus)
		api.OPTIONS("/statuses", cors)
		api.OPTIONS("/statuses/:id", cors)

		api.GET("/users", cors, controllers.GetUsers)
		api.GET("/users/:id", cors, controllers.GetUser)
		api.POST("/users", cors, controllers.CreateUser)
		api.PUT("/users/:id", cors, controllers.UpdateUser)
		api.DELETE("/users/:id", cors, controllers.DeleteUser)
		api.OPTIONS("/users", cors)
		api.OPTIONS("/users/:id", cors)

	}
}

func cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, AnonymousToken")
}
