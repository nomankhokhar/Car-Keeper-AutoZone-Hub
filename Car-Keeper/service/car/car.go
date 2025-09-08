package car

import (
	"context"

	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/store"
)

type CarService struct {
	store store.CarStoreInterface
}

func NewCarService(store store.CarStoreInterface) *CarService {
	return &CarService{
		store: store,
	}
}

func (cs *CarService) GetCarByID(ctx context.Context, id string) (*models.Car, error) {
	car, err := cs.store.GetCarById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) GetCarsByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	cars, err := s.store.GetCarByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func (s *CarService) CreateCar(ctx context.Context, carReq *models.CarRequest) (*models.Car, error) {
	if err := models.ValidateCarRequest(*carReq); err != nil {
		return nil, err
	}

	car, err := s.store.CreateCar(ctx, carReq)
	if err != nil {
		return nil, err
	}
	return &car, nil
}

func (s *CarService) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (*models.Car, error) {
	if err := models.ValidateCarRequest(*carReq); err != nil {
		return nil, err
	}

	updatedCar, err := s.store.UpdateCar(ctx, id, carReq)

	if err != nil {
		return nil, err
	}

	return &updatedCar, nil
}

func (s *CarService) DeleteCar(ctx context.Context, id string) (*models.Car, error) {
	deleteCar, err := s.store.DeleteCar(ctx, id)

	if err != nil {
		return nil, err
	}

	return &deleteCar, nil
}
