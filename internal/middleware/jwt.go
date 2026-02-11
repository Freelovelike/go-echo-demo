package middleware

import (
	"go-echo-demo/pkg"
	"net/http"
	"strings"

	"github.com/labstack/echo/v5"
)

// JWTAuth 是一个 JWT 认证中间件
// 它会从请求头中提取 token，验证并将 userid 存储到 context 中
func JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// 从 Authorization header 中获取 token
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "缺少认证令牌",
				})
			}

			// 检查是否是 Bearer token 格式
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "认证令牌格式错误",
				})
			}

			tokenString := parts[1]

			// 解析并验证 token
			claims, err := pkg.ParseToken(tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"code":    401,
					"message": "认证令牌无效或已过期",
					"error":   err.Error(),
				})
			}

			// 将 userid 存储到 context 中，供后续 handler 使用
			c.Set("userid", claims.ID)

			// 继续执行下一个 handler
			return next(c)
		}
	}
}

// GetUserID 从 context 中获取当前用户的 ID
// 返回 userid 和是否成功获取的布尔值
// 用于需要手动处理错误的场景
func GetUserID(c *echo.Context) (uint, bool) {
	userid := c.Get("userid")
	if userid == nil {
		return 0, false
	}

	id, ok := userid.(uint)
	return id, ok
}

// MustGetUserID 从 context 中获取当前用户的 ID
// 仅用于已经过 JWTAuth 中间件保护的路由
// 在这种情况下，userid 必然存在，无需错误检查
func MustGetUserID(c *echo.Context) uint {
	userid := c.Get("userid")
	return userid.(uint)
}
