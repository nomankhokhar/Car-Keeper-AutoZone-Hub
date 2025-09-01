package engine

import (
	"context"
	"database/sql"

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

}

func (e EngineStore) CreateEngine(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) UpdateEngine(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {

}
