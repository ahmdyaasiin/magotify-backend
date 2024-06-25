package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceTransactionRepository interface {
	//
	Update(tx *sqlx.Tx, transaction *entity.Transaction) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Transaction) error
	Create(tx *sqlx.Tx, transaction *entity.Transaction) error
	GetLast(tx *sqlx.Tx, transaction *entity.Transaction) error
	TransactionShop(tx *sqlx.Tx, user *entity.User, entity *[]model.ResponseTransactionShop) error
	TransactionPickUp(tx *sqlx.Tx, user *entity.User, entity *[]model.ResponseTransactionPickUp) error
	SpecificTransaction(tx *sqlx.Tx, transactionID string, dest *model.ResponseSpecificTransactionShop) error
}

type TransactionRepository struct {
	//
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) InterfaceTransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (r *TransactionRepository) TransactionPickUp(tx *sqlx.Tx, user *entity.User, entity *[]model.ResponseTransactionPickUp) error {
	q := `
SELECT
    o.id, d.name, a.address, o.total_amount, o.weight, o.created_at,
    ST_Distance_Sphere(
        POINT(w.longitude, w.latitude),
        POINT(a.longitude, a.latitude)
    ) AS distance_m
FROM
    orders o
JOIN
    addresses a ON o.address_id = a.id
JOIN
    drivers d ON o.driver_id = d.id
JOIN
    warehouses w ON d.warehouse_id = w.id
WHERE
    a.user_id = ?
    `

	err := tx.Select(entity, q, user.ID)
	if err != nil {
		return err
	}

	return err
}

func (r *TransactionRepository) TransactionShop(tx *sqlx.Tx, user *entity.User, entity *[]model.ResponseTransactionShop) error {
	q := `
SELECT
    t.id, t.invoice_number,
    a.name as address_name,
    (t.total_amount+t.shipping_costs) as total_price,
    t.created_at,
    count(ti.id)-1 as total_products,
    sum(p.weight*ti.quantity) as total_weight
FROM
    transactions t
JOIN
    addresses a ON t.address_id = a.id
JOIN
    transaction_items ti ON t.id = ti.transaction_id
JOIN
    products p ON ti.product_id = p.id
WHERE
    a.user_id = ?
GROUP BY
    t.invoice_number, a.name, (t.total_amount+t.shipping_costs), t.created_at
    `

	err := tx.Select(entity, q, user.ID)
	if err != nil {
		return err
	}

	return err
}

func (r *TransactionRepository) Update(tx *sqlx.Tx, transaction *entity.Transaction) error {
	_, err := tx.NamedExec(query.ForUpdate(transaction), transaction)
	return err
}

func (r *TransactionRepository) SpecificTransaction(tx *sqlx.Tx, transactionID string, dest *model.ResponseSpecificTransactionShop) error {
	q := `
SELECT
    t.id as transaction_id, t.invoice_number, t.total_amount, t.shipping_costs,
    t.status, t.service_type, t.service_name, COALESCE(t.receipt_number, '') as receipt_number, COALESCE(t.payment_type, '') as payment_type,
    t.created_at, a.address, a.name as address_name, COALESCE(v.id, '') as voucher_id, COALESCE(v.name, '') as voucher_name,
    COALESCE(v.amount, '') as voucher_amount, COALESCE(v.is_percent, '') as voucher_is_percent, COALESCE(v.url_logo, '') as voucher_url_logo
FROM
    transactions t
JOIN
    addresses a ON t.address_id = a.id
LEFT JOIN
    vouchers v ON t.voucher_id = v.id
WHERE
    t.id = ?
	`

	err := tx.Get(dest, q, transactionID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Transaction) error {
	q := `
SELECT
    t.id,
    t.invoice_number,
    t.total_amount,
    t.shipping_costs,
    t.status,
    t.service_name,
    t.service_type,
    t.created_at,
    t.updated_at,
    t.address_id,
    COALESCE(t.receipt_number, '') as receipt_number,
    COALESCE(t.payment_type, '') as payment_type,
    COALESCE(t.voucher_id, '') as voucher_id
FROM
    transactions t
WHERE %s = :value
    `
	q = fmt.Sprintf(q, column)
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

func (r *TransactionRepository) Create(tx *sqlx.Tx, transaction *entity.Transaction) error {

	fmt.Println(query.ForCreate(transaction))

	_, err := tx.NamedExec(query.ForCreate(transaction), transaction)
	return err
}

func (r *TransactionRepository) GetLast(tx *sqlx.Tx, transaction *entity.Transaction) error {
	q := "SELECT id, invoice_number, total_amount, shipping_costs, status, service_name, service_type, IFNULL(receipt_number, '') AS receipt_number, IFNULL(payment_type, '') AS payment_type, created_at, updated_at, address_id, IFNULL(voucher_id, '') AS voucher_id FROM transactions ORDER BY created_at DESC;"

	err := tx.Get(transaction, q)
	if err != nil {
		return err
	}

	return nil
}
