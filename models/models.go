package models

type User struct {
	ID       int32  `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Address  string `json:"address"  binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
