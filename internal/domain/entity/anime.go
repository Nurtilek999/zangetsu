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
	TitleRus    string  `json:"titleRus" validate:"required"`
	TitleEng    string  `json:"titleEng" validate:"min=1,max=100"`
	ReleaseDate int     `json:"releaseDate" validate:"min=1900,max=2050"`
	Duration    int     `json:"duration" validate:"min=1"`
	Director    string  `json:"director"` //validate:"regexp=^[А-Яа-я ]+$"
	Rating      float64 `json:"rating" validate:"min=0,max=10"`
	Views       int     `json:"views" validate:"min=0"`
	Description string  `json:"description" validate:"max=10"`
	Genres      []int   `json:"genres" validate:"min=1,max=5"`
}

type AnimeGenre struct {
	AnimeID int `json:"animeID"`
	GenreID int `json:"genreID"`
}
