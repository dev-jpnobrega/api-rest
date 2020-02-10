package repository

import (
	"fmt"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// UserRepository - Implementation (ref IUserRepository)
type UserRepository struct {
	model entity.User
}

// Login - User login in APP
func (p *UserRepository) Login(email string, pass string) (error, entity.User) {
	fmt.Println("repository:", email, pass)

	return nil, entity.User{
		Name:  "João Paulo",
		Email: "devjpnobrega@github.com",
	}
}
