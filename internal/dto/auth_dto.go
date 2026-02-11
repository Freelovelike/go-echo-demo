package dto
type LoginAndRegisterDto struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=6,max=32"`
}
