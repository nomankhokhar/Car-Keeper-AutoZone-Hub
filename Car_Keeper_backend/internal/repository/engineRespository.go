package repository

import (
	"Car_Keeper/internal/models"

	"gorm.io/gorm"
)

type engineRepository struct {
	db *gorm.DB
}

type EngineRepository interface {
	GetEngineByID(id string) (*models.Engine, error)
	CreateEngine(engine *models.Engine) error
}

func NewEngineRepository(db *gorm.DB) EngineRepository {
	return &engineRepository{db: db}
}

func (r *engineRepository) GetEngineByID(id string) (*models.Engine, error) {
	var engine models.Engine
	if err := r.db.First(&engine, "engine_id = ?", id).Error; err != nil {
		return nil, err
	}
	return &engine, nil
}

func (r *engineRepository) CreateEngine(engine *models.Engine) error {
	return r.db.Create(engine).Error
}
