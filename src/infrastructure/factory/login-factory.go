package factory

import (
	command "github.com/dev-jpnobrega/api-rest/src/domain/command"
	interfaces "github.com/dev-jpnobrega/api-rest/src/domain/contract/interface"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
	repository "github.com/dev-jpnobrega/api-rest/src/infrastructure/repository"
)

var dbMock = map[string]*entity.User{
	"jon@labstack.com": &entity.User{"Jon Snow", "jon@labstack.com"},
	"dev@dev.com.br":   &entity.User{"JP Nobrega", "dev@dev.com.br"},
}

// UserLoginFactory infc
func UserLoginFactory() interfaces.ICommand {
	return &command.LoginCommand{
		Repository: &repository.UserRepository{dbMock},
	}
}
