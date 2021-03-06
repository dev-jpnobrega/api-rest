package infrastructure

import (
	infrastructure "github.com/dev-jpnobrega/api-rest/src/db"
	command "github.com/dev-jpnobrega/api-rest/src/domain/command"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	repository "github.com/dev-jpnobrega/api-rest/src/infrastructure/repository"
)

// UserLoginFactory infc
func UserLoginFactory() interfaces.ICommand {
	return &command.LoginCommand{
		Repository: &repository.UserRepository{&infrastructure.DB{}},
	}
}
