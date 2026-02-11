package route

import (
	authHandler "go-echo-demo/internal/controller/auth"

	"github.com/labstack/echo/v5"
)

func SetupAuthRoutes(e *echo.Group) {
	authPath := e.Group("/auth")
	authPath.POST("/login", authHandler.LoginController)
	authPath.POST("/register", authHandler.RegisterController)
}
