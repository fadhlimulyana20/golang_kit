package usecase

import (
	"fmt"

	"template/internal/appctx"
	"template/internal/entities"
	"template/internal/params"
	"template/internal/repository"
	"template/utils/password"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type user struct {
	repo repository.UserRepository
	name string
}

type UserUsecase interface {
	// Create a new user record
	Create(param params.UserCreateParam) appctx.Response

	// List and filter user record
	List(params.UserListParams) appctx.Response

	// Update user record
	Update(params.UserUpdateParam) appctx.Response

	// Update user record
	Get(int) appctx.Response

	// Delete user record
	Delete(int) appctx.Response
}

func NewUserUsecase(db *gorm.DB) UserUsecase {
	return &user{
		repo: repository.NewUserRepository(db),
		name: "USER USECASE",
	}
}

func (u *user) Create(param params.UserCreateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", u.name))

	var user entities.User
	copier.Copy(&user, &param)

	// Generate Hash Password
	passwd, _ := password.HashPassword(user.Password)
	user.Password = passwd

	usr, err := u.repo.Create(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (u *user) List(param params.UserListParams) appctx.Response {
	log.Info(fmt.Sprintf("[%s][List] is executed", u.name))

	var usrs []entities.User
	users, count, err := u.repo.List(usrs, param)
	if err != nil {
		log.Error(err.Error())
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(users).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (u *user) Update(param params.UserUpdateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Update] is executed", u.name))

	var user entities.User
	user, err := u.repo.Get(user, param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	copier.Copy(&user, &param)

	usr, err := u.repo.Update(user)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (u *user) Get(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Get] is executed", u.name))

	var user entities.User
	user, err := u.repo.Get(user, ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Get] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(user)
}

func (u *user) Delete(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", u.name))

	var user entities.User
	user, err := u.repo.Get(user, ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", u.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(user)
}
