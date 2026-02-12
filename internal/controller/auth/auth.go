package auth_handler

import (
	"go-echo-demo/internal/dto"
	authService "go-echo-demo/internal/service/auth"
	"go-echo-demo/pkg/response"

	"github.com/labstack/echo/v5"
)

// LoginController godoc
// @Summary 用户登录
// @Description 通过用户名和密码获取 JWT Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginAndRegisterDto true "登录信息"
// @Success 200 {object} response.Response{data=vo.LoginAndRegisterVo} "成功"
// @Failure 400 {object} response.Response "参数错误"
// @Router /auth/login [post]
func LoginController(c *echo.Context) error {
	var dto dto.LoginAndRegisterDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	data, err := authService.UserLoginService(c, dto)
	if err != nil {
		return response.ResErr(c, response.CodeUserNotExist)
	}

	return response.ResOK(c, data)
}

// RegisterController godoc
// @Summary 用户注册
// @Description 注册新用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.LoginAndRegisterDto true "注册信息"
// @Success 200 {object} response.Response{data=model.User} "成功"
// @Failure 400 {object} response.Response "参数错误"
// @Router /auth/register [post]
func RegisterController(c *echo.Context) error {
	var dto dto.LoginAndRegisterDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	user, err := authService.UserRegisterService(dto)
	if err != nil {
		return response.ResErr(c, response.CodeUserExist)
	}
	return response.ResOK(c, user)
}
