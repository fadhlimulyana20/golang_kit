package respository

import (
	"fmt"
	"template/database"
	"template/internal/entities"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type userRepo struct {
	db   *gorm.DB
	name string
}

type UserRepository interface {
	Create(entities.User) (entities.User, error)
}

func NewUserRepository() UserRepository {
	return &userRepo{
		db:   database.ORM(),
		name: "USER REPOSITORY",
	}
}

func (u *userRepo) Create(user entities.User) (entities.User, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", u.name))

	if err := u.db.Create(&user).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return user, err
	}

	return user, nil
}