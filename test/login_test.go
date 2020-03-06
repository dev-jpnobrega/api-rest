package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	DB "github.com/dev-jpnobrega/api-rest/src/db"
	factory "github.com/dev-jpnobrega/api-rest/src/infrastructure/factory"
	helper "github.com/dev-jpnobrega/api-rest/src/infrastructure/helper"
	handler "github.com/dev-jpnobrega/api-rest/src/infrastructure/http"
	routers "github.com/dev-jpnobrega/api-rest/src/infrastructure/routers"
)

func init() {
	database := &DB.DB{}
	database.Connect("postgres", "postgres://postgres:postgres@localhost/profile?sslmode=disable")
}

func TestLoginInvalidArgs(t *testing.T) {
	e := helper.CreateEchoServer()

	req := httptest.NewRequest("POST", "/v1/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := routers.Adapter(
		factory.UserLoginFactory(),
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

	h := routers.Adapter(
		factory.UserLoginFactory(),
		handler.NewHandler(),
	)

	//expected := "{'StatusCode': 422,'Errors': [{'Code': 2,'Message': 'Email.invalid'},{'Code': 2,'Message': 'Pass.invalid'}]}"

	if assert.NoError(t, h(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
