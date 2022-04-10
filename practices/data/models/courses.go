package models

type Courses struct {
	ID          uint   `json:"-" gorm:"primary_key"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Images      string `json:"images,omitempty"`
	Price       uint   `json:"price,omitempty"`
}

type CreateCourseInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Images      string `json:"images" binding:"required"`
	Price       uint   `json:"price" binding:"required"`
}

type UpdateCourseInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Images      string `json:"images,omitempty"`
	Price       uint   `json:"price,omitempty"`
}
