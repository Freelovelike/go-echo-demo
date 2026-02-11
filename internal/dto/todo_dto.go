package dto
type CreateTodoDto struct {
	Title string `json:"title" validate:"required,min=1,max=255"`
}

type GetTodoListDto struct {
	Page  int `json:"page" validate:"omitempty,min=1"`
	Limit int `json:"limit" validate:"omitempty,min=1,max=100"`
}

type UpdateTodoDto struct {
	ID        uint   `json:"id" validate:"required"`
	Title     string `json:"title" validate:"required,min=1,max=255"`
	Completed bool   `json:"completed"`
}

type DeleteTodoDto struct {
	ID uint `json:"id" validate:"required"`
}
