package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceWarehouseRepository interface {
	//
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Warehouse) error
}

type WarehouseRepository struct {
	//
	DB *sqlx.DB
}

func NewWarehouseRepository(db *sqlx.DB) InterfaceWarehouseRepository {
	return &WarehouseRepository{
		DB: db,
	}
}

func (r *WarehouseRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Warehouse) error {
	q := fmt.Sprintf("SELECT * FROM warehouses WHERE %s = :value", column)
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
