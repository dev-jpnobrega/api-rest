package domain

import (
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
)

// ICommand infc
type ICommand interface {
	ModelValidate(p interface{}, model interface{}) error
	Execute(p value.DataInput) (error, value.DataResult)
}
