package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceDriverRepository interface {
	//
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Driver) error
	FindAvailableDriver(tx *sqlx.Tx, warehouseID string, vehicleID string, dest *entity.Driver) error
	GetSpecificDriver(tx *sqlx.Tx, dest *entity.Vehicle, vehicleID string, warehouseID string) error
	GetDriverExcept(tx *sqlx.Tx, dest *[]entity.Vehicle, vehicleID string, warehouseID string) error
}

type DriverRepository struct {
	//
	DB *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) InterfaceDriverRepository {
	return &DriverRepository{
		DB: db,
	}
}

func (r *DriverRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Driver) error {
	q := fmt.Sprintf("SELECT * FROM drivers WHERE %s = :value", column)
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

func (r *DriverRepository) FindAvailableDriver(tx *sqlx.Tx, warehouseID string, vehicleID string, dest *entity.Driver) error {
	q := `
SELECT
    d.id, d.vehicle_id,
    d.created_at
FROM
    drivers d
LEFT JOIN
    orders o ON d.id = o.driver_id
WHERE
    d.warehouse_id = ? AND
    d.vehicle_id = ?
GROUP BY
    d.id, d.created_at
HAVING
    SUM(IF(o.status = 'in-progress' or o.status = 'waiting-for-payment', 1, 0)) = 0
ORDER BY
    COUNT(o.id), d.created_at;
	`

	err := tx.Get(dest, q, warehouseID, vehicleID)
	if err != nil {
		return err
	}

	return nil
}

func (r *DriverRepository) GetSpecificDriver(tx *sqlx.Tx, dest *entity.Vehicle, vehicleID string, warehouseID string) error {
	q := `
SELECT
    distinct v.id, v.name, v.duration, v.description, v.url_photo, v.status, v.created_at, v.updated_at
FROM
    drivers d
JOIN
    vehicles v ON d.vehicle_id = v.id
WHERE
    d.vehicle_id = ? AND d.warehouse_id = ?
`

	err := tx.Get(dest, q, vehicleID, warehouseID)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}

func (r *DriverRepository) GetDriverExcept(tx *sqlx.Tx, dest *[]entity.Vehicle, vehicleID string, warehouseID string) error {
	q := `
SELECT
    distinct v.id, v.name, v.duration, v.description, v.url_photo, v.status, v.created_at, v.updated_at
FROM
    drivers d
JOIN
    vehicles v ON d.vehicle_id = v.id
WHERE
    d.vehicle_id != ? AND d.warehouse_id = ?
`

	err := tx.Select(dest, q, vehicleID, warehouseID)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}
