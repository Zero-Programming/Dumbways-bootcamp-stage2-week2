package repositories

import (
	"housy/models"

	"gorm.io/gorm"
)

type HouseRepository interface {
	FindHouses() ([]models.House, error)
	GetHouse(ID int) (models.House, error)
	CreateHouse(house models.House) (models.House, error)
}

func RepositoryHouse(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHouses() ([]models.House, error) {
	var houses []models.House
	err := r.db.Find(&houses).Error

	return houses, err
}

func (r *repository) GetHouse(ID int) (models.House, error) {
	var house models.House
	err := r.db.First(&house, ID).Error

	return house, err
}

// Write this code
func (r *repository) CreateHouse(house models.House) (models.House, error) {
	err := r.db.Create(&house).Error // Using Create method

	return house, err
}
