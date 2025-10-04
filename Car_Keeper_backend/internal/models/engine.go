// models/engine.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Engine struct {
	EngineID      uuid.UUID `gorm:"type:uuid;primaryKey;column:engine_id" json:"engine_id"`
	Displacement  int64     `gorm:"not null" json:"displacement"`
	NoOfCylinders int64     `gorm:"not null;column:no_of_cylinders" json:"no_of_cylinders"`
	CarRange      int64     `gorm:"not null" json:"car_range"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook
func (e *Engine) BeforeCreate(tx *gorm.DB) error {
	if e.EngineID == uuid.Nil {
		e.EngineID = uuid.New()
	}
	return nil
}

// TableName override
func (Engine) TableName() string {
	return "engines"
}

type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"no_of_cylinders"`
	CarRange      int64 `json:"car_range"`
}
