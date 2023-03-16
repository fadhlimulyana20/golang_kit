package repository

import (
	"fmt"

	"template/internal/config"
	"template/internal/entities"
	"template/internal/params"
	"template/utils/pagination/gorm_pagination"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RoleRepo struct {
	db   *gorm.DB
	name string
}

type RoleRepository interface {
	// Create a new role
	Create(role entities.Role) (entities.Role, error)
	// Update role
	Update(role entities.Role) (entities.Role, error)
	// List role
	List(param params.RoleFilterParam) ([]entities.Role, int, error)
	// Get Role
	Get(ID int) (entities.Role, error)
	// Get Role By name
	GetByName(name string) (entities.Role, error)
	// Delete Role
	Delete(ID int) (entities.Role, error)
}

// Create new role repository instance
func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepo{
		db:   db,
		name: "Role Repository",
	}
}

func (r *RoleRepo) Create(role entities.Role) (entities.Role, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", r.name))

	if err := r.db.Create(&role).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return role, err
	}

	return role, nil
}

func (r *RoleRepo) Get(ID int) (entities.Role, error) {
	log.Info(fmt.Sprintf("[%s][Get] is executed", r.name))
	var role entities.Role

	db := r.db

	if err := db.Debug().First(&role, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GET] %s", r.name, err.Error()))
		return role, err
	}

	return role, nil
}

func (r *RoleRepo) List(param params.RoleFilterParam) ([]entities.Role, int, error) {
	log.Info(fmt.Sprintf("[%s][Update] is executed", r.name))

	var roles []entities.Role

	var count int64
	r.db.Find(&roles).Count(&count)

	db := r.db

	if err := db.Debug().Scopes(gorm_pagination.Paginate(config.Pagination.Page, config.Pagination.PageLimit)).Order("created_at desc").Find(&roles).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][List] %s", r.name, err.Error()))
		return roles, int(count), err
	}

	return roles, int(count), nil
}

func (r *RoleRepo) Update(role entities.Role) (entities.Role, error) {
	log.Info(fmt.Sprintf("[%s][Create] is executed", r.name))

	if err := r.db.Model(&role).Updates(&role).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Create] %s", r.name, err.Error()))
		return role, err
	}

	return role, nil
}

func (r *RoleRepo) Delete(ID int) (entities.Role, error) {
	log.Info(fmt.Sprintf("[%s][Delete] is executed", r.name))
	var role entities.Role

	if err := r.db.Delete(&role, ID).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][Delete] %s", r.name, err.Error()))
		return role, err
	}

	return role, nil
}

func (r *RoleRepo) GetByName(name string) (entities.Role, error) {
	log.Info(fmt.Sprintf("[%s][GetByName] is executed", r.name))
	var role entities.Role

	if err := r.db.Debug().Where("name = ?", name).First(&role).Error; err != nil {
		log.Error(fmt.Sprintf("[%s][GetByName] %s", r.name, err.Error()))
		return role, err
	}

	return role, nil
}
