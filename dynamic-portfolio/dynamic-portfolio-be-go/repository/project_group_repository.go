package repository

import (
	"dynamic-portfolio-be-go/models"

	"gorm.io/gorm"
)

type ProjectGroupRepository struct {
	DB *gorm.DB
}

func (pg *ProjectGroupRepository) Create(projectGroup *models.ProjectGroup) (int, error){
	result := pg.DB.Create(&projectGroup)
	return int(result.RowsAffected), result.Error
}

func (pg *ProjectGroupRepository) Update(projectGroup *models.ProjectGroup) (int, error){
	result := pg.DB.Save(&projectGroup)
	return int(result.RowsAffected), result.Error
}

func (pg *ProjectGroupRepository) DeleteByID(id uint) (int, error){
	result := pg.DB.Delete(&models.ProjectGroup{}, id)
	return int(result.RowsAffected), result.Error
}

func (pg *ProjectGroupRepository) GetProjectGroupById(id uint) (models.ProjectGroup, error){
	var projectGroup models.ProjectGroup
	result := pg.DB.Preload("Projects").First(&projectGroup, id)	
	return projectGroup, result.Error
}

func (pg *ProjectGroupRepository) GetProjectGroupsByUserId(user_id uint) ([]models.ProjectGroup, error){
	var projectGroup []models.ProjectGroup
	// result := pg.DB.Preload("Projects").Where("user_id = ?", user_id).Find(&projectGroup)	
	result := pg.DB.
		Preload("Projects.ProjectImages").
		Where("user_id = ?", user_id).
		Find(&projectGroup)
	return projectGroup, result.Error
}
