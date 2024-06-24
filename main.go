package main

import (
	"github.com/dxckboi/hugeman-exam/config"
	"github.com/dxckboi/hugeman-exam/docs"
	"github.com/dxckboi/hugeman-exam/infra"
	"github.com/dxckboi/hugeman-exam/internal/handler"
	"github.com/dxckboi/hugeman-exam/internal/repo"
	"github.com/dxckboi/hugeman-exam/internal/service"
	"github.com/dxckboi/hugeman-exam/pkg/logger"
	"github.com/dxckboi/hugeman-exam/pkg/validator"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	logger.Init()
	validator.Init()
	config.Init()
	infra.InitDB()
}

func main() {
	app := gin.Default()
	docs.SwaggerInfo.Title = "Todo API"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Version = "1.0"

	app.Use(cors.Default())
	api := app.Group("/api")
	{
		router := api.Group("/todo")
		handler.NewTodoHandler(
			router,
			service.NewTodoService(
				repo.NewTodoRepo(infra.GetDB()),
			),
		)
	}

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(":8080")
}
