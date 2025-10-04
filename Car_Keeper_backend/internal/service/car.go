package service

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
	"context"

	"go.opentelemetry.io/otel"
)

type CarService interface {
	GetCarByID(id string) (*models.Car, error)
	GetCarByBrand(brand string) ([]models.Car, error)
	CreateCar(car *models.CarRequest) error
	UpdateCar(id string, car *models.CarRequest) error
	DeleteCar(id string) error
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) GetCarByID(id string) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(context.Background(), "GetCarByID-Service")
	defer span.End()

	return s.repo.GetCarByID(id)
}
func (s *carService) GetCarByBrand(brand string) ([]models.Car, error) {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(context.Background(), "GetCarByBrand-Service")
	defer span.End()

	return s.repo.GetCarByBrand(brand)
}

func (s *carService) CreateCar(carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(context.Background(), "CreateCar-Service")
	defer span.End()

	return s.repo.CreateCar(carReq)
}

func (s *carService) UpdateCar(id string, carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(context.Background(), "UpdateCar-Service")
	defer span.End()

	// For simplicity, we'll just call CreateCar for now.
	// In a real application, you'd implement an UpdateCar method in the repository.
	return s.repo.UpdateCar(id, carReq)
}

func (s *carService) DeleteCar(id string) error {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(context.Background(), "DeleteCar-Service")
	defer span.End()

	// For simplicity, we'll just call CreateCar for now.
	// In a real application, you'd implement a DeleteCar method in the repository.
	return s.repo.DeleteCar(id)
}
