package controllers

import (
	"dynamic-portfolio-be-go/models"
	"dynamic-portfolio-be-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectGroupController struct {
	ProjectGroupUsecase usecase.ProjectGroupUsecase
}

func (p *ProjectGroupController) GetProjectGroupsByUserId(c *gin.Context) {
	id := c.Param("id")
	user_id,err := strconv.ParseUint(id,10,64)
	result, err := p.ProjectGroupUsecase.GetProjectGroupsByUserId(uint(user_id))
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

func (p *ProjectGroupController) GetProjectGroupById(c *gin.Context) {
	id := c.Param("id")
	project_id,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}

	result, err := p.ProjectGroupUsecase.GetProjectGroupById(uint(project_id))
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

func (p *ProjectGroupController) AddProjectGroup(c *gin.Context) {
	var projectGroup models.ProjectGroup
	err := c.ShouldBind(&projectGroup)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	affectedRows, err := p.ProjectGroupUsecase.Create(&projectGroup)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Add Project Group Successful",
		Data: affectedRows,
	})
}

func (p *ProjectGroupController) UpdateProjectGroup(c *gin.Context) {
	var projectGroup models.ProjectGroup
	err := c.ShouldBind(&projectGroup)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	result, err := p.ProjectGroupUsecase.Update(&projectGroup)
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Update Project Group Successful",
		Data: result,
	})
}

func (p *ProjectGroupController) DeleteProjectGroupById(c *gin.Context) {
	id := c.Param("id")
	project_id,err := strconv.ParseUint(id,10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
	}

	result, err := p.ProjectGroupUsecase.DeleteByID(uint(project_id))
		if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return 
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Delete Project Group Successful",
		Data: result,
	})
}