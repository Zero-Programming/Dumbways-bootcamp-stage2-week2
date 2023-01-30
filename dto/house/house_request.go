package housesdto

import "gorm.io/datatypes"

type HouseRequest struct {
	Name      string         `json:"name" gorm:"type: varchar(225)" validate:"required"`
	City      string         `json:"city_name" gorm:"type: varchar(255)" validate:"required"`
	Address   string         `json:"address" gorm:"type: text" validate:"required"`
	Price     int            `json:"price" gorm:"type: int" validate:"required"`
	TypeRent  string         `json:"type_rent" gorm:"type: varchar(225)" validate:"required"`
	Amenities datatypes.JSON `json:"amenities" gorm:"type: JSON" validate:"required"`
	Bedroom   int            `json:"Bedroom" gorm:"type: int" validate:"required"`
	Bathroom  int            `json:"Bathroom" gorm:"type: int" validate:"required"`
	Image     string         `json:"image" gorm:"type: varchar(255)"`
}
