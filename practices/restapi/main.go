package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Postman collection can be find ./Golang Spring.postman_collection.json
type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Images      string `json:"images"`
	Price       int    `json:"price"`
}

var courses = []Course{
	{ID: "1", Name: "Golang for Junior", Description: "Golang for Junior", Images: "1.jpg", Price: 1000},
	{ID: "2", Name: "Golang for Middle", Description: "Golang for Middle", Images: "2.jpg", Price: 2000},
	{ID: "3", Name: "Golang for Senior", Description: "Golang for Senior", Images: "3.jpg", Price: 3000},
}

func getCourses(c *gin.Context) {
	c.JSON(http.StatusOK, courses)
}

func postCourse(c *gin.Context) {
	var newCourse Course
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	courses = append(courses, newCourse)
	c.JSON(http.StatusCreated, newCourse)
}

func getCourseByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range courses {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

func deleteCourseByID(c *gin.Context) {
	id := c.Param("id")
	for i, a := range courses {
		if a.ID == id {
			courses = append(courses[:i], courses[i+1:]...)
			return
		}
	}
	c.AbortWithStatus(http.StatusNotFound)
}

func main() {
	router := gin.Default()
	router.GET("/courses", getCourses)
	router.POST("/courses", postCourse)
	router.GET("/courses/:id", getCourseByID)
	router.DELETE("/courses/:id", deleteCourseByID)

	router.Run(":8000")
}
