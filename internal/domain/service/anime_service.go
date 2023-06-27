package service

import (
	"fmt"
	"strconv"
	"zangetsu/internal/domain/entity"
	"zangetsu/internal/domain/repository"
)

type AnimeService struct {
	animeRepo   repository.IAnimeRepository
	animeEsRepo repository.IAnimeESRepository
}

type IAnimeService interface {
	SaveAnime(anime *entity.AnimeViewModel) error
	SearchAnime(query string) ([]*entity.AnimeViewModel, error)
}

func NewAnimeService(animeRepo repository.IAnimeRepository, animeEsRepo repository.IAnimeESRepository) *AnimeService {
	var animeService = AnimeService{}
	animeService.animeRepo = animeRepo
	animeService.animeEsRepo = animeEsRepo
	return &animeService
}

func (s *AnimeService) SearchAnime(query string) ([]*entity.AnimeViewModel, error) {
	animeList, err := s.animeEsRepo.Search(query)
	if err != nil {
		return nil, err
	}
	return animeList, nil
}

func (s *AnimeService) SaveAnime(anime *entity.AnimeViewModel) error {
	// Save anime in Postgres

	var newAnime = entity.AnimeViewModel{
		TitleRus:    anime.TitleRus,
		TitleEng:    anime.TitleEng,
		ReleaseDate: anime.ReleaseDate,
		Duration:    anime.Duration,
		Director:    anime.Director,
		Rating:      anime.Rating,
		Views:       anime.Views,
		Description: anime.Description,
		Genres:      anime.Genres,
	}

	row := s.animeRepo.SaveAnime(&newAnime)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return err
	}

	for _, genre := range anime.Genres {
		err = s.animeRepo.SaveAnimeGenres(id, genre)
		if err != nil {
			err = s.animeRepo.DeleteAnimeGenres(id)
			if err != nil {
				return fmt.Errorf("Delete from table 'anime_genres' failed:\n%s\nDelete genres for animeID: %s", err.Error(), strconv.Itoa(id))
			}
			err = s.animeRepo.DeleteAnime(id)
			if err != nil {
				return fmt.Errorf("Delete from table 'anime' failed:\n%s\nDelete anime for id: %s", err.Error(), strconv.Itoa(id))
			}
			return err
		}

	}

	// Save anime in ElasticSearch

	err = s.animeEsRepo.CreateAnimeIndex()
	err = s.animeEsRepo.Index(anime)
	if err != nil {
		fmt.Errorf("Failed to save anime in ElasticSearch")
	}
	return nil

}
