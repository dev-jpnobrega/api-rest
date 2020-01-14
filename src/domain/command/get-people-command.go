package domain

import (
	"fmt"
	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
)

// GetPeopleParam param ss
type GetPeopleParam struct {
	name string `schema:"name"`
	age  string `schema:"age"`
}

// IPersonRepository infc
type IPersonRepository interface {
	GetPeople(params GetPeopleParam) (error, []entity.Person)
}

type GetPeopleCommand struct {
	params     GetPeopleParam
	Repository IPersonRepository
}

// Execute command
func (p *GetPeopleCommand) Execute(params values.DataInput) (error, values.DataResult) {

	value, ok := params.Args.(GetPeopleParam)
	var dataResult values.DataResult

	fmt.Println("[GetPeopleCommand][value]", value)
	fmt.Println("[GetPeopleCommand][parms]", params)
	fmt.Println("[GetPeopleCommand][ok]", ok)

	p.params.age = "31"

	err, data := p.Repository.GetPeople(p.params)

	if err != nil {
		panic(err)
	}

	dataResult.Data = data

	fmt.Println("people:", data, "getPeople")

	return nil, dataResult
}
