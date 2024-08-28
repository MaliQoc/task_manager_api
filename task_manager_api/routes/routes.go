package routes

import (
	"github.com/gin-gonic/gin"
	"restapi_base/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", controllers.GetHome)
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.POST("/tasks", controllers.CreateTask)

	return router
}
