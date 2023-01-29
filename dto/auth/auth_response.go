package authdto

type SignUpResponse struct {
	Username string `json:"username" gorm:"type: varchar(225)"`
	Message  string `json:"message" gorm:"type: varchar(225)"`
}

type SignInResponse struct {
	Username string `json:"username" gorm:"type: varchar(225)"`
	Token    string `json:"token" gorm:"type: varchar(225)"`
}
