package user_handler

import (
	"go-echo-demo/internal/middleware"
	user_service "go-echo-demo/internal/service/user"
	"go-echo-demo/pkg/response"

	"github.com/labstack/echo/v5"
)

// GetUserInfoController 获取用户信息
// 该接口受 JWT 中间件保护,只有携带有效 token 的请求才能访问
func GetUserInfoController(c *echo.Context) error {
	// 直接获取 userid，因为中间件已经验证通过

	userid := middleware.MustGetUserID(c)

	user, err := user_service.GetUserInfoService(c, userid)
	if err != nil {
		return response.ResErr(c, response.CodeDBError)
	}
	// 这里可以使用 userid 查询数据库获取用户详细信息
	// 示例返回
	return response.ResOK(c, user)
}
