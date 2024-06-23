package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceOrderRepository interface {
	//
	Update(tx *sqlx.Tx, order *entity.Order) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Order) error
	GetLast(tx *sqlx.Tx, order *entity.Order) error
	Create(tx *sqlx.Tx, order *entity.Order) error
}

type OrderRepository struct {
	//
	DB *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) InterfaceOrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (r *OrderRepository) Update(tx *sqlx.Tx, order *entity.Order) error {
	_, err := tx.NamedExec(query.ForUpdate(order), order)
	return err
}

func (r *OrderRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Order) error {
	q := fmt.Sprintf("SELECT * FROM orders WHERE %s = :value", column)
	param := map[string]any{
		"value": value,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Get(entity, param)
	if err != nil {
		return err
	}

	return err
}

func (r *OrderRepository) Create(tx *sqlx.Tx, order *entity.Order) error {

	_, err := tx.NamedExec(query.ForCreate(order), order)
	return err
}

func (r *OrderRepository) GetLast(tx *sqlx.Tx, order *entity.Order) error {
	q := `
SELECT
    o.invoice_number
FROM
    orders o
ORDER BY
    o.created_at DESC
    `

	err := tx.Get(order, q)
	if err != nil {
		return err
	}

	return nil
}
