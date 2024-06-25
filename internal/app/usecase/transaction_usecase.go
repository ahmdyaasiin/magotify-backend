package usecase

import (
	"fmt"
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
	UpdateExpiredTransaction()
	UpdateExpiredOrder()
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

func (u *TransactionUseCase) UpdateExpiredTransaction() {

	fmt.Println("go cron jalan nih cioks [transaction]")

	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		fmt.Println("go cron error make transaction [transaction]")
		return
	}

	err = u.TransactionRepository.UpdateStatusExpiredTransaction(tx)
	if err != nil {
		fmt.Println("go cron error exec repo [transaction]")
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error [t]: " + err.Error())
		return
	}
}

func (u *TransactionUseCase) UpdateExpiredOrder() {

	fmt.Println("go cron jalan nih cioks [order]")

	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		fmt.Println("go cron error make transaction [order]")
		return
	}

	err = u.TransactionRepository.UpdateStatusExpiredOrder(tx)
	if err != nil {
		fmt.Println(err)
		fmt.Println("go cron error exec repo [order]")
		return
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("commit error [o]: " + err.Error())
		return
	}
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
