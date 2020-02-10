package domain

import (
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// IUserRepository - Rs
type IUserRepository interface {
	Login(email string, pass string) (error, entity.User)
}