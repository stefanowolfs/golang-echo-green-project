package main

import (
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
	"green-echo/internal/app/green/connector"
	"green-echo/internal/app/green/handler"
	"green-echo/internal/app/green/repository"
	"green-echo/internal/app/green/service"
)

// @title Swagger Green API
// @version 1.0
// @description Green API
// @title Green API

// @host 127.0.0.1:8585
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json

func main() {
	r := handler.New()
	DB := connector.NewGormDB()
	ur := repository.NewUserPostgresRepository(DB)
	us := service.NewUserService(ur)
	h := handler.NewHandler(us)

	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/api")
	h.Register(v1)

	r.Logger.Fatal(r.Start("127.0.0.1:8585"))
}
