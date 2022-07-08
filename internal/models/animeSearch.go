package models

type AnimeSearchResult struct {
	Data []struct {
		MalId  int    `json:"mal_id"`
		URL    string `json:"url"`
		Images struct {
			JPG struct {
				ImageUrl   string `json:"image_url"`
				SmallImage string `json:"small_image_url"`
				LargeImage string `json:"large_image_url"`
			} `json:"jpg"`
		} `json:"images"`
		Trailer struct {
			YouTubeID            string `json:"youtube_id"`
			YoutubeTrailerURL    string `json:"url"`
			YoutubeTrailerImages struct {
				ImageUrl        string `json:"image_url"`
				SmallImageUrl   string `json:"small_image_url"`
				MediumImageUrl  string `json:"medium_image_url"`
				LargeImageUrl   string `json:"large_image_url"`
				MaximumImageUrl string `json:"maximum_image_url"`
			} `json:"images"`
		} `json:"trailer"`
		Title    string  `json:"title"`
		TitleEng string  `json:"title_english"`
		TitleJPN string  `json:"title_japanese"`
		Type     string  `json:"type"`
		Source   string  `json:"source"`
		Episodes int     `json:"episodes"`
		Status   string  `json:"status"`
		Rating   string  `json:"rating"`
		Score    float64 `json:"score"`
		Year     int     `json:"year"`
	} `json:"data"`
}
