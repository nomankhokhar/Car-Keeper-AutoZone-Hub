package repository

import (
	"Car_Keeper/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

type CarRepository interface {
	GetCarByID(ctx context.Context, id string) (*models.Car, error)
	GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error)
	CreateCar(ctx context.Context, carReq *models.CarRequest) error
	UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) error
	DeleteCar(ctx context.Context, id string) error
}

func (r *carRepository) GetCarByID(ctx context.Context, id string) (*models.Car, error) {
	tracer := otel.Tracer("CarRepository")
	ctx, span := tracer.Start(ctx, "GetCarByID-Repository")
	defer span.End()

	// Parse string to real UUID type
	carID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID format: %w", err)
	}

	var car models.Car
	if err := r.db.Preload("Engine").First(&car, "id = ?", carID).Error; err != nil {
		return nil, err
	}
	return &car, nil
}

func (r *carRepository) GetCarByBrand(ctx context.Context, brand string) ([]models.Car, error) {
	tracer := otel.Tracer("CarRepository")
	ctx, span := tracer.Start(ctx, "GetCarByBrand-Repository")
	defer span.End()

	var cars []models.Car
	if err := r.db.Preload("Engine").Where("brand = ?", brand).Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *carRepository) CreateCar(ctx context.Context, carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarRepository")
	ctx, span := tracer.Start(ctx, "CreateCar-Repository")
	defer span.End()

	// Map CarRequest to Car model
	car := models.Car{
		Name:     carReq.Name,
		Year:     carReq.Year,
		Brand:    carReq.Brand,
		FuelType: carReq.FuelType,
		Price:    carReq.Price,
		EngineID: carReq.EngineID, // ✅ Set foreign key,
	}
	// Create the car record in the database
	return r.db.Create(&car).Error
}

func (r *carRepository) UpdateCar(ctx context.Context, carID string, carReq *models.CarRequest) error {
	tracer := otel.Tracer("CarRepository")
	ctx, span := tracer.Start(ctx, "UpdateCar-Repository")
	defer span.End()

	// Map CarRequest to Car model
	// Parse string to real UUID type
	id, err := uuid.Parse(carID)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %w", err)
	}

	car := models.Car{
		ID:       id,
		Name:     carReq.Name,
		Year:     carReq.Year,
		Brand:    carReq.Brand,
		FuelType: carReq.FuelType,
		Price:    carReq.Price,
		EngineID: carReq.EngineID, // ✅ Set foreign key,
	}
	// Create the car record in the database
	return r.db.Save(&car).Error
}

func (r *carRepository) DeleteCar(ctx context.Context, id string) error {
	tracer := otel.Tracer("CarRepository")
	ctx, span := tracer.Start(ctx, "DeleteCar-Repository")
	defer span.End()

	// Parse string to real UUID type
	carID, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid UUID format: %w", err)
	}
	return r.db.Delete(&models.Car{}, carID).Error
}
