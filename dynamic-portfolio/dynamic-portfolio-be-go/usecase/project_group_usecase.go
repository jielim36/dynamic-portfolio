package usecase

import (
	"dynamic-portfolio-be-go/models"
	"dynamic-portfolio-be-go/repository"
)

type ProjectGroupUsecase struct {
	Repo repository.ProjectGroupRepository
}

func (pg *ProjectGroupUsecase) Create(projectGroup *models.ProjectGroup) (int, error) {
	return pg.Repo.Create(projectGroup)
}

func (pg *ProjectGroupUsecase) Update(projectGroup *models.ProjectGroup) (int, error){
	return pg.Repo.Update(projectGroup)
}

func (pg *ProjectGroupUsecase) DeleteByID(id uint) (int, error){
	return pg.Repo.DeleteByID(id)
}

func (pg *ProjectGroupUsecase) GetProjectGroupById(id uint) (models.ProjectGroup, error){
	return pg.Repo.GetProjectGroupById(id)
}

func (pg *ProjectGroupUsecase) GetProjectGroupsByUserId(userId uint) ([]models.ProjectGroup, error){
	return pg.Repo.GetProjectGroupsByUserId(userId)
}
