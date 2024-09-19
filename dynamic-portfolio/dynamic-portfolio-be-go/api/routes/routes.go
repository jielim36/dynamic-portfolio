package routes

import (
	"dynamic-portfolio-be-go/api/controllers"
	"dynamic-portfolio-be-go/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(env *bootstrap.Env, gin *gin.Engine, db *gorm.DB) {
	publicRouter := gin.Group("") 
	{
		publicRouter.POST("/login", controllers.Login)
	}

	protectedRouter := gin.Group("")
	{
		ProjectRouterInit(protectedRouter, db)
		ProjectGroupRouterInit(protectedRouter,db)
	}
}