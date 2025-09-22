package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nomankhokhar/Car-Keeper-AutoZone-Hub/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{
		db: db,
	}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {
	var engine models.Engine

	fmt.Println("Starting GetEngineById with ID:", id)

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Failed to begin transaction:", err)
		return engine, err
	}
	fmt.Println("Transaction started")

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	// Execute query
	fmt.Println("Executing SELECT query...")
	err = tx.QueryRowContext(
		ctx,
		"SELECT id, displacement, noofcylinders, carrange FROM engine WHERE id = $1",
		id,
	).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCyclinders, &engine.CarRange)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No engine found with ID:", id)
			return engine, nil
		}
		fmt.Println("Query execution error:", err)
		return engine, err
	}

	fmt.Println("Fetched engine successfully:", engine)
	return engine, nil
}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {
	fmt.Println("Starting CreateEngine with request:", engineReq)

	tx, err := e.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Println("Failed to begin transaction:", err)
		return models.Engine{}, err
	}
	fmt.Println("Transaction started")

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	engineID := uuid.New()
	fmt.Println("Generated new Engine ID:", engineID)

	res, err := tx.ExecContext(
		ctx,
		"INSERT INTO engine (id, displacement, noofcylinders, carrange) VALUES ($1, $2, $3, $4)",
		engineID, engineReq.Displacement, engineReq.NoOfCyclinders, engineReq.CarRange,
	)
	if err != nil {
		fmt.Println("Error executing INSERT query:", err)
		return models.Engine{}, err
	}
	fmt.Println("INSERT query executed successfully")

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("Error checking rows affected:", err)
		return models.Engine{}, err
	}
	fmt.Println("Rows affected by INSERT:", rows)

	if rows == 0 {
		fmt.Println("No rows inserted into engine table")
		return models.Engine{}, fmt.Errorf("no rows were inserted into engine table")
	}

	engine := models.Engine{
		EngineID:       engineID,
		Displacement:   engineReq.Displacement,
		NoOfCyclinders: engineReq.NoOfCyclinders,
		CarRange:       engineReq.CarRange,
	}
	fmt.Println("Engine created successfully:", engine)

	return engine, nil
}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {
	engineID, err := uuid.Parse(id)

	if err != nil {
		return models.Engine{}, err
	}

	tx, err := e.db.BeginTx(ctx, nil)
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	result, err := tx.ExecContext(ctx, "UPDATE engine SET displacement = $1, noofcylinders = $2, carrange = $3 WHERE id = $4", engineReq.Displacement, engineReq.NoOfCyclinders, engineReq.CarRange, engineID)
	if err != nil {
		return models.Engine{}, err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return models.Engine{}, err
	}

	if rowAffected == 0 {
		return models.Engine{}, errors.New("engine not found")
	}

	engine := models.Engine{
		EngineID:       engineID,
		Displacement:   engineReq.Displacement,
		NoOfCyclinders: engineReq.NoOfCyclinders,
		CarRange:       engineReq.CarRange,
	}

	return engine, nil
}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	engineID, err := uuid.Parse(id)

	if err != nil {
		return models.Engine{}, err
	}

	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, displacement, noofcylinders, carrange FROM engine WHERE id = $1", engineID).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCyclinders, &engine.CarRange)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}
		return engine, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM engine WHERE id = $1", engineID)

	if err != nil {
		return models.Engine{}, err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return models.Engine{}, err
	}

	if rowAffected == 0 {
		return models.Engine{}, errors.New("engine not found")
	}

	return engine, nil
}
