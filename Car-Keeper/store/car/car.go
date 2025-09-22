package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var Car models.Car

	query := `
		SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id,
		       c.price, c.created_at, c.updated_at,
		       e.displacement, e.noofcylinders, e.carrange
		FROM car c
		LEFT JOIN engine e ON c.engine_id = e.id
		WHERE c.id=$1
	`

	row := s.db.QueryRowContext(ctx, query, id)

	// ⚠️ Removed duplicate scan for EngineID
	err := row.Scan(
		&Car.ID,
		&Car.Name,
		&Car.Year,
		&Car.Brand,
		&Car.FuelType,
		&Car.Engine.EngineID,
		&Car.Price,
		&Car.CreatedAt,
		&Car.UpdatedAt,
		&Car.Engine.Displacement,
		&Car.Engine.NoOfCyclinders,
		&Car.Engine.CarRange,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return Car, nil
		}
		return Car, err
	}

	return Car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car

	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.noofcylinders, e.carrange FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE c.brand=$1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE brand=$1`
	}

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var car models.Car
		if isEngine {
			err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt, &car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCyclinders, &car.Engine.CarRange)
			if err != nil {
				return nil, err
			}
			car.Engine = car.Engine
		} else {
			err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.Engine.EngineID, &car.Price, &car.CreatedAt, &car.UpdatedAt)
			if err != nil {
				return nil, err
			}
		}
		cars = append(cars, car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var engineID uuid.UUID
	err := s.db.QueryRowContext(ctx, "SELECT id FROM engine WHERE id=$1", carReq.Engine.EngineID).Scan(&engineID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("engine not found")
		}
		return createdCar, err
	}

	carID := uuid.New()

	createdAt := time.Now()
	updatedAt := time.Now()

	newCar := models.Car{
		ID:       carID,
		Name:     carReq.Name,
		Year:     carReq.Year,
		Brand:    carReq.Brand,
		FuelType: carReq.FuelType,
		Engine: models.Engine{
			EngineID:       engineID,
			Displacement:   carReq.Engine.Displacement,
			NoOfCyclinders: carReq.Engine.NoOfCyclinders,
			CarRange:       carReq.Engine.CarRange,
		},
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	// Begin the transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query, newCar.ID, newCar.Name, newCar.Year, newCar.Brand, newCar.FuelType, newCar.Engine.EngineID, newCar.Price, newCar.CreatedAt, newCar.UpdatedAt).Scan(&createdCar.ID, &createdCar.Name, &createdCar.Year, &createdCar.Brand, &createdCar.FuelType, &createdCar.Engine.EngineID, &createdCar.Price, &createdCar.CreatedAt, &createdCar.UpdatedAt)
	if err != nil {
		return createdCar, err
	}

	return createdCar, nil
}

func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	var updateCar models.Car

	// Begin the transaction

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updateCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `UPDATE car SET name=$1, year=$2, brand=$3, fuel_type=$4, engine_id=$5, price=$6, updated_at=$7 WHERE id=$8 RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query, carReq.Name, carReq.Year, carReq.Brand, carReq.FuelType, carReq.Engine.EngineID, carReq.Price, time.Now(), id).Scan(&updateCar.ID, &updateCar.Name, &updateCar.Year, &updateCar.Brand, &updateCar.FuelType, &updateCar.Engine.EngineID, &updateCar.Price, &updateCar.CreatedAt, &updateCar.UpdatedAt)
	if err != nil {
		return updateCar, err
	}

	return updateCar, nil
}

func (s Store) DeleteCar(ctx context.Context, id string) (models.Car, error) {
	var deletedCar models.Car

	// Begin the transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return deletedCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.QueryRowContext(ctx, `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE id=$1`, id).Scan(&deletedCar.ID, &deletedCar.Name, &deletedCar.Year, &deletedCar.Brand, &deletedCar.FuelType, &deletedCar.Engine.EngineID, &deletedCar.Price, &deletedCar.CreatedAt, &deletedCar.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return deletedCar, errors.New("car not found")
		}
		return models.Car{}, err
	}

	result, err := tx.ExecContext(ctx, `DELETE FROM car WHERE id=$1`, id)
	if err != nil {
		return models.Car{}, err
	}
	if rowsAffected, err := result.RowsAffected(); err != nil {
		return models.Car{}, err
	} else if rowsAffected == 0 {
		return models.Car{}, errors.New("no rows affected")
	}

	return deletedCar, nil
}
