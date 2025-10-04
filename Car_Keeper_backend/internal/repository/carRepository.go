package repository

import (
	"Car_Keeper/internal/models"

	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

type CarRepository interface {
	GetCarByID(id string) (*models.Car, error)
}

func (r *carRepository) GetCarByID(id string) (*models.Car, error) {
	var car models.Car
	if err := r.db.Preload("Engine").First(&car, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &car, nil
}
