package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type RoleRepository interface {
	ReadRoles()
	ReadRoleById(id uint64)
	CreateRole()
	UpdateRole(id uint64, body any)
	DeleteRole(id uint64)
	AssignRole(id uint64, userId uint64)
}

func NewRoleRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ReadRoles() {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) ReadRoleById(id uint64) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) CreateRole(body any) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) UpdateRole(id uint64, body any) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) DeleteRole(id uint64) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) AssignRole(id uint64, userId uint64) {
	fmt.Println("Not Implemented yet", r.db)
}
