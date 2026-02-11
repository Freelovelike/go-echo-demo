package todo_handler

import (
	"go-echo-demo/internal/dto"
	"go-echo-demo/internal/middleware"
	todo_service "go-echo-demo/internal/service/todo"
	"go-echo-demo/pkg/response"

	"github.com/labstack/echo/v5"
)

func CreateTodoController(c *echo.Context) error {
	userId := middleware.MustGetUserID(c)
	var dto dto.CreateTodoDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	todo, err := todo_service.CreateTodoService(c, dto.Title, userId)
	if err != nil {
		return response.ResErr(c, response.CodeDBError, err.Error())
	}
	return response.ResOK(c, todo)
}

func ListTodoController(c *echo.Context) error {
	userId := middleware.MustGetUserID(c)
	var dto dto.GetTodoListDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	list, err := todo_service.ListTodoService(c, userId, dto)
	if err != nil {
		return response.ResErr(c, response.CodeDBError, err.Error())
	}
	return response.ResOK(c, list)
}

func UpdateTodoController(c *echo.Context) error {
	userId := middleware.MustGetUserID(c)
	var dto dto.UpdateTodoDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	err := todo_service.UpdateTodoService(c, userId, dto)
	if err != nil {
		return response.ResErr(c, response.CodeDBError, err.Error())
	}
	return response.ResOK(c, nil)
}

func DeleteTodoController(c *echo.Context) error {
	userId := middleware.MustGetUserID(c)
	var dto dto.DeleteTodoDto
	if err := c.Bind(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, "无法解析请求参数")
	}
	if err := c.Validate(&dto); err != nil {
		return response.ResErr(c, response.CodeInvalidParam, err.Error())
	}
	err := todo_service.DeleteTodoService(c, userId, dto.ID)
	if err != nil {
		return response.ResErr(c, response.CodeDBError, err.Error())
	}
	return response.ResOK(c, nil)
}
