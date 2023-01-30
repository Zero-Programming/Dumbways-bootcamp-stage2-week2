package models

type Transaction struct {
	CheckIn       string `json:"check_in"`
	CheckOut      string `json:"check_out"`
	HouseId       int    `json:"house_id"`
	House         House  `json:"house"`
	UserId        int    `json:"user_id"`
	User          User   `json:"user" gorm:""`
	Total         int    `json:"total"`
	StatusPayment string `json:"status_payment"`
}
