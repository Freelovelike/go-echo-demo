package route

import (
	todo_handler "go-echo-demo/internal/controller/todo"

	"github.com/labstack/echo/v5"
)

func SetupTodoRoutes(e *echo.Group) {
	todoPath := e.Group("/todo")
	todoPath.POST("/create", todo_handler.CreateTodoController)
	todoPath.GET("/list", todo_handler.ListTodoController)
	todoPath.PUT("/update", todo_handler.UpdateTodoController)
	todoPath.DELETE("/delete", todo_handler.DeleteTodoController)
}
