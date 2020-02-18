package domain

import (
	value "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
)

// ICommand infc
type ICommand interface {
	GetModelValidate() interface{}
	Execute(value.RequestData) (value.ResponseData, *value.ResponseError)
}
