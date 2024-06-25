package model

import "github.com/ahmdyaasiin/magotify-backend/internal/app/entity"

type ServicesOngkir struct {
	Name        string `json:"name"`
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
	Estimation  string `json:"estimation"`
	Note        string `json:"note"`
}

type RecentAddress struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	District    string  `json:"district"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postal_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	IsPrimary   bool    `json:"is_primary"`
	PhoneNumber string  `json:"phone_number"`
}

type PaymentShop struct {
	RecentAddress entity.Address   `json:"recent_address"`
	Addresses     []entity.Address `json:"addresses"`
	Product       []ProductCart    `json:"products"`
	Services      []ServicesOngkir `json:"services"`
	RecentVoucher entity.Voucher   `json:"recent_voucher"`
	Vouchers      []entity.Voucher `json:"vouchers"`
}

type PaymentPickUp struct {
	RecentAddress    entity.Address   `json:"recent_address"`
	Addresses        []entity.Address `json:"addresses"`
	WarehouseDetails entity.Warehouse `json:"warehouse_details"`
	RecentVehicles   entity.Vehicle   `json:"recent_vehicles"`
	Vehicles         []entity.Vehicle `json:"vehicles"`
	RecentVoucher    entity.Voucher   `json:"recent_voucher"`
	Vouchers         []entity.Voucher `json:"vouchers"`
	Distance         float64          `json:"distance"`
}

type RequestCreatePayment struct {
	ProductIDs     []string `json:"product_ids" validate:"required"`
	Quantities     []string `json:"quantities" validate:"required"`
	AddressID      string   `json:"address_id" validate:"required"`
	VoucherID      string   `json:"voucher_id"`
	ExpeditionName string   `json:"expedition_name" validate:"required"`
	ExpeditionType string   `json:"expedition_type" validate:"required"`
}

type RequestCreatePickUp struct {
	Weight      float64 `json:"weight"`
	AddressID   string  `json:"address_id"`
	WarehouseID string  `json:"warehouse_id"`
	VehicleID   string  `json:"vehicle_id"`
	VoucherID   string  `json:"voucher_id"`
	//
}

type FirebaseTracking struct {
	DriverID        string  `json:"driver_id"`
	DriverLatitude  float64 `json:"driver_latitude"`
	DriverLongitude float64 `json:"driver_longitude"`
	UserID          string  `json:"user_id"`
	CreatedAt       int64   `json:"created_at"`
	UpdatedAt       int64   `json:"updated_at"`
}

type RequestValidatePayment struct {
	OrderID           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	SignatureKey      string `json:"signature_key"`
	StatusCode        string `json:"status_code"`
	GrossAmount       string `json:"gross_amount"`
}

type ResponseCreatePayment struct {
	PaymentID string `json:"payment_id"`
}

type ResponseValidatePayment struct {
	IsPaid bool `json:"is_paid"`
}
