package usecase

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type InterfaceTransactionUseCase interface {
	//
	HistoryShop(auth string) (*[]model.ResponseTransactionShop, error)
	HistoryPickUp(auth string) (*[]model.ResponseTransactionPickUp, error)
	SpecificPickUp(auth string, transactionID string) (*model.ResponseSpecificTransactionPickUp, error)
	SpecificShop(auth string, transactionID string) (*model.ResponseSpecificTransactionShop, error)
}

type TransactionUseCase struct {
	//
	DB                    *sqlx.DB
	Log                   *logrus.Logger
	TransactionRepository repository.InterfaceTransactionRepository
	UserRepository        repository.InterfaceUserRepository
	OrderRepository       repository.InterfaceOrderRepository
	PaymentRepository     repository.InterfacePaymentRepository
}

func NewTransactionUseCase(db *sqlx.DB, log *logrus.Logger, tr repository.InterfaceTransactionRepository, ur repository.InterfaceUserRepository, or repository.InterfaceOrderRepository, pr repository.InterfacePaymentRepository) InterfaceTransactionUseCase {
	//
	return &TransactionUseCase{
		DB:                    db,
		Log:                   log,
		TransactionRepository: tr,
		UserRepository:        ur,
		OrderRepository:       or,
		PaymentRepository:     pr,
	}
}

func (u *TransactionUseCase) SpecificPickUp(auth string, transactionID string) (*model.ResponseSpecificTransactionPickUp, error) {
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

	mod := new(model.ResponseSpecificTransactionPickUp)
	err = u.OrderRepository.SpecificOrder(tx, transactionID, mod)
	if err != nil {
		return nil, err
	}

	return mod, nil
}

func (u *TransactionUseCase) SpecificShop(auth string, transactionID string) (*model.ResponseSpecificTransactionShop, error) {
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

	transaction := new(model.ResponseSpecificTransactionShop)

	err = u.TransactionRepository.SpecificTransaction(tx, transactionID, transaction)
	if err != nil {
		return nil, err
	}

	// add products for transaction variable
	err = u.PaymentRepository.ProductsForTransactionDetails(tx, transactionID, &transaction.Products)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *TransactionUseCase) HistoryPickUp(auth string) (*[]model.ResponseTransactionPickUp, error) {
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

	var orders []model.ResponseTransactionPickUp
	err = u.TransactionRepository.TransactionPickUp(tx, user, &orders)
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (u *TransactionUseCase) HistoryShop(auth string) (*[]model.ResponseTransactionShop, error) {
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

	var transaction []model.ResponseTransactionShop
	err = u.TransactionRepository.TransactionShop(tx, user, &transaction)
	if err != nil {
		return nil, err
	}

	// first product
	// first image of product

	return &transaction, nil
}
