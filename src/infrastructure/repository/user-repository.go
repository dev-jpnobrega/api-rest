package infrastructure

import (
	infrastructure "github.com/dev-jpnobrega/api-rest/src/db"
	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// UserRepository - Implementation (ref IUserRepository)
type UserRepository struct {
	DB infrastructure.IDB
}

// Login - User login in APP
func (p *UserRepository) Login(email string, pass string) (*entity.User, *values.ResponseError) {
	user := &entity.User{}
	err := &values.ResponseError{}

	p.DB.GetModel(&entity.User{}).Where("email = ?", email).First(&user)

	if user == nil {
		return nil, err.New("user.not.found", 2, 422)
	}

	return user, nil
}
