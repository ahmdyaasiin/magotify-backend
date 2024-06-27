package test

import (
	"encoding/json"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	res "github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	DeleteAll()

	requestBody := model.RequestUserRegister{
		FullName:    "Raymond Ananda",
		Email:       testEmail,
		Password:    "ini_password_raymond_ananda",
		PhoneNumber: "085982950102",
		Address:     "Jl. Cengger Ayam DLM No.18 A, Tulusrejo",
		District:    "Kec. Lowokwaru",
		City:        "Kota Malang",
		State:       "Jawa Timur",
		PostalCode:  "65141",
		Latitude:    -7.947938841664528,
		Longitude:   112.63089962806873,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/auth/register", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusCreated, responseBody.Status.Code)

	responseMap := responseBody.Data.(map[string]interface{})
	responseJSON, err := json.Marshal(responseMap)
	assert.Nil(t, err)

	var user model.ResponseUser
	err = json.Unmarshal(responseJSON, &user)
	assert.Nil(t, err)

	assert.NotEmpty(t, user.Token)
}

func TestRegisterError(t *testing.T) {
	//
	DeleteAll()

	requestBody := model.RequestUserRegister{
		FullName:    "Raymond Ananda",
		Email:       "",
		Password:    "",
		PhoneNumber: "",
		Address:     "",
		District:    "",
		City:        "",
		State:       "",
		PostalCode:  "",
		Latitude:    0,
		Longitude:   0,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/auth/register", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, responseBody.Status.Code)
	assert.NotNil(t, responseBody.Errors)
}

func TestRegisterDuplicate(t *testing.T) {
	//
	TestRegister(t)

	requestBody := model.RequestUserRegister{
		FullName:    "Raymond Ananda",
		Email:       testEmail,
		Password:    "ini_password_raymond_ananda",
		PhoneNumber: "085982950102",
		Address:     "Jl. Cengger Ayam DLM No.18 A, Tulusrejo",
		District:    "Kec. Lowokwaru",
		City:        "Kota Malang",
		State:       "Jawa Timur",
		PostalCode:  "65141",
		Latitude:    -7.947938841664528,
		Longitude:   112.63089962806873,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/auth/register", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusConflict, responseBody.Status.Code)
	assert.Nil(t, responseBody.Data)
}

func TestLogin(t *testing.T) {
	//
	TestRegister(t)

	requestBody := model.RequestUserLogin{
		Email:    testEmail,
		Password: "ini_password_raymond_ananda",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, responseBody.Status.Code)

	responseMap := responseBody.Data.(map[string]interface{})
	responseJSON, err := json.Marshal(responseMap)
	assert.Nil(t, err)

	var user model.ResponseUser
	err = json.Unmarshal(responseJSON, &user)
	assert.Nil(t, err)

	assert.NotEmpty(t, user.Token)
}

func TestLoginIncorrectCredentials(t *testing.T) {
	//
	TestRegister(t)

	requestBody := model.RequestUserLogin{
		Email:    testEmail,
		Password: "ini_password_raymond_ananda_yang_salah",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/auth/login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, responseBody.Status.Code)
	assert.Nil(t, responseBody.Data)
}
