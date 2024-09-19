package main

import (
	"dynamic-portfolio-be-go/api/routes"
	"dynamic-portfolio-be-go/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db, err := bootstrap.NewMySQLDatabase(env)
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}
	defer bootstrap.CloseMySQLDB(db)

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	routes.Init(env, gin, db)

	gin.Run(env.ServerAddress)
}
