package domain

import (
	"fmt"

	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// GetPeopleParam param ss
type GetPeopleParam2 struct {
	name string
	age  string
}

// IPersonRepository infc
type IPersonRepository interface {
	GetPeople(params GetPeopleParam2) (error, []entity.Person)
}

type GetPeopleCommand struct {
	params     GetPeopleParam2
	Repository IPersonRepository
}

// GetModelValidate command
func (p *GetPeopleCommand) GetModelValidate() interface{} {
	return nil
}

// Execute command
func (p *GetPeopleCommand) Execute(input values.DataInput) (error, values.DataResult) {

	context := input

	fmt.Println("COMMAND", context)

	//value, ok := params.Args.(GetPeopleParam)
	var dataResult values.DataResult
	err, data := p.Repository.GetPeople(p.params)

	if err != nil {
		panic(err)
	}

	dataResult.Data = data

	fmt.Println("people:", data, "getPeople")

	return nil, dataResult
}
