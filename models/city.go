package models

// User model struct
type City struct {
	ID   int    `json:"id"`
	Name string `json:"name" gorm:"type: varchar(255);primaryKey"`
}

type CityResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CityResponse) TableName() string {
	return "list_as"
}