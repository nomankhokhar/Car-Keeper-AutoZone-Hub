package store

import (
	"context"

	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
)

type CarStoreInterface interface {
	GetCarByID(ctx context.Context, id string) (models.Car, error)
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
	CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error)
	UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error)
	DeleteCar(ctx context.Context, id string) (models.Car, error)
}

type EngineStoreInterface interface {
	GetEngineById(ctx context.Context, id string) (models.Engine, error)
	CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error)
	DeleteEngine(ctx context.Context, id string) (models.Engine, error)
}
