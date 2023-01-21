package usecase

import (
	"fmt"
	"template/internal/appctx"
	"template/internal/entities"
	"template/internal/respository"
	"template/utils/json"
	"template/utils/validator"

	log "github.com/sirupsen/logrus"
)

type user struct {
	repo respository.UserRepository
	name string
}

type UserUsecase interface {
	Create(d appctx.Data) appctx.Response
}

func NewUserUsecase() UserUsecase {
	return &user{
		repo: respository.NewUserRepository(),
		name: "USER USECASE",
	}
}

func (u *user) Create(d appctx.Data) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", u.name))

	var user entities.User
	if err := json.Decode(d.Request.Body, &user); err != nil {
		log.Error("Cannot decode json")
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	if err := validator.Validate(user); err != nil {
		log.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	usr, err := u.repo.Create(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}
