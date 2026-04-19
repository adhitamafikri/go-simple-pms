package repositories

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRespository interface {
	ReadUsers()
	ReadUserById(id uint64)
	CreateUser(body any)
	UpdateUser(id uint64, body any)
	DeleteUser(id uint64)
	ReadMe()
}

func NewUserRepository(db *sqlx.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ReadUsers() {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) ReadUserById(id uint64) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) CreateUser(body any) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) UpdateUser(id uint64, body any) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) DeleteUser(id uint64) {
	fmt.Println("Not Implemented yet", r.db)
}
func (r *repository) ReadMe() {
	fmt.Println("Not Implemented yet", r.db)
}
