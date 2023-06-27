package entity

type Anime struct {
	ID          int     `json:"ID"`
	TitleRus    string  `json:"titleRus"`
	TitleEng    string  `json:"titleEng"`
	ReleaseDate int     `json:"releaseDate"`
	Duration    int     `json:"duration"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Views       int     `json:"views"`
	Description string  `json:"description"`
	Genres      []int   `json:"genres"`
}

type AnimeViewModel struct {
	TitleRus    string  `json:"titleRus"`
	TitleEng    string  `json:"titleEng"`
	ReleaseDate int     `json:"releaseDate"`
	Duration    int     `json:"duration"`
	Director    string  `json:"director"`
	Rating      float64 `json:"rating"`
	Views       int     `json:"views"`
	Description string  `json:"description"`
	Genres      []int   `json:"genres"`
}

type AnimeGenre struct {
	AnimeID int `json:"animeID"`
	GenreID int `json:"genreID"`
}
