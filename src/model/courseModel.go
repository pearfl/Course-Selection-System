package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Course struct {
	CourseID    string `json:"course_id" form:"course_id" gorm:"primary_key"`
	TeacherID   string `json:"teacher_id" form:"teacher_id"`
	Name        string `json:"name" form:"name"`
	Capacity    int    `json:"capacity" form:"capacity"`
	CapSelected int    `json:"cap_select" form:"cap_select"`
}

func (Course) TableName() string {
	return "course"
}

func (course *Course) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4().String()
	return scope.SetColumn("course_id", uuid)
}