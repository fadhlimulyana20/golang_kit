package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"template/internal/appctx"
	"template/internal/params"
	"template/internal/usecase"
	"template/utils/json"
	"template/utils/validator"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type role struct {
	handler Handler
	usecase usecase.RoleUsecase
	name    string
}

type RoleHandler interface {
	// Create a new role
	Create(w http.ResponseWriter, r *http.Request)
	// Get list of roles
	Read(w http.ResponseWriter, r *http.Request)
	// Update a role
	Update(w http.ResponseWriter, r *http.Request)
	// Delete a role
	Delete(w http.ResponseWriter, r *http.Request)
	// Get detaile of Role
	Detail(w http.ResponseWriter, r *http.Request)
	// Assign Role to User
	AssignRole(w http.ResponseWriter, r *http.Request)
	// Revoke Role from User
	RevokeRole(w http.ResponseWriter, r *http.Request)
}

func NewRoleHandler(db *gorm.DB) RoleHandler {
	return &role{
		name:    "Role Handler",
		usecase: usecase.NewRoleUsecase(db),
	}
}

func (ro *role) Create(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Create] is executed", ro.name))
	startTime := time.Now()

	var param params.RoleCreateParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error()).WithCode(http.StatusBadRequest)
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := ro.usecase.Create(param)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) Read(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Read] is executed", ro.name))
	startTime := time.Now()

	var param params.RoleFilterParam
	ctx := appctx.NewResponse()

	if err := decoder.Decode(&param, r.URL.Query()); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := ro.usecase.List(param)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) Update(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Update] is executed", ro.name))
	startTime := time.Now()

	var param params.RoleEditParam
	ctx := appctx.NewResponse()

	id := chi.URLParam(r, "id")
	param.ID, _ = strconv.Atoi(id)

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := ro.usecase.Update(param)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) Delete(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Delete] is executed", ro.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	roleID, _ := strconv.Atoi(id)

	resp := ro.usecase.Delete(roleID)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) Detail(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][Detail] is executed", ro.name))
	startTime := time.Now()

	id := chi.URLParam(r, "id")
	roleID, _ := strconv.Atoi(id)

	resp := ro.usecase.Detail(roleID)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) AssignRole(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][AssignRole] is executed", ro.name))
	startTime := time.Now()

	var param params.RoleAssignParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := ro.usecase.Assign(param.UserID, param.Role)
	ro.handler.Response(w, resp, startTime, time.Now())
}

func (ro *role) RevokeRole(w http.ResponseWriter, r *http.Request) {
	logrus.Info(fmt.Sprintf("[%s][RevokeRole] is executed", ro.name))
	startTime := time.Now()

	var param params.RoleAssignParam
	ctx := appctx.NewResponse()

	if err := json.Decode(r.Body, &param); err != nil {
		logrus.Error("Cannot decode json")
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	if err := validator.Validate(param); err != nil {
		logrus.Error(err.Error())
		ctx = ctx.WithErrors(err.Error())
		ro.handler.Response(w, *ctx, startTime, time.Now())
		return
	}

	resp := ro.usecase.Revoke(param.UserID, param.Role)
	ro.handler.Response(w, resp, startTime, time.Now())
}
