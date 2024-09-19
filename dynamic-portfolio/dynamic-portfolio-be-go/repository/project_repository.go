package repository

import (
	"dynamic-portfolio-be-go/models"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	DB *gorm.DB
}

// func handleError(err error) {
// 	defer 

// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func (p *ProjectRepository) Create(project *models.Project) (int, error){
	result := p.DB.Create(&project)
	return int(result.RowsAffected), result.Error
}

func (p *ProjectRepository) Update(project *models.Project) (int, error){
	result := p.DB.Save(&project)
	return int(result.RowsAffected), result.Error
}

func (p *ProjectRepository) DeleteByID(id uint) (int, error){
	result := p.DB.Delete(&models.Project{}, id)
	return int(result.RowsAffected), result.Error
}

func  (p *ProjectRepository) GetProjectsByGroupId(groupId uint) ([]models.Project, error){
	var projects []models.Project
	result := p.DB.Where("group_id = ?", groupId).Find(&projects)
	return projects, result.Error
}

func (p *ProjectRepository) GetProjectById(id uint) (models.Project, error){
	var project models.Project
	result := p.DB.First(&project, id)
	return project, result.Error
}

