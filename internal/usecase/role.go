package usecase

import (
	"fmt"

	"template/internal/appctx"
	"template/internal/entities"
	"template/internal/params"
	"template/internal/repository"

	"gorm.io/gorm"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

type role struct {
	repo     repository.RoleRepository
	userRepo repository.UserRepository
	name     string
}

type RoleUsecase interface {
	// Create new role
	Create(param params.RoleCreateParam) appctx.Response
	// Edit a role
	Update(param params.RoleEditParam) appctx.Response
	// Get role list
	List(param params.RoleFilterParam) appctx.Response
	// Get detail role
	Detail(ID int) appctx.Response
	// Delete role
	Delete(ID int) appctx.Response
	// Assign Role to user
	Assign(userID int, roleName string) appctx.Response
	// Revoke Role from user
	Revoke(userID int, roleName string) appctx.Response
}

func NewRoleUsecase(db *gorm.DB) RoleUsecase {
	return &role{
		repo:     repository.NewRoleRepository(db),
		userRepo: repository.NewUserRepository(db),
		name:     "Role Usecase",
	}
}

func (r *role) Create(param params.RoleCreateParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Create] is executed", r.name))

	var role entities.Role
	copier.Copy(&role, &param)

	usr, err := r.repo.Create(role)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (r *role) Update(param params.RoleEditParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Update] is executed", r.name))

	// get role
	var role entities.Role
	role, err := r.repo.Get(param.ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	copier.Copy(&role, &param)

	usr, err := r.repo.Update(role)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Update] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(usr)
}

func (r *role) List(param params.RoleFilterParam) appctx.Response {
	log.Info(fmt.Sprintf("[%s][List] is executed", r.name))

	// get role list
	var roles []entities.Role
	roles, count, err := r.repo.List(param)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrors(err.Error())
	}

	return *appctx.NewResponse().WithData(roles).WithMeta(int64(param.Page), int64(param.Limit), int64(count))
}

func (r *role) Detail(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Detail] is executed", r.name))

	// get role
	var role entities.Role
	role, err := r.repo.Get(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Detail] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithData(role)
}

func (r *role) Delete(ID int) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", r.name))

	// get role
	_, err := r.repo.Delete(ID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("role deleted sucessfully")
}

func (r *role) Assign(userID int, roleName string) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Assign] is executed", r.name))

	var role entities.Role
	role, err := r.repo.GetByName(roleName)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	var user entities.User
	user, err = r.userRepo.Get(user, userID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	_, err = r.userRepo.AddRole(user, role)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Assign] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("role assigned sucessfully")
}

func (r *role) Revoke(userID int, roleName string) appctx.Response {
	log.Info(fmt.Sprintf("[%s][Revoke] is executed", r.name))

	var role entities.Role
	role, err := r.repo.GetByName(roleName)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	var user entities.User
	user, err = r.userRepo.Get(user, userID)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	_, err = r.userRepo.RemoveRole(user, role)
	if err != nil {
		log.Error(fmt.Sprintf("[%s][Revoke] %s", r.name, err.Error()))
		return *appctx.NewResponse().WithErrorObj(err)
	}

	return *appctx.NewResponse().WithMessage("role revoked sucessfully")
}
