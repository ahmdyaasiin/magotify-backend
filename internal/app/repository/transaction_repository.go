package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
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
	UpdateStatusExpiredTransaction(tx *sqlx.Tx) error
	UpdateStatusExpiredOrder(tx *sqlx.Tx) error
	UpdateSpecificExpiredOrder(tx *sqlx.Tx, orderID string) error
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
    t.id,
    t.invoice_number,
    a.name AS address_name,
    (t.total_amount + t.shipping_costs) AS total_price,
    t.created_at,
    (
        SELECT ti2.quantity
        FROM transaction_items ti2
        WHERE ti2.transaction_id = t.id
        ORDER BY ti2.created_at
        LIMIT 1
    ) AS product_quantity,
    (
        SELECT p.name
        FROM transaction_items ti2
        JOIN products p ON ti2.product_id = p.id
        WHERE ti2.transaction_id = t.id
        ORDER BY ti2.created_at
        LIMIT 1
    ) AS product_name,
    (
        SELECT m.url_photo
        FROM transaction_items ti2
        JOIN products p ON ti2.product_id = p.id
        JOIN media m ON p.id = m.product_id
        WHERE ti2.transaction_id = t.id
        ORDER BY ti2.created_at
        LIMIT 1
    ) AS product_image,
    COUNT(ti.id) - 1 AS total_products,
    SUM(p.weight * ti.quantity) AS total_weight
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
    t.id, t.invoice_number, a.name, (t.total_amount + t.shipping_costs), t.created_at
ORDER BY
    t.created_at DESC;
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

func (r *TransactionRepository) UpdateStatusExpiredTransaction(tx *sqlx.Tx) error {
	fiveMinutesAgo := time.Now().Local().Add(-1 * 5 * time.Minute).UnixNano()

	q := `
		UPDATE transactions t SET t.status='cancel' WHERE t.status = 'waiting-for-payment' AND ? > t.created_at
	`

	res, err := tx.Exec(q, fiveMinutesAgo)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("rows affected nih [t]: " + strconv.FormatInt(rowsAffected, 10))

	return nil
}

func (r *TransactionRepository) UpdateSpecificExpiredOrder(tx *sqlx.Tx, orderID string) error {
	q := `
UPDATE orders o SET o.status='cancel', o.driver_id = null WHERE o.id = ?
	`

	_, err := tx.Exec(q, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (r *TransactionRepository) UpdateStatusExpiredOrder(tx *sqlx.Tx) error {
	fiveMinutesAgo := time.Now().Local().Add(-1 * 5 * time.Minute).UnixNano()

	q := `
		UPDATE orders o SET o.status='cancel', o.driver_id = null WHERE o.status = 'waiting-for-payment' AND ? > o.created_at
	`

	res, err := tx.Exec(q, fiveMinutesAgo)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println("rows affected nih [o]: " + strconv.FormatInt(rowsAffected, 10))

	return nil
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
