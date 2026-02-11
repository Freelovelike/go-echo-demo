package route

import (
	user_handler "go-echo-demo/internal/controller/user"

	"github.com/labstack/echo/v5"
)

func SetupUserRoutes(e *echo.Group) {
	userPath := e.Group("/user")

	// 应用 JWT 认证中间件，保护用户相关接口

	// 需要认证的用户接口
	userPath.GET("/info", user_handler.GetUserInfoController)
}
