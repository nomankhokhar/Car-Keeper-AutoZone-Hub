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
	GetEngineByID(id string) (*models.Engine, error)
	CreateEngine(engine *models.Engine) error
	UpdateEngine(engine *models.Engine) error
	DeleteEngine(engineID string) error
}

func NewEngineRepository(db *gorm.DB) EngineRepository {
	return &engineRepository{db: db}
}

func (r *engineRepository) GetEngineByID(id string) (*models.Engine, error) {
	tracer := otel.Tracer("EngineRepository")
	_, span := tracer.Start(context.Background(), "GetEngineByID-Repository")
	defer span.End()

	var engine models.Engine
	if err := r.db.First(&engine, "engine_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &engine, nil
}

func (r *engineRepository) CreateEngine(engine *models.Engine) error {
	tracer := otel.Tracer("EngineRepository")
	_, span := tracer.Start(context.Background(), "CreateEngine-Repository")
	defer span.End()

	return r.db.Create(engine).Error
}

func (r *engineRepository) UpdateEngine(engine *models.Engine) error {
	tracer := otel.Tracer("EngineRepository")
	_, span := tracer.Start(context.Background(), "UpdateEngine-Repository")
	defer span.End()

	return r.db.Save(engine).Error
}

func (r *engineRepository) DeleteEngine(engineID string) error {
	tracer := otel.Tracer("EngineRepository")
	_, span := tracer.Start(context.Background(), "DeleteEngine-Repository")
	defer span.End()

	return r.db.Delete(&models.Engine{}, "engine_id = ?", engineID).Error
}
