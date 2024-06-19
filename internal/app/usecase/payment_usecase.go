package usecase

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/rajaongkir"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type InterfacePaymentUseCase interface {
	//
	PickUp(auth string, warehouseID string, addressID string, vehiclesID string, voucherID string) (*model.PaymentPickUp, error)
	CreateShop(auth string, request *model.RequestCreatePayment) (*model.ResponseCreatePayment, error)
	Shop(auth string, productIds []string, addressID string, quantites []string, voucherID string) (*model.PaymentShop, error)
}

type PaymentUseCase struct {
	//
	DB                  *sqlx.DB
	Log                 *logrus.Logger
	PaymentRepository   repository.InterfacePaymentRepository
	UserRepository      repository.InterfaceUserRepository
	AddressRepository   repository.InterfaceAddressRepository
	VoucherRepository   repository.InterfaceVoucherRepository
	WarehouseRepository repository.InterfaceWarehouseRepository
	VehicleRepository   repository.InterfaceVehicleRepository
	DriverRepository    repository.InterfaceDriverRepository
	ProductRepository   repository.InterfaceProductRepository
}

func NewPaymentUseCase(db *sqlx.DB, log *logrus.Logger, pr repository.InterfacePaymentRepository, ur repository.InterfaceUserRepository, ar repository.InterfaceAddressRepository, vr repository.InterfaceVoucherRepository, wr repository.InterfaceWarehouseRepository, ver repository.InterfaceVehicleRepository, dr repository.InterfaceDriverRepository, prr repository.InterfaceProductRepository) InterfacePaymentUseCase {
	//
	return &PaymentUseCase{
		DB:                  db,
		Log:                 log,
		PaymentRepository:   pr,
		UserRepository:      ur,
		AddressRepository:   ar,
		VoucherRepository:   vr,
		WarehouseRepository: wr,
		VehicleRepository:   ver,
		DriverRepository:    dr,
		ProductRepository:   prr,
	}
}

func (u *PaymentUseCase) PickUp(auth string, warehouseID string, addressID string, vehiclesID string, voucherID string) (*model.PaymentPickUp, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	// get address
	var address entity.Address

	if addressID == "" {
		fmt.Println("kosong")
		err = u.AddressRepository.GetPrimaryAddress(tx, &address, user)
	} else {
		fmt.Println("isi")
		err = u.AddressRepository.FindAddressUser(tx, "id", addressID, &address, user)
	}

	if err != nil {
		fmt.Println("0")
		return nil, err
	}

	var addresses []entity.Address

	err = u.AddressRepository.FindExcept(tx, address.ID, &addresses, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	// warehouse details
	var warehouse entity.Warehouse

	err = u.WarehouseRepository.FindBy(tx, "id", warehouseID, &warehouse)
	if err != nil {
		return nil, err
	}

	// vehicle
	var vehicle entity.Vehicle

	if vehiclesID != "" {
		err = u.DriverRepository.GetSpecificDriver(tx, &vehicle, vehiclesID, warehouseID)
		if err != nil {
			return nil, err
		}
	}

	// get vehicles except
	var vehicles []entity.Vehicle

	err = u.DriverRepository.GetDriverExcept(tx, &vehicles, vehiclesID, warehouseID)
	if err != nil {
		fmt.Println("error nih")
		return nil, err
	}

	// get vouchers
	var voucher entity.Voucher

	if voucherID != "" {
		err = u.VoucherRepository.FindBy(tx, "id", voucherID, &voucher, user)
		if err != nil {
			fmt.Println("voucher not found")
			return nil, err
		}
	}

	// get all vouchers except
	var vouchers []entity.Voucher

	err = u.VoucherRepository.FindExcept(tx, voucher.ID, &vouchers, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	return &model.PaymentPickUp{
		RecentAddress:    address,
		Addresses:        addresses,
		WarehouseDetails: warehouse,
		RecentVehicles:   vehicle,
		Vehicles:         vehicles,
		RecentVoucher:    voucher,
		Vouchers:         vouchers,
	}, nil
}

func (u *PaymentUseCase) CreateShop(auth string, request *model.RequestCreatePayment) (*model.ResponseCreatePayment, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	var totalAmount float64

	// check product IDs and quantities
	for i := range request.ProductIDs {
		q, err := strconv.Atoi(request.Quantities[i])
		if err != nil {
			return nil, err
		}

		if q <= 0 {
			return nil, fiber.ErrBadRequest
		}

		var product *entity.Product
		err = u.ProductRepository.FindBy(tx, "id", request.ProductIDs[i], product)
		if err != nil {
			return nil, err
		}

		if product.Quantity < q {
			return nil, fiber.ErrConflict
		}

		totalAmount += product.Price * float64(q)

	}

	// check address ID
	var address entity.Address

	err = u.AddressRepository.FindBy(tx, "id", request.AddressID, &address)
	if err != nil {
		return nil, err
	}

	// check voucher ID
	var voucher entity.Voucher

	if request.VoucherID != "" {
		err = u.VoucherRepository.FindBy(tx, "id", request.VoucherID, &voucher, user)
		if err != nil {
			return nil, err
		}
	}

	// check Expedition

	return nil, nil
}

func (u *PaymentUseCase) Shop(auth string, productIds []string, addressID string, quantites []string, voucherID string) (*model.PaymentShop, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	// get address
	var address entity.Address

	if addressID == "" {
		fmt.Println("kosong")
		err = u.AddressRepository.GetPrimaryAddress(tx, &address, user)
	} else {
		fmt.Println("isi")
		err = u.AddressRepository.FindAddressUser(tx, "id", addressID, &address, user)
	}

	if err != nil {
		fmt.Println("0")
		return nil, err
	}

	var addresses []entity.Address

	err = u.AddressRepository.FindExcept(tx, address.ID, &addresses, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	// get vouchers
	var voucher entity.Voucher

	if voucherID != "" {
		err = u.VoucherRepository.FindBy(tx, "id", voucherID, &voucher, user)
		if err != nil {
			fmt.Println("voucher not found")
			return nil, err
		}
	}

	// get all vouchers except
	var vouchers []entity.Voucher

	err = u.VoucherRepository.FindExcept(tx, voucher.ID, &vouchers, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	// ge tall products
	var products []model.ProductCart

	err = u.PaymentRepository.ProductsPaymentShop(tx, &products, productIds)
	if err != nil {
		fmt.Println("2")
		return nil, err
	}

	for i := range products {
		q, err := strconv.Atoi(quantites[i])
		if err != nil {
			return nil, err
		}

		products[i].QuantityCart = q
	}

	// weight
	var weight float64
	for _, p := range products {
		weight += p.Weight
	}

	// services
	var services []model.ServicesOngkir

	postalCode, err := strconv.Atoi(address.PostalCode)
	if err != nil {
		return nil, err
	}

	err = rajaongkir.CheckCost(postalCode, weight, &services)
	if err != nil {
		return nil, err
	}

	return &model.PaymentShop{
		RecentAddress: address,
		Addresses:     addresses,
		Product:       products,
		Services:      services,
		RecentVoucher: voucher,
		Vouchers:      vouchers,
	}, nil
}
