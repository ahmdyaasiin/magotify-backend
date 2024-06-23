package repository

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceTransactionItemRepository interface {
	//
	Create(tx *sqlx.Tx, transaction *entity.TransactionItem) error
}

type TransactionItemRepository struct {
	//
	DB *sqlx.DB
}

func NewTransactionItemRepository(db *sqlx.DB) InterfaceTransactionItemRepository {
	return &TransactionItemRepository{
		DB: db,
	}
}

func (r *TransactionItemRepository) Create(tx *sqlx.Tx, transaction *entity.TransactionItem) error {
	_, err := tx.NamedExec(query.ForCreate(transaction), transaction)
	return err
}
