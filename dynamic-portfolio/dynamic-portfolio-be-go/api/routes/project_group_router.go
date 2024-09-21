package routes

import (
	"dynamic-portfolio-be-go/api/controllers"
	"dynamic-portfolio-be-go/repository"
	"dynamic-portfolio-be-go/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProjectGroupRouterInit(parentRouter *gin.RouterGroup, db *gorm.DB) {
	repo := repository.ProjectGroupRepository{
		DB: db,
	}
	usecase := usecase.ProjectGroupUsecase{
		Repo: repo,
	}
	projectGroupController := controllers.ProjectGroupController{
		ProjectGroupUsecase: usecase,
	}

	projectRouter := parentRouter.Group("/groups")
	{
		projectRouter.GET("/:id", projectGroupController.GetProjectGroupById)
		projectRouter.GET("/user/:id", projectGroupController.GetProjectGroupsByUserId)
		projectRouter.POST("", projectGroupController.AddProjectGroup)
		projectRouter.PUT("", projectGroupController.UpdateProjectGroup)
		projectRouter.DELETE("", projectGroupController.DeleteProjectGroupById)
		projectRouter.OPTIONS("",func(c *gin.Context) {
			c.Status(http.StatusNoContent)
		})
	}
}
