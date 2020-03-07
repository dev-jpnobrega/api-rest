package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	command "github.com/dev-jpnobrega/api-rest/src/domain/command"
	values "github.com/dev-jpnobrega/api-rest/src/domain/contract/value"
	entity "github.com/dev-jpnobrega/api-rest/src/domain/entity"
	helper "github.com/dev-jpnobrega/api-rest/src/infrastructure/helper"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	routers "github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
)

var usersFaker []*entity.User

type UserRepositoryFaker struct{}

// Login - User login in APP
func (p *UserRepositoryFaker) Login(email string, pass string) (*entity.User, *values.ResponseError) {
	user := &entity.User{}

	for _, u := range usersFaker {
		if email == u.Email {
			user = u
		}
	}

	return user, nil
}

func init() {
	usersFaker = append(usersFaker, &entity.User{
		ID:    uuid.New(),
		Name:  "JP",
		Email: "dev@dev.com.br",
	})

	usersFaker = append(usersFaker, &entity.User{
		ID:    uuid.New(),
		Name:  "ED",
		Email: "dev@ed.con.br",
	})
}

func TestLoginInvalidArgs(t *testing.T) {
	e := helper.CreateEchoServer()

	req := httptest.NewRequest("POST", "/v1/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	command := &command.LoginCommand{
		Repository: &UserRepositoryFaker{},
	}

	h := routers.Adapter(
		command,
		handler.NewHandler(),
	)

	/*
		expected := `{
			"StatusCode": 422,
			"Errors": [
				{
					"Code": 2,
					"Message": "Email.invalid"
				},
				{
					"Code": 2,
					"Message": "Pass.invalid"
				}
			]
		}`
	*/

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
		// assert.Equal(t, expected, rec.Body.String())
	}
}

func TestLoginOk(t *testing.T) {
	e := helper.CreateEchoServer()

	body := `{"email": "dev@dev.com.br", "pass": "megadrive"}`

	req := httptest.NewRequest("POST", "/v1/login", strings.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	command := &command.LoginCommand{
		Repository: &UserRepositoryFaker{},
	}

	h := routers.Adapter(
		command,
		handler.NewHandler(),
	)

	//expected := "{'StatusCode': 422,'Errors': [{'Code': 2,'Message': 'Email.invalid'},{'Code': 2,'Message': 'Pass.invalid'}]}"

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
