package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceVehicleRepository interface {
	//
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Vehicle) error
	FindExcept(tx *sqlx.Tx, vehicleID string, vehicles *[]entity.Vehicle, warehouse *entity.Warehouse) error
}

type VehicleRepository struct {
	//
	DB *sqlx.DB
}

func NewVehicleRepository(db *sqlx.DB) InterfaceVehicleRepository {
	return &VehicleRepository{
		DB: db,
	}
}

func (r *VehicleRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Vehicle) error {
	q := fmt.Sprintf("SELECT * FROM vehicles WHERE %s = :value", column)
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

func (r *VehicleRepository) FindExcept(tx *sqlx.Tx, vehicleID string, vehicles *[]entity.Vehicle, warehouse *entity.Warehouse) error {
	q := "SELECT * FROM vehicles WHERE id != :vehicle_id AND warehouse_id = :warehouse_id"
	param := map[string]any{
		"vehicle_id":   vehicleID,
		"warehouse_id": warehouse.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Select(vehicles, param)
	if err != nil {
		return err
	}

	return nil
}
