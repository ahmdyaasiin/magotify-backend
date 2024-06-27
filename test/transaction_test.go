package test

import (
	"encoding/json"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	res "github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDetailsProduct(t *testing.T) {
	//
	TestLogin(t)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/product/%s/details", GetRandomProduct()), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

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

	var r model.ProductDetails
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.PD)
	assert.NotNil(t, r.TotalCart)
	assert.NotNil(t, r.ProductDiscount)
}

func TestDetailsProductError(t *testing.T) {
	//
	TestLogin(t)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/product/%s/details", "ab0b3edd-ef6b-4633-b3c4-cac8ac1b776c"), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusInternalServerError, responseBody.Status.Code)
	assert.Nil(t, responseBody.Data)
}

func TestListTransactionShop(t *testing.T) {
	//
	TestValidatePaymentShop(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/transaction/shop", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	var responseBody res.Final
	err = json.Unmarshal(bytes, &responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, responseBody.Status.Code)

	responseData, ok := responseBody.Data.([]interface{})
	if !ok {
		t.Fatalf("it's not list")
	}

	responseJSON, err := json.Marshal(responseData)
	assert.Nil(t, err)

	var r []model.ResponseTransactionShop
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	for _, v := range r {
		assert.NotNil(t, v.ID)
	}
}

func TestDetailsTransactionShop(t *testing.T) {
	//
	TestValidatePaymentShop(t)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/transaction/shop/%s", GetTestTransactionIDShop()), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

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

	var r model.ResponseSpecificTransactionShop
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.Products)
	assert.NotNil(t, r.TransactionID)
	assert.NotNil(t, r.InvoiceNumber)
}

func TestListTransactionPickUp(t *testing.T) {
	//
	TestValidatePaymentPickUp(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/transaction/pick_up", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	var responseBody res.Final
	err = json.Unmarshal(bytes, &responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, responseBody.Status.Code)

	responseData, ok := responseBody.Data.([]interface{})
	if !ok {
		t.Fatalf("it's not list")
	}

	responseJSON, err := json.Marshal(responseData)
	assert.Nil(t, err)

	var r []model.ResponseTransactionPickUp
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	for _, v := range r {
		assert.NotNil(t, v.ID)
	}
}

func TestDetailsTransactionPickUp(t *testing.T) {
	//
	TestValidatePaymentPickUp(t)

	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/transaction/pick_up/%s", GetTestTransactionIDPickUp()), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

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

	var r model.ResponseSpecificTransactionPickUp
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.TransactionID)
	assert.NotNil(t, r.InvoiceNumber)
}
