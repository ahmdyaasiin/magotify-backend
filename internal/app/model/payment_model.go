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
}

type RequestCreatePayment struct {
	ProductIDs     []string `json:"product_ids"`
	Quantities     []string `json:"quantities"`
	AddressID      string   `json:"address_id"`
	VoucherID      string   `json:"voucher_id"`
	ExpeditionName string   `json:"expedition_name"`
	ExpeditionType string   `json:"expedition_type"`
}

type ResponseCreatePayment struct {
	PaymentID string `json:"payment_id"`
}
