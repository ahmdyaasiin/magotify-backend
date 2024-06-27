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

func TestGetCart(t *testing.T) {
	//
	TestAddCart(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/user/cart", nil)
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

	var r model.MyCart
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.Product)
	assert.NotNil(t, r.TotalCart)
	assert.NotNil(t, r.HotItems)
}

func TestAddCart(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestAddCart{
		ProductID: GetRandomProduct(),
		Quantity:  2,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/cart/manage", strings.NewReader(string(bodyJson)))
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

	var r model.ResponseAddCart
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.TotalCart)
}

func TestUpdateCart(t *testing.T) {
	//
	TestAddCart(t)

	requestBody := model.RequestAddCart{
		ProductID: GetRandomProductOnCart(),
		Quantity:  1,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/cart/manage", strings.NewReader(string(bodyJson)))
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

	var r model.ResponseAddCart
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.TotalCart)
}

func TestDeleteCart(t *testing.T) {
	//
	TestAddCart(t)

	requestBody := model.RequestAddCart{
		ProductID: GetRandomProductOnCart(),
		Quantity:  0,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/cart/manage", strings.NewReader(string(bodyJson)))
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

	var r model.ResponseAddCart
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.TotalCart)
}

func TestGetCartUnauthorized(t *testing.T) {
	//
	TestAddCart(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/user/cart", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer wrong_bearer_token")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, responseBody.Status.Code)
}

func TestAddCartProductNotFound(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestAddCart{
		ProductID: "wrong_product_id",
		Quantity:  2,
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/cart/manage", strings.NewReader(string(bodyJson)))
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

func TestGetWishlist(t *testing.T) {
	//
	TestManageWishlist(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/user/wishlist", nil)
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

	var r model.MyWishlist
	err = json.Unmarshal(responseJSON, &r)
	assert.Nil(t, err)

	assert.NotNil(t, r.TotalCart)
	assert.NotNil(t, r.Product)
}

func TestManageWishlist(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestManageWishlist{
		ProductID: GetRandomProduct(),
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/wishlist/manage", strings.NewReader(string(bodyJson)))
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
	assert.NotNil(t, responseBody.Data)
}

func TestGetWishlistUnauthorized(t *testing.T) {
	//
	TestManageWishlist(t)

	request := httptest.NewRequest(http.MethodGet, "/v1/user/wishlist", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer wrong_bearer_token")

	response, err := fiber.Test(request, 10*1000)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(res.Final)
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusUnauthorized, responseBody.Status.Code)
}

func TestManageWishlistProductNotFound(t *testing.T) {
	//
	TestLogin(t)

	requestBody := model.RequestManageWishlist{
		ProductID: "wrong_product_id",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/v1/user/wishlist/manage", strings.NewReader(string(bodyJson)))
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
