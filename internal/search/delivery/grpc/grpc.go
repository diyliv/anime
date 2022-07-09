package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/diyliv/anime/internal/models"
	"github.com/diyliv/anime/internal/search"
	"github.com/diyliv/anime/pkg/logger"
	searchpb "github.com/diyliv/anime/proto/animeSearch"
	"github.com/golang/protobuf/ptypes/wrappers"
)

type searchService struct {
	logger   *logger.Logger
	searchUC search.UseCase
}

func NewSearchService(logger *logger.Logger, searchUC search.UseCase) *searchService {
	return &searchService{logger: logger, searchUC: searchUC}
}

func (s *searchService) Search(ctx context.Context, r *wrappers.StringValue) (*searchpb.FinSearchResp, error) {
	name := r.GetValue()

	if strings.Contains(name, " ") {
		name = strings.Replace(name, " ", "%20", -1)
	}

	resp, err := http.Get(fmt.Sprintf("https://api.jikan.moe/v4/anime?q=%s", name))
	if err != nil {
		s.logger.Error("Error while dialing with API: " + err.Error())
	}

	if resp.StatusCode == 400 {
		return nil, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error("Error while reading body: " + err.Error())
	}

	var res models.AnimeSearchResult

	if err := json.Unmarshal(body, &res); err != nil {
		s.logger.Error("Error while unmarshalling: " + err.Error())
	}

	serviceResp := make([]*searchpb.SearchResponse, 0)

	for _, v := range res.Data {
		serviceResp = append(serviceResp, &searchpb.SearchResponse{
			&searchpb.AnimeSearch{
				MalId: int64(v.MalId),
				Url:   v.URL,
				JpgImages: &searchpb.Images{
					ImageUrl:      v.Images.JPG.ImageUrl,
					SmallImageUrl: v.Images.JPG.SmallImage,
					LargeImageUrl: v.Images.JPG.LargeImage,
				},
				YoutubeTrailer: &searchpb.Trailer{
					YoutubeID:         v.Trailer.YouTubeID,
					YoutubeTrailerURL: v.Trailer.YoutubeTrailerURL,
					YoutubeTrailerImages: &searchpb.YoutubeTrailerImages{
						ImageUrl:        v.Trailer.YoutubeTrailerImages.ImageUrl,
						SmallImageUrl:   v.Trailer.YoutubeTrailerImages.SmallImageUrl,
						MediumImageUrl:  v.Trailer.YoutubeTrailerImages.MediumImageUrl,
						LargeImageUrl:   v.Trailer.YoutubeTrailerImages.LargeImageUrl,
						MaximumImageUrl: v.Trailer.YoutubeTrailerImages.MaximumImageUrl,
					},
				},
				Title:    v.Title,
				TitleEng: v.TitleEng,
				Type:     v.Type,
				Source:   v.Source,
				Episodes: int64(v.Episodes),
				Status:   v.Status,
				Rating:   v.Rating,
				Score:    float32(v.Score),
				Year:     int64(v.Year)}})
	}
	return &searchpb.FinSearchResp{Finresp: serviceResp}, nil
}
