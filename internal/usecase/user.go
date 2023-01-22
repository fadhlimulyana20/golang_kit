package usecase

import (
	"fmt"
	"template/internal/appctx"
	"template/internal/entities"
	"template/internal/params"
	"template/internal/respository"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type user struct {
	repo respository.UserRepository
	name string
}

type UserUsecase interface {
	Create(param params.UserCreateParam) appctx.Response
}

func NewUserUsecase() UserUsecase {
	return &user{
		repo: respository.NewUserRepository(),
		name: "USER USECASE",
	}
}

func (u *user) Create(param params.UserCreateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", u.name))

	var user entities.User
	copier.Copy(&user, &param)

	usr, err := u.repo.Create(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}
