package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"zangetsu/internal/domain/entity"
	"zangetsu/internal/domain/repository"
)

type UserService struct {
	userRepo repository.IUserRepository
}

type IUserService interface {
	GetHash(pwd []byte) string
	SignUp(user *entity.UserViewModel) error
	RegistrationByGmail(user *entity.UserRegistrationModel) error
}

func NewUserService(userRepo repository.IUserRepository) *UserService {
	var userService = UserService{}
	userService.userRepo = userRepo
	return &userService
}

func (s *UserService) GetHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (s *UserService) SignUp(userVM *entity.UserViewModel) error {
	row := s.userRepo.GetUser(userVM.Email)
	var user entity.User
	err := row.Scan(&user.ID, &user.RoleID, &user.FirstName, &user.SecondName, &user.Email, &user.Password, &user.RegisteredDate, &user.GmailBind)
	if len(user.Email) != 0 || user.Email != "" {
		return fmt.Errorf("User already registered")
	}

	passwordHash := s.GetHash([]byte(userVM.Password))
	currentDate := time.Now().Format("2006-01-02")

	err = s.userRepo.SaveUser(userVM, 3, passwordHash, currentDate, false)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) RegistrationByGmail(userRM *entity.UserRegistrationModel) error {
	row := s.userRepo.GetUser(userRM.Email)

	var user entity.User
	err := row.Scan(&user.ID, &user.RoleID, &user.FirstName, &user.SecondName, &user.Email, &user.Password, &user.RegisteredDate, &user.GmailBind)
	if len(user.Email) != 0 || user.Email != "" {
		return fmt.Errorf("User already registered")
	}
	userVM := entity.UserViewModel{
		FirstName:  userRM.GivenName,
		SecondName: userRM.FamilyName,
		Email:      userRM.Email,
		Password:   "",
	}
	currentDate := time.Now().Format("2006-01-02")

	err = s.userRepo.SaveUser(&userVM, 3, "", currentDate, true)
	if err != nil {
		return err
	}

	return nil
}
