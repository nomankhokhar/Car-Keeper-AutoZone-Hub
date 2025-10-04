package service

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
)

type CarService interface {
	GetCarByID(id string) (*models.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) GetCarByID(id string) (*models.Car, error) {
	return s.repo.GetCarByID(id)
}
