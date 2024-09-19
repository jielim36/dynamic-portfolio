package routes

import (
	"dynamic-portfolio-be-go/api/controllers"
	"dynamic-portfolio-be-go/repository"
	"dynamic-portfolio-be-go/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProjectRouterInit(parentRouter *gin.RouterGroup, db *gorm.DB) {
	repo := repository.ProjectRepository{
		DB: db,
	}
	usecase := usecase.ProjectUsecase{
		Repo: repo,
	}
	projectController := controllers.ProjectController{
		ProjectUsecase: usecase,
	}

	projectRouter := parentRouter.Group("/projects")
	{
		projectRouter.GET("/", projectController.GetProjects)
		projectRouter.POST("/",projectController.AddProject)
		projectRouter.PUT("/", projectController.UpdateProject)
		projectRouter.DELETE("/:id", projectController.DeleteProjectById)
		projectRouter.GET("/:id", projectController.GetProjectById)
		projectRouter.GET("/group/:groupId", projectController.GetProjectsByGroupId)
	}
}
