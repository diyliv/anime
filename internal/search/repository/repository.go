package repository

import (
	"database/sql"

	"github.com/diyliv/anime/internal/models"
	"github.com/diyliv/anime/pkg/logger"
)

type animeRepo struct {
	psqlDB sql.DB
	logger *logger.Logger
}

func NewAnimeRepo(psqlDB sql.DB, logger *logger.Logger) *animeRepo {
	return &animeRepo{psqlDB: psqlDB, logger: logger}
}

func (r *animeRepo) Search(name string) (*models.AnimeSearchResult, error) {
	return nil, nil
}
