package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceDriverRepository interface {
	//
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
