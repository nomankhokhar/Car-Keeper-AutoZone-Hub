package repository

import (
	"Car_Keeper/internal/models"
	"context"

	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

type engineRepository struct {
	db *gorm.DB
}

type EngineRepository interface {
	GetEngineByID(ctx context.Context, id string) (*models.Engine, error)
	CreateEngine(ctx context.Context, engine *models.Engine) error
	UpdateEngine(ctx context.Context, engine *models.Engine) error
	DeleteEngine(ctx context.Context, engineID string) error
}

func NewEngineRepository(db *gorm.DB) EngineRepository {
	return &engineRepository{db: db}
}

func (r *engineRepository) GetEngineByID(ctx context.Context, id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineRepository")
	ctx, span := tracer.Start(ctx, "GetEngineByID-Repository")
	defer span.End()

	var engine models.Engine
	if err := r.db.First(&engine, "engine_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &engine, nil
}

func (r *engineRepository) CreateEngine(ctx context.Context, engine *models.Engine) error {
	tracer := otel.Tracer("EngineRepository")
	ctx, span := tracer.Start(ctx, "CreateEngine-Repository")
	defer span.End()

	return r.db.Create(engine).Error
}

func (r *engineRepository) UpdateEngine(ctx context.Context, engine *models.Engine) error {
	tracer := otel.Tracer("EngineRepository")
	ctx, span := tracer.Start(ctx, "UpdateEngine-Repository")
	defer span.End()

	return r.db.Save(engine).Error
}

func (r *engineRepository) DeleteEngine(ctx context.Context, engineID string) error {
	tracer := otel.Tracer("EngineRepository")
	ctx, span := tracer.Start(ctx, "DeleteEngine-Repository")
	defer span.End()

	return r.db.Delete(&models.Engine{}, "engine_id = ?", engineID).Error
}
