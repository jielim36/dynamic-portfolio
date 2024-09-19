package controllers

import (
	"dynamic-portfolio-be-go/models"
	"dynamic-portfolio-be-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	ProjectUsecase usecase.ProjectUsecase
}

func (p *ProjectController) GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, "hi")
}

func (p *ProjectController) AddProject(c *gin.Context) {
	var project models.Project
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	_, err = p.ProjectUsecase.Create(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Add Project Successful",
		Data: true,
	})
}

func (p *ProjectController) UpdateProject(c *gin.Context) {
	var project models.Project
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	result, err := p.ProjectUsecase.Update(&project)
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Update Project Successful",
		Data: result,
	})
}

func (p *ProjectController) DeleteProjectById(c *gin.Context) {
	id := c.Param("id")
	project_id,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}

	result, err := p.ProjectUsecase.DeleteByID(uint(project_id))
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Delete Project Successful",
		Data: result,
	})
}

func (p *ProjectController) GetProjectsByGroupId(c *gin.Context) {
	id := c.Param("groupId")
	groupId,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}

	result, err := p.ProjectUsecase.GetProjectsByGroupId(uint(groupId))
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Successful",
		Data: result,
	})
}

func (p *ProjectController) GetProjectById(c *gin.Context) {
	id := c.Param("id")
	project_id,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}

	result, err := p.ProjectUsecase.GetProjectById(uint(project_id))
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Successful",
		Data: result,
	})
}