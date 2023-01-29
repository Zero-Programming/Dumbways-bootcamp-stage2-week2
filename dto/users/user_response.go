package usersdto

import "housy/models"

type UserResponse struct {
	ID       int                   `json:"id"`
	Fullname string                `json:"fullname"`
	Username string                `json:"username"`
	Email    string                `json:"email"`
	Password string                `json:"password"`
	ListAs   models.ListAsResponse `json:"list_as"`
	Gendre   string                `json:"gendre"`
	Phone    string                `json:"phone"`
	Address  string                `json:"address"`
	Image    string                `json:"imgae"`
}
