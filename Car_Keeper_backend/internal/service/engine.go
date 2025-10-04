package service

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
	"context"

	"go.opentelemetry.io/otel"
)

type EngineService interface {
	GetEngineByID(id string) (*models.Engine, error)
	CreateEngine(engineReq *models.EngineRequest) (*models.Engine, error)
	UpdateEngine(id string, engineReq *models.EngineRequest) (*models.Engine, error)
	DeleteEngine(id string) error
}

type engineService struct {
	repo repository.EngineRepository
}

func NewEngineService(repo repository.EngineRepository) EngineService {
	return &engineService{repo: repo}
}
func (s *engineService) GetEngineByID(id string) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	_, span := trace.Start(context.Background(), "GetEngineByID-Service")
	defer span.End()

	return s.repo.GetEngineByID(id)
}

func (s *engineService) CreateEngine(engineReq *models.EngineRequest) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	_, span := trace.Start(context.Background(), "CreateEngine-Service")
	defer span.End()

	engine := &models.Engine{
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	if err := s.repo.CreateEngine(engine); err != nil {
		return nil, err
	}
	return engine, nil
}

func (s *engineService) UpdateEngine(id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	_, span := trace.Start(context.Background(), "UpdateEngine-Service")
	defer span.End()

	engine, err := s.repo.GetEngineByID(id)
	if err != nil {
		return nil, err
	}

	engine.Displacement = engineReq.Displacement
	engine.NoOfCylinders = engineReq.NoOfCylinders
	engine.CarRange = engineReq.CarRange

	if err := s.repo.UpdateEngine(engine); err != nil {
		return nil, err
	}
	return engine, nil
}

func (s *engineService) DeleteEngine(engineID string) error {
	trace := otel.Tracer("EngineService")
	_, span := trace.Start(context.Background(), "DeleteEngine-Service")
	defer span.End()

	return s.repo.DeleteEngine(engineID)
}
