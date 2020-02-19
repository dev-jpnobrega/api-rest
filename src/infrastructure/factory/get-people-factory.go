package infrastructure

import (
	command "github.com/dev-jpnobrega/api-rest/src/domain/command"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	repository "github.com/dev-jpnobrega/api-rest/src/infrastructure/repository"
)

// GetPeopleFactory infc
func GetPeopleFactory() interfaces.ICommand {
	return &command.GetPeopleCommand{
		Repository: &repository.PersonRepository{},
	}
}
