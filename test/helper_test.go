package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/firebase"
	"io"
	"log"
	"net/http"
	"os"
)

func DeleteAll() {
	UpdateOrders()
	UpdateTestOrder()
	UpdateTestTransaction()
	DeleteTestWishlist()
	DeleteTestCart()
	DeleteTestAddress()
	DeleteTestUser()
}

func UpdateOrders() {
	//
	q := `UPDATE orders o SET o.status = 'done' WHERE o.status = 'waiting-for-payment' OR o.status = 'in-progress'`

	_, err := db.Exec(q)
	if err != nil {
		return
	}
}

func UpdateTestOrder() {
	//
	addressID := GetRandomAddressID()
	q := `UPDATE orders o JOIN addresses a ON o.address_id = a.id JOIN users u ON a.user_id = u.id SET o.address_id = ? WHERE u.email = ?`

	_, err := db.Exec(q, addressID, testEmail)
	if err != nil {
		return
	}
}

func UpdateTestTransaction() {
	//
	addressID := GetRandomAddressID()
	q := `UPDATE transactions t JOIN addresses a ON t.address_id = a.id JOIN users u ON a.user_id = u.id SET t.address_id = ? WHERE u.email = ?`

	_, err := db.Exec(q, addressID, testEmail)
	if err != nil {
		return
	}
}

func DeleteTestWishlist() {
	//
	q := `DELETE w FROM wishlists w JOIN users u ON w.user_id = u.id WHERE u.email = ?`

	_, err := db.Exec(q, testEmail)
	if err != nil {
		return
	}
}

func DeleteTestCart() {
	//
	q := `DELETE c FROM carts c JOIN users u ON c.user_id = u.id WHERE u.email = ?`

	_, err := db.Exec(q, testEmail)
	if err != nil {
		return
	}
}

func DeleteTestAddress() {
	//
	q := `DELETE a FROM addresses a JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	_, err := db.Exec(q, testEmail)
	if err != nil {
		return
	}
}

func DeleteTestUser() {
	q := `DELETE u FROM users u WHERE u.email = ?`

	_, err := db.Exec(q, testEmail)
	if err != nil {
		return
	}
}

func GetRandomAddressID() string {
	//
	q := `SELECT a.id FROM addresses a JOIN users u ON a.user_id = u.id WHERE u.email != ? ORDER BY RAND() LIMIT 1`

	var addressID string
	err := db.Get(&addressID, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get random address_id: %v", err)
	}

	return addressID
}

func GetRandomProductOnCart() string {
	q := `SELECT c.product_id FROM carts c JOIN users u ON c.user_id = u.id WHERE u.email = ? LIMIT 1`

	var productID string
	err := db.Get(&productID, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get product_id: %v", err)
	}

	return productID
}

func GetRandomProduct() string {
	q := `SELECT p.id FROM products p WHERE p.quantity > 2 ORDER BY RAND() LIMIT 1`

	var productID string
	err := db.Get(&productID, q)
	if err != nil {
		log.Fatalf("failed to get product_id: %v", err)
	}

	return productID
}

func GetRandomProducts(limit int) []string {
	q := `SELECT p.id FROM products p WHERE p.quantity > 2 ORDER BY RAND() LIMIT ?`

	var productIDs []string
	err := db.Select(&productIDs, q, limit)
	if err != nil {
		log.Fatalf("failed to get product_ids: %v", err)
	}

	return productIDs
}

func GetTestUserID() string {
	q := `SELECT u.id FROM users u WHERE u.email = ?`

	var userID string
	err := db.Get(&userID, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get user_id: %v", err)
	}

	return userID
}

func GetTestAddressID() string {
	q := `SELECT a.id FROM addresses a JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	var addressID string
	err := db.Get(&addressID, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get address_id: %v", err)
	}

	return addressID
}

func GetTestInvoiceNumberShop() string {
	//
	q := `SELECT t.invoice_number FROM transactions t JOIN addresses a ON t.address_id = a.id JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	var invoiceNumber string
	err := db.Get(&invoiceNumber, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get invoice_number [shop]: %v", err)
	}

	return invoiceNumber
}

func GetTestInvoiceNumberPickUp() string {
	//
	q := `SELECT o.invoice_number FROM orders o JOIN addresses a ON o.address_id = a.id JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	var invoiceNumber string
	err := db.Get(&invoiceNumber, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get invoice_number [pick_up]: %v", err)
	}

	return invoiceNumber
}

func GetTestTransactionIDShop() string {
	q := `SELECT t.id FROM transactions t JOIN addresses a ON t.address_id = a.id JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	var invoiceNumber string
	err := db.Get(&invoiceNumber, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get transaction_id [shop]: %v", err)
	}

	return invoiceNumber
}

func GetTestTransactionIDPickUp() string {
	q := `SELECT o.id FROM orders o JOIN addresses a ON o.address_id = a.id JOIN users u ON a.user_id = u.id WHERE u.email = ?`

	var invoiceNumber string
	err := db.Get(&invoiceNumber, q, testEmail)
	if err != nil {
		log.Fatalf("failed to get transaction_id [pick_up]: %v", err)
	}

	return invoiceNumber
}

func GetWarehouseID() string {
	//
	q := `SELECT w.id FROM warehouses w JOIN addresses a ON a.id = ? ORDER BY ST_Distance_Sphere(POINT(w.longitude, w.latitude),POINT(a.longitude, a.latitude)) LIMIT 1`

	var warehouseID string
	err := db.Get(&warehouseID, q, GetTestAddressID())
	if err != nil {
		log.Fatalf("failed to get warehouse_id: %v", err)
	}

	return warehouseID
}

func GetVehicleID() string {
	//
	q := `SELECT v.id FROM vehicles v WHERE v.name = 'Bike Pick Up' LIMIT 1`

	var vehicleID string
	err := db.Get(&vehicleID, q)
	if err != nil {
		log.Fatalf("failed to get vehicle_id: %v", err)
	}

	return vehicleID
}

func GetAuthorization() string {
	customToken, err := firebase.CreateCustomToken(GetTestUserID())
	requestBody := struct {
		Token             string `json:"token"`
		ReturnSecureToken bool   `json:"returnSecureToken"`
	}{
		Token:             customToken,
		ReturnSecureToken: true,
	}

	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("failed to marshal request body: %v", err)
	}

	request, err := http.NewRequest(http.MethodPost, "https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key="+os.Getenv("FIREBASE_API_KEY"), bytes.NewReader(bodyJson))
	if err != nil {
		log.Fatalf("failed to create new request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("failed to send request: %v", err)
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	responseBody := new(model.SignInWithCustomToken)
	err = json.Unmarshal(bytes, responseBody)
	if err != nil {
		log.Fatalf("failed to unmarshal response body: %v", err)
	}

	return fmt.Sprintf("Bearer %s", responseBody.IDToken)
}
