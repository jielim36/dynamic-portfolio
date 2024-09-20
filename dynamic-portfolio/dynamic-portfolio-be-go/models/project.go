package models

import (
	"time"
)

type ProjectGroup struct {
	BaseModel
	UserId   uint       `json:"user_id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Projects []Project `json:"projects,omitempty" gorm:"foreignKey:GroupId"`
}

type Project struct {
	BaseModel
	GroupId   uint       `json:"group_id" gorm:"foreignKey:GroupId" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	Description string `json:"description"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   *time.Time    `json:"end_date,omitempty"`   // 使用指针，以便处理可选字段
	ProjectImages []ProjectImages `json:"project_images" gorm:"foreignKey:ProjectId"`
}

type ProjectImages struct {
	BaseModel
	ProjectId uint `json:"project_id" binding:"required" gorm:"foreignKey:ProjectId"`
	ImageUrl string `json:"image_url" binding:"required"`
}