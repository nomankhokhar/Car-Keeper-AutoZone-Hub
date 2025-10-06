package service

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
	"context"

	"go.opentelemetry.io/otel"
)

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

type CarService interface {
	GetCarByID(ctx context.Context, id string) (*models.Car, error)
	GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error)
	CreateCar(ctx context.Context, car *models.CarRequest) error
	UpdateCar(ctx context.Context, id string, car *models.CarRequest) error
	DeleteCar(ctx context.Context, id string) error
}

func (s *carService) GetCarByID(ctx context.Context, id string) (*models.Car, error) {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "GetCarByID-Service")
	defer span.End()

	return s.repo.GetCarByID(ctx, id)
}

func (s *carService) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	tracer := otel.Tracer("CarService")
	_, span := tracer.Start(ctx, "GetCarByBrand-Service")
	defer span.End()

	return s.repo.GetCarByBrand(ctx, brand)
}

func (s *carService) CreateCar(ctx context.Context, carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "CreateCar-Service")
	defer span.End()

	return s.repo.CreateCar(ctx, carReq)
}

func (s *carService) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "UpdateCar-Service")
	defer span.End()

	// For simplicity, we'll just call CreateCar for now.
	// In a real application, you'd implement an UpdateCar method in the repository.
	return s.repo.UpdateCar(ctx, id, carReq)
}

func (s *carService) DeleteCar(ctx context.Context, id string) error {
	tracer := otel.Tracer("CarService")
	ctx, span := tracer.Start(ctx, "DeleteCar-Service")
	defer span.End()

	// For simplicity, we'll just call CreateCar for now.
	// In a real application, you'd implement a DeleteCar method in the repository.
	return s.repo.DeleteCar(ctx, id)
}
