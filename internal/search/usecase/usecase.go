package usecase

import (
	"github.com/diyliv/anime/internal/models"
	"github.com/diyliv/anime/internal/search"
)

type searchUC struct {
	searchRepo search.PostgresRepository
}

func NewSearchUC(searchRepo search.PostgresRepository) *searchUC {
	return &searchUC{searchRepo: searchRepo}
}

func (u *searchUC) Search(name string) (*models.AnimeSearchResult, error) {
	return u.searchRepo.Search(name)
}
