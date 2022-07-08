package search

import "github.com/diyliv/anime/internal/models"

type UseCase interface {
	Search(string) (*models.AnimeSearchResult, error)
}
