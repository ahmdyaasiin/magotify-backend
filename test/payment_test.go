package test

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	res "github.com/ahmdyaasiin/magotify-backend/internal/pkg/response"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestCreatePaymentShop(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestCreatePayment{
		ProductIDs:     GetRandomProducts(2),
		Quantities:     []string{"1", "2"},
		AddressID:      GetTestAddressID(),
		ExpeditionName: "jne",
		ExpeditionType: "REG",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/shop/create", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

	response, err := fiber.Test(request, 10*1500)
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

	var r model.ResponseCreatePayment
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.ID)
	assert.NotNil(t, r.PaymentID)
}

func TestCreatePaymentShopError(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestCreatePayment{
		ProductIDs:     []string{},
		Quantities:     []string{},
		AddressID:      "",
		ExpeditionName: "",
		ExpeditionType: "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/shop/create", strings.NewReader(string(bodyJson)))
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

	assert.Equal(t, http.StatusBadRequest, responseBody.Status.Code)
	assert.Nil(t, responseBody.Data)
}

func TestValidatePaymentShop(t *testing.T) {
	//
	TestCreatePaymentShop(t)

	requestBody := model.RequestValidatePayment{
		OrderID:           GetTestInvoiceNumberShop(),
		TransactionStatus: "settlement",
		PaymentType:       "bank_transfer",
		FraudStatus:       "accept",
		SignatureKey:      "",
		StatusCode:        "200",
		GrossAmount:       "216801.00",
	}

	signaturePayload := requestBody.OrderID + requestBody.StatusCode + requestBody.GrossAmount + os.Getenv("MIDTRANS_SERVER_KEY")
	sha512Value := sha512.New()
	sha512Value.Write([]byte(signaturePayload))

	signatureKey := fmt.Sprintf("%x", sha512Value.Sum(nil))
	requestBody.SignatureKey = signatureKey

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/shop/validate", strings.NewReader(string(bodyJson)))
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
}

func TestCreatePaymentPickUp(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestCreatePickUp{
		Weight:      4.2,
		AddressID:   GetTestAddressID(),
		WarehouseID: GetWarehouseID(),
		VehicleID:   GetVehicleID(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/pick_up/create", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", GetAuthorization())

	response, err := fiber.Test(request, 10*1500)
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

	var r model.ResponseCreatePayment
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.ID)
	assert.NotNil(t, r.PaymentID)
}

func TestCreatePaymentPickUpError(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestCreatePickUp{
		Weight:      0,
		AddressID:   "",
		WarehouseID: "",
		VehicleID:   "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/pick_up/create", strings.NewReader(string(bodyJson)))
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

	assert.Equal(t, http.StatusBadRequest, responseBody.Status.Code)
	assert.Nil(t, responseBody.Data)
}

func TestValidatePaymentPickUp(t *testing.T) {
	//
	TestCreatePaymentPickUp(t)

	requestBody := model.RequestValidatePayment{
		OrderID:           GetTestInvoiceNumberPickUp(),
		TransactionStatus: "settlement",
		PaymentType:       "bank_transfer",
		FraudStatus:       "accept",
		SignatureKey:      "",
		StatusCode:        "200",
		GrossAmount:       "216801.00",
	}

	signaturePayload := requestBody.OrderID + requestBody.StatusCode + requestBody.GrossAmount + os.Getenv("MIDTRANS_SERVER_KEY")
	sha512Value := sha512.New()
	sha512Value.Write([]byte(signaturePayload))

	signatureKey := fmt.Sprintf("%x", sha512Value.Sum(nil))
	requestBody.SignatureKey = signatureKey

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/payment/pick_up/validate", strings.NewReader(string(bodyJson)))
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
}
