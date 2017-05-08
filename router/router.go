package router

import (
	"github.com/coderockr/APIVagaFrontend/controllers"

	"github.com/gin-gonic/gin"
	"net/http"
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
		api.OPTIONS("/friends", cors)
		api.OPTIONS("/friends/:id", cors)

		api.GET("/statuses", controllers.GetStatuses)
		api.GET("/statuses/:id", controllers.GetStatus)
		api.POST("/statuses", controllers.CreateStatus)
		api.PUT("/statuses/:id", controllers.UpdateStatus)
		api.DELETE("/statuses/:id", controllers.DeleteStatus)
		api.OPTIONS("/statuses", cors)
		api.OPTIONS("/statuses/:id", cors)

		api.GET("/users", controllers.GetUsers)
		api.GET("/users/:id", controllers.GetUser)
		api.POST("/users", controllers.CreateUser)
		api.PUT("/users/:id", controllers.UpdateUser)
		api.DELETE("/users/:id", controllers.DeleteUser)
		api.OPTIONS("/users", cors)
		api.OPTIONS("/users/:id", cors)

	}
}

func cors(c *gin.Context) {
    c.Header("Access-Control-Allow-Origin", "*")
    c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
    c.JSON(http.StatusOK, struct{}{})
}
