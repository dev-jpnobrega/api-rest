package domain

import (
	"fmt"

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

func (l *LoginCommand) Execute(input values.RequestData) (error, values.ResponseData) {
	var dataResult values.ResponseData
	args := input.Args.(*values.InputLogin)

	fmt.Println("command[args]", args)

	_, user := l.Repository.Login(args.Email, args.Pass)

	dataResult.Data = user

	return nil, dataResult
}
