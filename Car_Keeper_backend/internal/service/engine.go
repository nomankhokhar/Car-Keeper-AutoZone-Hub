package service

import (
	"Car_Keeper/internal/models"
	"Car_Keeper/internal/repository"
	"context"

	"go.opentelemetry.io/otel"
)

type EngineService interface {
	GetEngineByID(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error)
	UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error)
	DeleteEngine(ctx context.Context, id string) error
}

type engineService struct {
	repo repository.EngineRepository
}

func NewEngineService(repo repository.EngineRepository) EngineService {
	return &engineService{repo: repo}
}
func (s *engineService) GetEngineByID(ctx context.Context, id string) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	ctx, span := trace.Start(ctx, "GetEngineByID-Service")
	defer span.End()

	return s.repo.GetEngineByID(ctx, id)
}

func (s *engineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	ctx, span := trace.Start(ctx, "CreateEngine-Service")
	defer span.End()

	engine := &models.Engine{
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}
	if err := s.repo.CreateEngine(ctx, engine); err != nil {
		return nil, err
	}
	return engine, nil
}

func (s *engineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	trace := otel.Tracer("EngineService")
	ctx, span := trace.Start(ctx, "UpdateEngine-Service")
	defer span.End()

	engine, err := s.repo.GetEngineByID(ctx, id)
	if err != nil {
		return nil, err
	}

	engine.Displacement = engineReq.Displacement
	engine.NoOfCylinders = engineReq.NoOfCylinders
	engine.CarRange = engineReq.CarRange

	if err := s.repo.UpdateEngine(ctx, engine); err != nil {
		return nil, err
	}
	return engine, nil
}

func (s *engineService) DeleteEngine(ctx context.Context, engineID string) error {
	trace := otel.Tracer("EngineService")
	ctx, span := trace.Start(ctx, "DeleteEngine-Service")
	defer span.End()

	return s.repo.DeleteEngine(ctx, engineID)
}
