package todo_handler

import (
	"go-echo-demo/internal/dto"
	"go-echo-demo/internal/middleware"
	todo_service "go-echo-demo/internal/service/todo"
	"go-echo-demo/pkg/response"

	"github.com/labstack/echo/v5"
)

// CreateTodoController godoc
// @Summary 创建待办事项
// @Description 创建一个新的待办事项
// @Tags Todo
// @Accept json
// @Produce json
// @Param request body dto.CreateTodoDto true "待办事项信息"
// @Success 200 {object} response.Response{data=model.Todo} "成功"
// @Failure 401 {object} response.Response "未授权"
// @Router /todo [post]
// @Security ApiKeyAuth
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

// ListTodoController godoc
// @Summary 获取待办事项列表
// @Description 获取当前用户的待办事项列表，支持分页
// @Tags Todo
// @Accept json
// @Produce json
// @Param request query dto.GetTodoListDto true "分页信息"
// @Success 200 {object} response.Response "成功"
// @Failure 401 {object} response.Response "未授权"
// @Router /todo [get]
// @Security ApiKeyAuth
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

// UpdateTodoController godoc
// @Summary 更新待办事项
// @Description 更新现有脚本的状态或标题
// @Tags Todo
// @Accept json
// @Produce json
// @Param request body dto.UpdateTodoDto true "更新信息"
// @Success 200 {object} response.Response "成功"
// @Failure 401 {object} response.Response "未授权"
// @Router /todo [put]
// @Security ApiKeyAuth
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

// DeleteTodoController godoc
// @Summary 删除待办事项
// @Description 根据 ID 删除待办事项
// @Tags Todo
// @Accept json
// @Produce json
// @Param request body dto.DeleteTodoDto true "删除信息"
// @Success 200 {object} response.Response "成功"
// @Failure 401 {object} response.Response "未授权"
// @Router /todo [delete]
// @Security ApiKeyAuth
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
