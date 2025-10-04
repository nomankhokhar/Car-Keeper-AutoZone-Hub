package database

import (
	"Car_Keeper/internal/config"
	"Car_Keeper/internal/models"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	// Migrate Engine first, then Car
	return db.AutoMigrate(
		&models.Engine{},
		&models.Car{},
	)
}

func InitDatabase(cfg *config.Config) (*gorm.DB, error) {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		return nil, err
	}

	// Auto-migrate models
	if err := AutoMigrate(db); err != nil {
		return nil, err
	}

	// Insert dummy data after migration
	if err := seedData(db); err != nil {
		log.Printf("⚠️ Warning: Failed to seed data: %v", err)
	}

	return db, nil
}

func seedData(db *gorm.DB) error {
	// Insert Engines first
	engines := []models.Engine{
		{EngineID: uuid.MustParse("e1f86b1a-0873-4c19-bae2-fc60329d0140"), Displacement: 2000, NoOfCylinders: 4, CarRange: 600},
		{EngineID: uuid.MustParse("f4a9c66b-8e38-419b-93c4-215d5cefb318"), Displacement: 1600, NoOfCylinders: 4, CarRange: 550},
		{EngineID: uuid.MustParse("cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c"), Displacement: 3000, NoOfCylinders: 6, CarRange: 700},
		{EngineID: uuid.MustParse("9746be12-07b7-42a3-b8ab-7d1f209b63d7"), Displacement: 1800, NoOfCylinders: 4, CarRange: 500},
	}
	for _, e := range engines {
		db.FirstOrCreate(&e, models.Engine{EngineID: e.EngineID})
	}

	// Then insert Cars
	cars := []models.Car{
		{ID: uuid.MustParse("c7c1a6d5-1ec4-4c64-a59a-8a2f6f3d2bf3"), Name: "Honda Civic", Year: "2023", Brand: "Honda", FuelType: "Gasoline", EngineID: uuid.MustParse("e1f86b1a-0873-4c19-bae2-fc60329d0140"), Price: 25000},
		{ID: uuid.MustParse("9d6a56f8-79c3-4931-a5c0-6b290c84ba2f"), Name: "Toyota Corolla", Year: "2022", Brand: "Toyota", FuelType: "Gasoline", EngineID: uuid.MustParse("f4a9c66b-8e38-419b-93c4-215d5cefb318"), Price: 22000},
		{ID: uuid.MustParse("9b9437c4-3ed1-45a5-b240-0fe3e24e0e4e"), Name: "Ford Mustang", Year: "2024", Brand: "Ford", FuelType: "Gasoline", EngineID: uuid.MustParse("cc2c2a7d-2e21-4f59-b7b8-bd9e5e4cf04c"), Price: 40000},
		{ID: uuid.MustParse("5e9df51a-8d7a-4d84-9c58-4ccfe5c7db06"), Name: "BMW 3 Series", Year: "2023", Brand: "BMW", FuelType: "Gasoline", EngineID: uuid.MustParse("9746be12-07b7-42a3-b8ab-7d1f209b63d7"), Price: 35000},
	}
	for _, c := range cars {
		db.FirstOrCreate(&c, models.Car{ID: c.ID})
	}

	return nil
}
