package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

// CustomValidator 实现了 Echo 的 Validator 接口
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate 处理结构体校验
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// 返回具体的校验错误
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
