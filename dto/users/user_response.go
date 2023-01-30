package usersdto

type UserResponse struct {
	ID         int    `json:"id"`
	Fullname   string `json:"fullname"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	ListAsRole string `json:"list_as"`
	Gendre     string `json:"gendre"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	Image      string `json:"imgae"`
}
