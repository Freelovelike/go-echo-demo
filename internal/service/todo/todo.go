package todo_service

import (
	"fmt"
	"go-echo-demo/internal/dto"
	"go-echo-demo/internal/model"
	"go-echo-demo/internal/vo"
	"go-echo-demo/pkg/db"
	"go-echo-demo/pkg/utils"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func CreateTodoService(c *echo.Context, title string, userId uint) (model.Todo, error) {

	todo := model.Todo{
		Title:  title,
		UserID: userId,
	}
	result := db.DB.WithContext(c.Request().Context()).Create(&todo)
	if result.Error != nil {
		return todo, result.Error
	}
	return todo, nil
}

func ListTodoService(c *echo.Context, userId uint, dto dto.GetTodoListDto) (vo.PageResult[model.Todo], error) {
	var list []model.Todo
	var total int64

	query := db.DB.WithContext(c.Request().Context()).Model(&model.Todo{}).Where("user_id = ?", userId)

	err := query.Session(&gorm.Session{Context: c.Request().Context()}).Count(&total).Error
	if err != nil {
		return vo.PageResult[model.Todo]{}, err
	}

	err = query.Session(&gorm.Session{Context: c.Request().Context()}).Scopes(utils.Paginate(dto.Page, dto.Limit)).Find(&list).Error

	if err != nil {
		return vo.PageResult[model.Todo]{}, err
	}
	return vo.PageResult[model.Todo]{
		List:  list,
		Total: int(total),
		Page:  dto.Page,
		Limit: dto.Limit,
	}, nil
}

func UpdateTodoService(c *echo.Context, userId uint, dto dto.UpdateTodoDto) error {
	var todo model.Todo
	result := db.DB.WithContext(c.Request().Context()).Where("id = ? and user_id = ?", dto.ID, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}
	todo.Title = dto.Title
	todo.Completed = dto.Completed
	result = db.DB.WithContext(c.Request().Context()).Save(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTodoService(c *echo.Context, userId uint, id uint) error {
	var todo model.Todo

	fmt.Println("userId", userId)
	fmt.Println("id", id)

	result := db.DB.WithContext(c.Request().Context()).Where("id = ? and user_id = ?", id, userId).First(&todo)
	if result.Error != nil {
		return result.Error
	}
	result = db.DB.WithContext(c.Request().Context()).Delete(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
