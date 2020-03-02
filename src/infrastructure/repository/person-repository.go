package infrastructure

import (
	"fmt"

	command "github.com/dev-jpnobrega/api-rest/src/domain/command"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

type PersonRepository struct {
	model entity.Person
}

func (p *PersonRepository) GetPeople(params command.GetPeopleParam2) (error, []entity.Person) {
	fmt.Println("repository:", params)

	return nil, []entity.Person{
		entity.Person{
			Name: "João Paulo",
			Age:  31,
		},
		entity.Person{
			Name: "Paulo",
			Age:  32,
		},
	}
}
