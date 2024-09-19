package usecase

import (
	"dynamic-portfolio-be-go/models"
	"dynamic-portfolio-be-go/repository"
)

type ProjectUsecase struct {
	Repo repository.ProjectRepository
}

func (p *ProjectUsecase) Create(project *models.Project) (int, error) {
	return p.Repo.Create(project)
}

func (p *ProjectUsecase) Update(project *models.Project) (int, error){
	return p.Repo.Update(project)
}

func (p *ProjectUsecase) DeleteByID(id uint) (int, error){
	return p.Repo.DeleteByID(id)

}

func  (p *ProjectUsecase) GetProjectsByGroupId(groupId uint) ([]models.Project, error){
	return p.Repo.GetProjectsByGroupId(groupId)

}

func (p *ProjectUsecase) GetProjectById(id uint) (models.Project, error){
	return p.Repo.GetProjectById(id)
}

