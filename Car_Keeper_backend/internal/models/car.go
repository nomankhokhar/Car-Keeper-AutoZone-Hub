// models/car.go
package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name     string    `gorm:"not null" json:"name"`
	Year     string    `gorm:"size:4;not null" json:"year"`
	Brand    string    `gorm:"not null" json:"brand"`
	FuelType string    `gorm:"not null" json:"fuel_type"`
	Price    float64   `gorm:"not null" json:"price"`

	EngineID uuid.UUID `gorm:"type:uuid;not null" json:"engine_id"`
	Engine   Engine    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"engine"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// BeforeCreate hook to generate UUID
func (c *Car) BeforeCreate(tx *gorm.DB) error {
	if c.ID == uuid.Nil {
		c.ID = uuid.New()
	}
	return nil
}

type CarRequest struct {
	Name     string    `json:"name" binding:"required"`
	Year     string    `json:"year" binding:"required,len=4"`
	Brand    string    `json:"brand" binding:"required"`
	FuelType string    `json:"fuel_type" binding:"required,oneof=petrol diesel electric hybrid"`
	EngineID uuid.UUID `json:"engine_id" binding:"required"` // use ID, not full object
	Price    float64   `json:"price" binding:"required"`
}
