package main

import (
	"dynamic-portfolio-be-go/api/routes"
	"dynamic-portfolio-be-go/bootstrap"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
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

	r := gin.Default()

	routes.Init(env, r, db)

	// CORS
	var client_url = fmt.Sprintf("http://%s:%s", env.ClientHost, env.ClientPort)
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{client_url}, // 设置允许的域名
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // 设置允许的 HTTP 方法
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 设置允许的头信息
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,   // 是否允许发送 Cookie 信息
        MaxAge: 12 * time.Hour,   // 预检请求的缓存时间
		AllowOriginFunc: func(origin string) bool {
			return origin == client_url
		},
    }))

	r.Use(func(c *gin.Context) {
		log.Println("Request URL:", c.Request.URL)
		log.Println("Request Method:", c.Request.Method)
		c.Next()
		log.Println("Response Status:", c.Writer.Status())
	})


	r.Run(env.ServerAddress)
}
