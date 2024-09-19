package models

import (
	"time"
)

type Project struct {
	BaseModel
	GroupId   uint       `json:"group_id" gorm:"foreignKey:GroupId" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   *time.Time    `json:"end_date,omitempty"`   // 使用指针，以便处理可选字段
	ImagesURL []string  `json:"images_url,omitempty" gorm:"-"` // 使用 `omitempty` 标签以处理可选字段
}

type ProjectGroup struct {
	BaseModel
	UserId   uint       `json:"user_id" binding:"required"`
	Name     string    `json:"name" binding:"required"`
	Projects []Project `json:"projects,omitempty" gorm:"foreignKey:GroupId"`
}