package models

type User struct {
	ID       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}
type UserResponse struct {
	ID    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
	Token string `json:"token" form:"token"`
}
