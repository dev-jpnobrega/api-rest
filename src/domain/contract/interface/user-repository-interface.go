package domain

import (
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// IUserRepository - Rs
type IUserRepository interface {
	Login(email string, pass string) (*entity.User, *value.ResponseError)
}
