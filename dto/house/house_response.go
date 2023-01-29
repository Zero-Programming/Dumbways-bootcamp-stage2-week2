package housesdto

import (
	"housy/models"
)

type HouseResponse struct {
	ID        int                 `json:"id"`
	Name      string              `json:"name"`
	City      models.CityResponse `json:"city"`
	Address   string              `json:"address" `
	Price     int                 `json:"price"`
	TypeRent  string              `json:"type_rent"`
	Amenities []string            `json:"amenities" `
	Bedroom   int                 `json:"bedroom" `
	Bathroom  int                 `json:"bathroom" `
	Image     []string            `json:"image"`
}
