package routes

import (
	"golang-gin/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/api/v1/users", controllers.FindUsers)
	r.POST("/api/v1/users", controllers.CreateUser)
	r.GET("/api/v1/users/:id", controllers.FindUser)
	r.PUT("/api/v1/users/:id", controllers.UpdateUser)
	r.DELETE("api/v1/users/:id", controllers.DeleteUser)
	r.GET("/api/v1/tasks", controllers.FindTasks)
	r.POST("/api/v1/tasks", controllers.CreateTask)
	r.GET("/api/v1/tasks/:id", controllers.FindTask)
	r.PUT("/api/v1/tasks/:id", controllers.UpdateTask)
	r.DELETE("api/v1/tasks/:id", controllers.DeleteTask)
	return r
}
