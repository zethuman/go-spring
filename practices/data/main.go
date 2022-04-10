package main

import (
	"database/handler"
	"database/models"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	models.ConnectDB()

	route := gin.Default()
	route.GET("/courses", handler.GetAll)
	route.GET("/courses/:id", handler.GetById)
	route.POST("/courses", handler.CreateCourse)
	route.PUT("/courses/:id", handler.UpdateCourse)
	route.DELETE("/courses/:id", handler.DeleteCourse)

	route.Run(":8000")
}
