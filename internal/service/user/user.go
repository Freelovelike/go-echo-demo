package user_service

import (
	"go-echo-demo/pkg/db"

	"go-echo-demo/internal/model"

	"github.com/labstack/echo/v5"
)

func GetUserInfoService(c *echo.Context, userid uint) (model.User, error) {
	var user model.User
	result := db.DB.WithContext(c.Request().Context()).Where("id = ?", userid).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
