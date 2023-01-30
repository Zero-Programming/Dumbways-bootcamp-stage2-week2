package housesdto

import (
	"housy/models"

	"gorm.io/datatypes"
)

type HouseResponse struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	City      models.CityResponse `json:"city"`
	Address   string              `json:"address" `
	Price     int                 `json:"price"`
	TypeRent  string              `json:"type_rent"`
	Amenities datatypes.JSON      `json:"amenities" `
	Bedroom   int                 `json:"bedroom" `
	Bathroom  int                 `json:"bathroom" `
	Image     datatypes.JSON      `json:"image"`
}
