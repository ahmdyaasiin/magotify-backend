package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceOrderRepository interface {
	//
	Update(tx *sqlx.Tx, order *entity.Order) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Order) error
	GetLast(tx *sqlx.Tx, order *entity.Order) error
	Create(tx *sqlx.Tx, order *entity.Order) error
	SpecificOrder(tx *sqlx.Tx, orderID string, dest *model.ResponseSpecificTransactionPickUp) error
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
	q := fmt.Sprintf("SELECT id, invoice_number, total_amount, weight, status, COALESCE(payment_type, '') as payment_type, created_at, updated_at, address_id, COALESCE(driver_id, '') as driver_id, COALESCE(voucher_id, '') as voucher_id FROM orders WHERE %s = :value", column)
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

func (r *OrderRepository) SpecificOrder(tx *sqlx.Tx, orderID string, dest *model.ResponseSpecificTransactionPickUp) error {
	q := `SELECT
    o.id as transaction_id,
    o.invoice_number,
    o.total_amount,
    o.weight, d.plate_number, v.name as vehicle_name,
    v.status as vehicle_status, w.name as warehouse_name, w.address as warehouse_address,
    v.duration as vehicle_duration, v.description as vehicle_description,
    a.name as address_name, a.address as user_address,
    ST_Distance_Sphere(
        POINT(w.longitude, w.latitude),
        POINT(a.longitude, a.latitude)
    ) AS distance_m
FROM
    orders o
JOIN
    drivers d ON o.driver_id = d.id
JOIN
    warehouses w ON d.warehouse_id = w.id
JOIN
    vehicles v ON d.vehicle_id = v.id
JOIN
    addresses a ON o.address_id = a.id
WHERE
    o.id = ?
    `

	err := tx.Get(dest, q, orderID)
	if err != nil {
		return err
	}

	return nil
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
