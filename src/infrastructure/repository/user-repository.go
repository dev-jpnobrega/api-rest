package infrastructure

import (
	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// UserRepository - Implementation (ref IUserRepository)
type UserRepository struct {
	DB map[string]*entity.User
}

// Login - User login in APP
func (p *UserRepository) Login(email string, pass string) (user *entity.User, err *values.ResponseError) {
	user = p.DB[email]

	if user == nil {
		err = err.New("user.not.found", 2, 422)
	}

	return
}
