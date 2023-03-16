package params

import "template/internal/params/generics"

type RoleCreateParam struct {
	Name string `json:"name" validate:"required"`
}

type RoleEditParam struct {
	ID int `json:"id"`
	RoleCreateParam
}

type RoleFilterParam struct {
	Name string `json:"name"`
	generics.GenericFilter
}

type RoleAssignParam struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
}
