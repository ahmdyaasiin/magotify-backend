package usecase

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/rajaongkir"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
)

type InterfacePaymentUseCase interface {
	//
	CreateShop(auth string, request *model.RequestCreatePayment) (*model.ResponseCreatePayment, error)
	Shop(auth string, productIds []string, addressID string, quantites []string) (*model.PaymentShop, error)
}

type PaymentUseCase struct {
	//
	DB                *sqlx.DB
	Log               *logrus.Logger
	PaymentRepository repository.InterfacePaymentRepository
	UserRepository    repository.InterfaceUserRepository
	AddressRepository repository.InterfaceAddressRepository
}

func NewPaymentUseCase(db *sqlx.DB, log *logrus.Logger, pr repository.InterfacePaymentRepository, ur repository.InterfaceUserRepository, ar repository.InterfaceAddressRepository) InterfacePaymentUseCase {
	//
	return &PaymentUseCase{
		DB:                db,
		Log:               log,
		PaymentRepository: pr,
		UserRepository:    ur,
		AddressRepository: ar,
	}
}

func (u *PaymentUseCase) CreateShop(auth string, request *model.RequestCreatePayment) (*model.ResponseCreatePayment, error) {

	return nil, nil
}

func (u *PaymentUseCase) Shop(auth string, productIds []string, addressID string, quantites []string) (*model.PaymentShop, error) {
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
	if err != nil {
		fmt.Println("1")
		return nil, err
	}

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
	}, nil
}
