package route

import (
	"go-echo-demo/internal/middleware"
	"time"

	"github.com/labstack/echo/v5"
	echoMiddleware "github.com/labstack/echo/v5/middleware"
)

func Init(e *echo.Echo) {
	basePath := e.Group("/api")
	basePath.Use(echoMiddleware.ContextTimeout(5 * time.Second))
	publicPath := basePath.Group("/public")
	SetupAuthRoutes(publicPath)

	providePath := basePath.Group("")
	providePath.Use(middleware.JWTAuth())
	SetupUserRoutes(providePath)
	SetupTodoRoutes(providePath)
}
