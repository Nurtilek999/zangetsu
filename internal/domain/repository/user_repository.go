package repository

import (
	"database/sql"
	"zangetsu/internal/domain/entity"
	"zangetsu/pkg/logging"
)

type UserRepository struct {
	db     *sql.DB
	logger logging.Logger
}

type IUserRepository interface {
	SaveUser(user *entity.UserViewModel, roleID int, passwordHash string, regDate string, gmailBind bool) error
	GetUser(email string) *sql.Row
}

func NewUserRepository(db *sql.DB, logger logging.Logger) *UserRepository {
	var userRepo UserRepository
	userRepo.db = db
	userRepo.logger = logger
	return &userRepo
}

func (r *UserRepository) SaveUser(user *entity.UserViewModel, roleID int, passwordHash string, regDate string, gmailBind bool) error {
	_, err := r.db.Exec(`insert into users(role_id, first_name, second_name, email, password_hash, registered_date, gmail_bind) values ($1, $2, $3, $4, $5, $6, $7)`, roleID, user.FirstName, user.SecondName, user.Email, passwordHash, regDate, gmailBind)
	if err != nil {
		r.logger.Errorf(err.Error())
		return err
	}
	return nil
}

func (r *UserRepository) GetUser(email string) *sql.Row {
	row := r.db.QueryRow(`select * from users where email = $1`, &email)
	return row
}
