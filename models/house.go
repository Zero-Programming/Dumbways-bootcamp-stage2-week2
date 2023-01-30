package models

import (
	"time"

	"gorm.io/datatypes"
)

// House model struct
type House struct {
	ID        int            `json:"id" gorm:"primaryKey:autoIncrement"`
	Name      string         `json:"name" gorm:"type: varchar(255)"`
	CityName  string         `json:"-" gorm:"type: varchar(255)"`
	City      City           `json:"city" gorm:"foreignKey:CityName"`
	Address   string         `json:"address" gorm:"type: text"`
	Price     int            `json:"price" gorm:"type: int"`
	TypeRent  string         `json:"type_rent" gorm:"type: varchar(255)"`
	Amenities datatypes.JSON `json:"amenities" gorm:"type: json"`
	Bedroom   int            `json:"bedroom" gorm:"type: int"`
	Bathroom  int            `json:"bathroom" gorm:"type: int"`
	Image     datatypes.JSON `json:"image" gorm:"type: json"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type HouseResponse struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	City      City           `json:"city"`
	Address   string         `json:"address" `
	Price     int            `json:"price"`
	TypeRent  string         `json:"type_rent"`
	Amenities datatypes.JSON `json:"amenities" `
	Bedroom   int            `json:"bedroom" `
	Bathroom  int            `json:"bathroom" `
	Image     datatypes.JSON `json:"image"`
}

func (HouseResponse) TableName() string {
	return "houses"
}
