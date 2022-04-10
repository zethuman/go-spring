package handler

import (
	"database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(context *gin.Context) {
	var courses []models.Courses
	models.DB.Find(&courses)
	context.JSON(http.StatusOK, gin.H{"courses": courses})
}

func GetById(context *gin.Context) {
	var course models.Courses
	models.DB.Find(&course, "id = ?", context.Param("id"))
	context.JSON(http.StatusOK, gin.H{"courses": course})
}

func CreateCourse(context *gin.Context) {
	var input models.CreateCourseInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course := models.Courses{
		Name:        input.Name,
		Description: input.Description,
		Images:      input.Images,
		Price:       input.Price,
	}

	models.DB.Create(&course)
	context.JSON(http.StatusOK, gin.H{"course": course})
}

func UpdateCourse(context *gin.Context) {
	var course models.Courses
	if err := models.DB.Where("id=?", context.Param("id")).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record is not found"})
		return
	}

	var input models.UpdateCourseInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&course).Update(&input)
	context.JSON(http.StatusOK, gin.H{"course": input})
}

func DeleteCourse(context *gin.Context) {
	var course models.Courses
	if err := models.DB.Where("id=?", context.Param("id")).First(&course).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record is not found"})
		return
	}
	models.DB.Delete(&course)
	context.JSON(http.StatusOK, gin.H{"course": "deleted"})
}
