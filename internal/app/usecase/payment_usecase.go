package usecase

import (
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/firebase"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/midtrans"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/rajaongkir"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	BIKE  = 250
	CAR   = 750
	TRUCK = 2000
)

type InterfacePaymentUseCase interface {
	//
	ValidatePickUp(request *model.RequestValidatePayment) (*model.ResponseValidatePayment, error)
	ValidateShop(request *model.RequestValidatePayment) (*model.ResponseValidatePayment, error)
	PickUp(auth string, warehouseID string, addressID string, vehiclesID string, voucherID string) (*model.PaymentPickUp, error)
	CreateShop(auth string, request *model.RequestCreatePayment) (*model.ResponseCreatePayment, error)
	CreatePickUp(auth string, request *model.RequestCreatePickUp) (*model.ResponseCreatePayment, error)
	Shop(auth string, productIds []string, addressID string, quantites []string, voucherID string) (*model.PaymentShop, error)
}

type PaymentUseCase struct {
	//
	DB                        *sqlx.DB
	Log                       *logrus.Logger
	PaymentRepository         repository.InterfacePaymentRepository
	UserRepository            repository.InterfaceUserRepository
	AddressRepository         repository.InterfaceAddressRepository
	VoucherRepository         repository.InterfaceVoucherRepository
	WarehouseRepository       repository.InterfaceWarehouseRepository
	VehicleRepository         repository.InterfaceVehicleRepository
	DriverRepository          repository.InterfaceDriverRepository
	ProductRepository         repository.InterfaceProductRepository
	TransactionRepository     repository.InterfaceTransactionRepository
	TransactionItemRepository repository.InterfaceTransactionItemRepository
	OrderRepository           repository.InterfaceOrderRepository
}

func NewPaymentUseCase(db *sqlx.DB, log *logrus.Logger, pr repository.InterfacePaymentRepository, ur repository.InterfaceUserRepository, ar repository.InterfaceAddressRepository, vr repository.InterfaceVoucherRepository, wr repository.InterfaceWarehouseRepository, ver repository.InterfaceVehicleRepository, dr repository.InterfaceDriverRepository, prr repository.InterfaceProductRepository, tr repository.InterfaceTransactionRepository, tir repository.InterfaceTransactionItemRepository, or repository.InterfaceOrderRepository) InterfacePaymentUseCase {
	//
	return &PaymentUseCase{
		DB:                        db,
		Log:                       log,
		PaymentRepository:         pr,
		UserRepository:            ur,
		AddressRepository:         ar,
		VoucherRepository:         vr,
		WarehouseRepository:       wr,
		VehicleRepository:         ver,
		DriverRepository:          dr,
		ProductRepository:         prr,
		TransactionRepository:     tr,
		TransactionItemRepository: tir,
		OrderRepository:           or,
	}
}

func (u *PaymentUseCase) ValidatePickUp(request *model.RequestValidatePayment) (*model.ResponseValidatePayment, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// signature filter
	signaturePayload := request.OrderID + request.StatusCode + request.GrossAmount + os.Getenv("MIDTRANS_SERVER_KEY")
	sha512Value := sha512.New()
	sha512Value.Write([]byte(signaturePayload))

	signatureKey := fmt.Sprintf("%x", sha512Value.Sum(nil))

	if signatureKey != request.SignatureKey {
		return nil, errors.New("invalid signature key")
	}

	order := new(entity.Order)
	err = u.OrderRepository.FindBy(tx, "invoice_number", request.OrderID, order)
	if err != nil {
		return nil, err
	}

	if order.Status != "waiting-for-payment" {
		return nil, errors.New("already paid")
	}

	paymentStatus := false
	if request.PaymentType == string(snap.PaymentTypeCreditCard) {
		paymentStatus = (request.TransactionStatus == "capture") && (request.FraudStatus == "accept")
	} else {
		paymentStatus = (request.TransactionStatus == "settlement") && (request.FraudStatus == "accept")
	}

	if paymentStatus {
		// update transaction status
		order.Status = "in-progress"
		order.PaymentType = request.PaymentType

		// update driverID

		err = u.OrderRepository.Update(tx, order)
		if err != nil {
			return nil, err
		}

		var driver entity.Driver
		err = u.DriverRepository.FindBy(tx, "id", order.DriverID, &driver)
		if err != nil {
			return nil, err
		}

		fmt.Println("passed 1")

		var warehouse entity.Warehouse
		err = u.WarehouseRepository.FindBy(tx, "id", driver.WarehouseID, &warehouse)
		if err != nil {
			fmt.Println("error warehouse nih: " + err.Error())
			return nil, err
		}

		fmt.Println("passed 2")

		var address entity.Address
		err = u.AddressRepository.FindBy(tx, "id", order.AddressID, &address)
		if err != nil {
			return nil, err
		}

		fmt.Println("passed 3")

		// setup firebase
		timeNow := time.Now().Local().UnixNano()
		firebaseTracking := &model.FirebaseTracking{
			DriverID:        order.DriverID,
			DriverLatitude:  warehouse.Latitude,
			DriverLongitude: warehouse.Longitude,
			UserID:          address.UserID,
			CreatedAt:       timeNow,
			UpdatedAt:       timeNow,
		}

		// add chat node

		err = firebase.SaveData("tracking/"+order.ID, firebaseTracking)
		if err != nil {
			return nil, err
		}
	} else if request.TransactionStatus == "expire" {
		//
		err = u.TransactionRepository.UpdateSpecificExpiredOrder(tx, order.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseValidatePayment{
		IsPaid: paymentStatus,
	}, nil
}

func (u *PaymentUseCase) ValidateShop(request *model.RequestValidatePayment) (*model.ResponseValidatePayment, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	fmt.Println(request)

	// signature filter
	signaturePayload := request.OrderID + request.StatusCode + request.GrossAmount + os.Getenv("MIDTRANS_SERVER_KEY")
	sha512Value := sha512.New()
	sha512Value.Write([]byte(signaturePayload))

	signatureKey := fmt.Sprintf("%x", sha512Value.Sum(nil))

	if signatureKey != request.SignatureKey {
		return nil, errors.New("invalid signature key")
	}

	transaction := new(entity.Transaction)
	err = u.TransactionRepository.FindBy(tx, "invoice_number", request.OrderID, transaction)
	if err != nil {
		fmt.Println("salahnya di sini nih pasti cok: " + err.Error())
		return nil, err
	}

	if transaction.Status != "waiting-for-payment" {
		return nil, errors.New("already paid")
	}

	paymentStatus := false
	if request.PaymentType == string(snap.PaymentTypeCreditCard) {
		paymentStatus = (request.TransactionStatus == "capture") && (request.FraudStatus == "accept")
	} else {
		paymentStatus = (request.TransactionStatus == "settlement") && (request.FraudStatus == "accept")
	}

	if paymentStatus {
		// update transaction status
		transaction.Status = "packing"
		transaction.PaymentType = request.PaymentType
		err = u.TransactionRepository.Update(tx, transaction)
		if err != nil {
			return nil, err
		}
	} else if request.TransactionStatus == "expire" {
		// expired

		transaction.Status = "cancel"
		err = u.TransactionRepository.Update(tx, transaction)
		if err != nil {
			return nil, err
		}

		// if expired add the quantity of products back
		err = u.ProductRepository.RollBackQuantityIfCancel(tx, transaction.ID)
		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseValidatePayment{
		IsPaid: paymentStatus,
	}, nil
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

	fmt.Println("1")

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

	fmt.Println("2")

	// get vehicles except
	var vehicles []entity.Vehicle

	err = u.DriverRepository.GetDriverExcept(tx, &vehicles, vehiclesID, warehouseID)
	if err != nil {
		fmt.Println("error nih")
		return nil, err
	}

	var distance float64
	err = u.AddressRepository.GetDistance(tx, &distance, address.ID, warehouse.ID)
	if err != nil {
		fmt.Println("jarak")
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

	fmt.Println("3")

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
		Distance:         distance,
	}, nil
}

func (u *PaymentUseCase) CreatePickUp(auth string, request *model.RequestCreatePickUp) (*model.ResponseCreatePayment, error) {
	tx, err := u.DB.Beginx()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	// check weight
	if request.Weight <= 0 {
		return nil, errors.New("weight should be positive")
	}

	// check address ID
	var address entity.Address

	err = u.AddressRepository.FindBy(tx, "id", request.AddressID, &address)
	if err != nil {
		return nil, err
	}

	// check warehouse and vehicle
	var driver entity.Driver

	err = u.DriverRepository.FindAvailableDriver(tx, request.WarehouseID, request.VehicleID, &driver)
	if err != nil {

		fmt.Println("cioks: " + err.Error())
		fmt.Println("driver not found")

		return nil, err
	}

	fmt.Println("1")

	// check voucher ID
	var voucher entity.Voucher

	if request.VoucherID != "" {
		err = u.VoucherRepository.FindBy(tx, "id", request.VoucherID, &voucher, user)
		if err != nil {
			return nil, err
		}

		// voucher
		voucher.Status = false
		err = u.VoucherRepository.Update(tx, &voucher)
		if err != nil {
			return nil, err
		}
	}

	var vehicle entity.Vehicle
	err = u.VehicleRepository.FindBy(tx, "id", request.VehicleID, &vehicle)
	if err != nil {
		return nil, err
	}

	var distance float64
	err = u.AddressRepository.GetDistance(tx, &distance, request.AddressID, request.WarehouseID)
	if err != nil {
		return nil, err
	}

	fmt.Println("2")

	totalPrice := distance
	if strings.Contains(vehicle.Name, "Bike") {
		totalPrice *= float64(BIKE)
	} else if strings.Contains(vehicle.Name, "Car") {
		totalPrice *= float64(CAR)
	} else {
		totalPrice *= float64(TRUCK)
	}

	currentTime := time.Now()
	date := currentTime.Format("20060102")

	var order entity.Order
	err = u.OrderRepository.GetLast(tx, &order)
	if err != nil {
		fmt.Println("errornya di sini nih")
		return nil, err
	}

	lastInvoiceNumber := strings.Split(order.InvoiceNumber, "/")[3]

	last, err := strconv.Atoi(lastInvoiceNumber)
	if err != nil {
		return nil, err
	}

	transactionID := fmt.Sprintf("INV/%s/PCK/%d", date, last+1)

	// store to database
	timeNow := time.Now().Local().UnixNano()
	newOrder := &entity.Order{
		ID:            uuid.NewString(),
		InvoiceNumber: transactionID,
		TotalAmount:   totalPrice,
		Weight:        request.Weight,
		Status:        "waiting-for-payment",
		PaymentType:   "",
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		AddressID:     request.AddressID,
		DriverID:      driver.ID,
	}

	fmt.Println("3")

	if voucher.ID != "" {
		newOrder.VoucherID = voucher.ID
	}

	err = u.OrderRepository.Create(tx, newOrder)
	if err != nil {
		fmt.Println("error di sini")
		return nil, err
	}

	midtransToken, err := midtrans.CreateToken(transactionID, int64(totalPrice))
	if err != nil {
		return nil, err
	}

	// commit
	err = tx.Commit()
	//err = errors.New("testing")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.ResponseCreatePayment{
		PaymentID: midtransToken,
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
	var totalWeight float64

	// check product IDs and quantities
	for i := range request.ProductIDs {
		q, err := strconv.Atoi(request.Quantities[i])
		if err != nil {
			return nil, err
		}

		if q <= 0 {
			return nil, fiber.ErrBadRequest
		}

		var product entity.Product
		err = u.ProductRepository.FindBy(tx, "id", request.ProductIDs[i], &product)
		if err != nil {
			return nil, err
		}

		if product.Quantity < q {
			return nil, fiber.ErrConflict
		}

		totalAmount += product.Price * float64(q)
		totalWeight += product.Weight * float64(q)

	}

	fmt.Println("passed")

	// check address ID
	var address entity.Address

	fmt.Println(request)

	err = u.AddressRepository.FindBy(tx, "id", request.AddressID, &address)
	if err != nil {
		fmt.Println("ini nih pasti")
		return nil, err
	}

	// check voucher ID
	var voucher entity.Voucher

	if request.VoucherID != "" {
		err = u.VoucherRepository.FindBy(tx, "id", request.VoucherID, &voucher, user)
		if err != nil {
			return nil, err
		}

		// voucher
		voucher.Status = false
		err = u.VoucherRepository.Update(tx, &voucher)
		if err != nil {
			return nil, err
		}

		fmt.Println(voucher)
	}
	fmt.Println("passed 2")

	// check Expedition
	var services []model.ServicesOngkir

	postalCode, err := strconv.Atoi(address.PostalCode)
	if err != nil {
		return nil, err
	}

	err = rajaongkir.CheckCost(postalCode, totalWeight, &services)
	if err != nil {
		return nil, err
	}

	fmt.Println("passed 3")

	ong := 0
	for _, s := range services {
		if s.Name == request.ExpeditionName && s.Service == request.ExpeditionType {
			ong = s.Cost
			break
		}
	}

	if ong == 0 {
		return nil, fiber.ErrBadRequest
	}

	// store to database
	currentTime := time.Now()
	date := currentTime.Format("20060102")

	var transaction entity.Transaction
	err = u.TransactionRepository.GetLast(tx, &transaction)
	if err != nil {
		fmt.Println("errornya di sini nih")
		return nil, err
	}

	lastInvoiceNumber := strings.Split(transaction.InvoiceNumber, "/")[3]

	last, err := strconv.Atoi(lastInvoiceNumber)
	if err != nil {
		return nil, err
	}

	transactionID := fmt.Sprintf("INV/%s/SHP/%d", date, last+1)

	timeNow := time.Now().Local().UnixNano()
	newTransaction := &entity.Transaction{
		ID:            uuid.NewString(),
		InvoiceNumber: transactionID,
		TotalAmount:   totalAmount,
		ShippingCosts: float64(ong),
		Status:        "waiting-for-payment",
		ServiceName:   request.ExpeditionName,
		ServiceType:   request.ExpeditionType,
		ReceiptNumber: "",
		PaymentType:   "",
		CreatedAt:     timeNow,
		UpdatedAt:     timeNow,
		AddressID:     request.AddressID,
	}

	if voucher.ID != "" {
		newTransaction.VoucherID = voucher.ID
	}

	err = u.TransactionRepository.Create(tx, newTransaction)
	if err != nil {
		fmt.Println("error di sini")
		return nil, err
	}

	// quantity product
	for i, p := range request.ProductIDs {
		var product entity.Product
		err = u.ProductRepository.FindBy(tx, "id", p, &product)
		if err != nil {
			return nil, err
		}

		q, err := strconv.Atoi(request.Quantities[i])
		if err != nil {
			return nil, err
		}

		product.Quantity -= q
		err = u.ProductRepository.Update(tx, &product)
		if err != nil {
			return nil, err
		}

		fmt.Printf("quantity updated: %d\n", product.Quantity)

		timeNow = time.Now().Local().UnixNano()
		transactionItem := &entity.TransactionItem{
			ID:            uuid.NewString(),
			Quantity:      q,
			TotalPrice:    product.Price * float64(q),
			TransactionID: newTransaction.ID,
			ProductID:     p,
			CreatedAt:     timeNow,
		}

		err = u.TransactionItemRepository.Create(tx, transactionItem)
		if err != nil {
			return nil, err
		}
	}

	// minus by voucher if exists
	if voucher.ID != "" {
		if voucher.IsPercent == true {
			totalAmount -= totalAmount * float64(voucher.Amount)
		} else {
			totalAmount -= float64(voucher.Amount)
		}
	}

	fmt.Println(transactionID)

	midtransToken, err := midtrans.CreateToken(transactionID, int64(totalAmount))
	if err != nil {
		return nil, err
	}

	// commit
	err = tx.Commit()
	//err = errors.New("testing")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.ResponseCreatePayment{
		PaymentID: midtransToken,
	}, nil
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
