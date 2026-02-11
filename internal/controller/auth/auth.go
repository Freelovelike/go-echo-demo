package auth_handler

import (
	"go-echo-demo/internal/dto"
	authService "go-echo-demo/internal/service/auth"
	"go-echo-demo/pkg/response"

	"github.com/labstack/echo/v5"
)

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
