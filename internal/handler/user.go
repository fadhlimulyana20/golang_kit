package handler

import (
	"fmt"
	"net/http"
	"template/internal/appctx"
	"template/internal/usecase"
	"time"

	"github.com/sirupsen/logrus"
)

type user struct {
	handler handler
	usecase usecase.UserUsecase
	name    string
}

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler() UserHandler {
	return &user{
		usecase: usecase.NewUserUsecase(),
		name:    "USER HANDLER",
	}
}

func (u *user) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", u.name))
	d := appctx.Data{
		Request: r,
	}

	resp := u.usecase.Create(d)
	u.handler.Response(w, resp, time.Now())
}
