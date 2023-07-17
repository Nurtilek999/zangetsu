package repository

import (
	"database/sql"
	"github.com/olivere/elastic/v7"
	"zangetsu/internal/domain/entity"
	"zangetsu/pkg/logging"
)

type Repository struct {
	IUserRepository
	IAnimeRepository
	IAnimeESRepository
}

type IUserRepository interface {
	SaveUser(user *entity.UserViewModel, roleID int, passwordHash string, regDate string, gmailBind bool) error
	GetUser(email string) *sql.Row
}

type IAnimeRepository interface {
	BeginTransaction() (*sql.Tx, error)
	SaveAnime(anime *entity.AnimeViewModel) *sql.Row
	GetLastID() (*sql.Rows, error)
	SaveAnimeGenres(animeID int, genreID int) error
	DeleteAnimeGenres(animeID int) error
	DeleteAnime(animeID int) error
}

type IAnimeESRepository interface {
	Index(anime *entity.AnimeViewModel) error
	Search(query string) ([]*entity.AnimeViewModel, error)
	CreateAnimeIndex() error
}

func NewRepository(db *sql.DB, client *elastic.Client, logger logging.Logger, index string) *Repository {
	return &Repository{
		NewUserRepository(db, logger),
		NewAnimeRepository(db, logger),
		NewElasticsearchAnimeRepository(client, index, logger),
	}
}
