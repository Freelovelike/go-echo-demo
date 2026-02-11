package response

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type Response struct {
	Code ResCode `json:"code"`
	Msg  string  `json:"msg"`
	Data any     `json:"data"`
}

func ResOK(c *echo.Context, data any) error {
	return c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	})
}

func ResErr(c *echo.Context, code ResCode, msg ...string) error {
	var msgStr string
	if len(msg) > 0 {
		msgStr = msg[0]
	} else {
		msgStr = code.Msg()
	}
	return c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msgStr,
		Data: nil,
	})
}
