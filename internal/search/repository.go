package search

import "github.com/diyliv/anime/internal/models"

type PostgresRepository interface {
	Search(string) (*models.AnimeSearchResult, error)
}
