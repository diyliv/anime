package repository

import (
	"database/sql"

	"github.com/diyliv/anime/internal/models"
	"go.uber.org/zap"
)

type animeRepo struct {
	psqlDB sql.DB
	logger *zap.Logger
}

func NewAnimeRepo(psqlDB sql.DB, logger *zap.Logger) *animeRepo {
	return &animeRepo{psqlDB: psqlDB, logger: logger}
}

func (r *animeRepo) Search(name string) (*models.AnimeSearchResult, error) {
	return nil, nil
}
