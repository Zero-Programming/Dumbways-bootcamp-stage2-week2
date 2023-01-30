package models

import "time"

// User model struct
type User struct {
	ID         int       `json:"id"`
	Fullname   string    `json:"fullname" gorm:"type: varchar(255)"`
	Username   string    `json:"username" gorm:"type: varchar(255)"`
	Email      string    `json:"email" gorm:"type: varchar(255)"`
	Password   string    `json:"password" gorm:"type: varchar(255)"`
	ListAsRole string    `json:"-" gorm:"type: varchar(225)"`
	ListAs     ListAs    `json:"list_as" gorm:"foreignKey:ListAsRole"`
	Gender     string    `json:"gender" gorm:"type: varchar(255)"`
	Phone      string    `json:"phone" gorm:"type: varchar(255)"`
	Address    string    `json:"address" gorm:"type: varchar(225)"`
	Image      string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserProfileRespone struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	ListAs   ListAs `json:"list_as"`
	Gender   string `json:"gender"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Image    string `json:"image"`
}

func (UserProfileRespone) TableName() string {
	return "users"
}
