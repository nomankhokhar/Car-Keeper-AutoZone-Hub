package engine

import (
	"context"

	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/store"
	"go.opentelemetry.io/otel"
)

type EngineService struct {
	store store.EngineStoreInterface
}

func NewEngineService(store store.EngineStoreInterface) *EngineService {
	return &EngineService{
		store: store,
	}
}

func (s *EngineService) GetEngineByID(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "GetEngineByID-Service")
	defer span.End()

	engine, err := s.store.GetEngineById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engine, nil
}

func (s *EngineService) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "CreateEngine-Service")
	defer span.End()

	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	engineCreated, err := s.store.CreateEngine(ctx, engineReq)
	if err != nil {
		return nil, err
	}
	return &engineCreated, nil
}

func (s *EngineService) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Service")
	defer span.End()

	if err := models.ValidateEngineRequest(*engineReq); err != nil {
		return nil, err
	}

	engineUpdated, err := s.store.UpdateEngine(ctx, id, engineReq)
	if err != nil {
		return nil, err
	}
	return &engineUpdated, nil
}

func (s *EngineService) DeleteEngine(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineService")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Service")
	defer span.End()

	engineDeleted, err := s.store.DeleteEngine(ctx, id)
	if err != nil {
		return nil, err
	}
	return &engineDeleted, nil
}
