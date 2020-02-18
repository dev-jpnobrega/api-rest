package domain

import (
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
)

type LoginCommand struct {
	Repository interfaces.IUserRepository
}

// GetModelValidate command
func (l *LoginCommand) GetModelValidate() interface{} {
	return &values.InputLogin{}
}

func (l *LoginCommand) Execute(input values.RequestData) (dataResult values.ResponseData, err *values.ResponseError) {
	args := input.Args.(*values.InputLogin)

	user, err := l.Repository.Login(args.Email, args.Pass)

	if err != nil {
		return
	}

	dataResult.StatusCode = 200
	dataResult.Data = user

	return
}
