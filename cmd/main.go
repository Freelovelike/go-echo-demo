package main

import (
	"go-echo-demo/internal/model"
	"go-echo-demo/internal/route"

	"go-echo-demo/pkg/db"

	"go-echo-demo/pkg/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// @title Go Echo Demo API
// @version 1.0
// @description This is a sample Todo API server.
// @host localhost:1323
// @BasePath /api
func main() {
	db.InitDB("postgres://freelove:hwc20010616@localhost:5432/postgres")
	db.InitRedis("localhost:6379")
	err := db.DB.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		panic("自动迁移失败: " + err.Error())
	}
	e := echo.New()
	// e.Use(middleware.CSRF())
	// 注册验证器
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	e.Use(middleware.RequestLogger())
	route.Init(e)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
