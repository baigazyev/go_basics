package models

type Role struct {
	RoleID   int    `json:"role_id" db:"role_id"`
	RoleName string `json:"role_name" db:"role_name"`
}
