package repository

import (
	"database/sql"
	"zangetsu/internal/domain/entity"
	"zangetsu/pkg/logging"
)

type AnimeRepository struct {
	db     *sql.DB
	logger logging.Logger
}

type IAnimeRepository interface {
	BeginTransaction() (*sql.Tx, error)
	SaveAnime(anime *entity.AnimeViewModel) *sql.Row
	GetLastID() (*sql.Rows, error)
	SaveAnimeGenres(animeID int, genreID int) error
	DeleteAnimeGenres(animeID int) error
	DeleteAnime(animeID int) error
}

func NewAnimeRepository(db *sql.DB, logger logging.Logger) *AnimeRepository {
	var animeRepo = AnimeRepository{}
	animeRepo.db = db
	animeRepo.logger = logger
	return &animeRepo
}

func (r *AnimeRepository) BeginTransaction() (*sql.Tx, error) {
	return r.db.Begin()
}

func (r *AnimeRepository) GetLastID() (*sql.Rows, error) {
	rows, err := r.db.Query(`select max(id) from anime`)
	if err != nil {
		r.logger.Errorf(err.Error())
		return nil, err
	}
	return rows, nil
}

func (r *AnimeRepository) SaveAnime(anime *entity.AnimeViewModel) *sql.Row {
	row := r.db.QueryRow(`insert into anime(title_rus, title_eng, duration, director, rating, views, description, release_date) values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`, anime.TitleRus, anime.TitleEng, anime.Duration, anime.Director, anime.Rating, anime.Views, anime.Description, anime.ReleaseDate)
	return row
}

func (r *AnimeRepository) SaveAnimeGenres(animeID int, genreID int) error {
	_, err := r.db.Exec(`insert into anime_genres(anime_id, genre_id) values($1, $2)`, animeID, genreID)
	if err != nil {
		r.logger.Errorf(err.Error())
		return err
	}
	return nil
}

func (r *AnimeRepository) DeleteAnime(animeID int) error {
	_, err := r.db.Exec(`delete from anime where id = $1`, animeID)
	if err != nil {
		r.logger.Errorf(err.Error())
		return err
	}
	return nil
}

func (r *AnimeRepository) DeleteAnimeGenres(animeID int) error {
	_, err := r.db.Exec(`delete from anime_genres where anime_id = $1`, animeID)
	if err != nil {
		r.logger.Errorf(err.Error())
		return err
	}
	return nil
}
