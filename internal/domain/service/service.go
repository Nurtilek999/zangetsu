package service

import (
	"zangetsu/internal/domain/entity"
	"zangetsu/internal/domain/repository"
	"zangetsu/pkg/logging"
)

type Service struct {
	IUserService
	IAnimeService
}

type IUserService interface {
	SignUp(user *entity.UserViewModel) error
	RegistrationByGmail(user *entity.UserRegistrationModel) error
}

type IAnimeService interface {
	SaveAnime(anime *entity.AnimeViewModel) error
	SearchAnime(query string) ([]*entity.AnimeViewModel, error)
}

func NewService(r *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		NewUserService(r.IUserRepository, logger),
		NewAnimeService(r.IAnimeRepository, r.IAnimeESRepository, logger),
	}
}
