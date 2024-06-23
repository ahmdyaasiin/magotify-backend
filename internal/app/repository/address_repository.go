package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceAddressRepository interface {
	//
	GetDistance(tx *sqlx.Tx, dest *float64, addressID string, warehouseID string) error
	Create(tx *sqlx.Tx, address *entity.Address) error
	GetPrimaryAddress(tx *sqlx.Tx, address *entity.Address, user *entity.User) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Address) error
	FindAddressUser(tx *sqlx.Tx, column string, value string, entity *entity.Address, user *entity.User) error
	FindExcept(tx *sqlx.Tx, addressID string, addresses *[]entity.Address, user *entity.User) error
}

type AddressRepository struct {
	//
	DB *sqlx.DB
}

func NewAddressRepository(db *sqlx.DB) InterfaceAddressRepository {
	return &AddressRepository{
		DB: db,
	}
}

func (r *AddressRepository) Create(tx *sqlx.Tx, address *entity.Address) error {
	_, err := tx.NamedExec(query.ForCreate(address), address)
	return err
}

func (r *AddressRepository) GetPrimaryAddress(tx *sqlx.Tx, address *entity.Address, user *entity.User) error {
	q := "SELECT * FROM addresses WHERE is_primary = 1 AND user_id = :user_id"
	param := map[string]any{
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Get(address, param)
	if err != nil {
		return err
	}

	return nil
}

func (r *AddressRepository) GetDistance(tx *sqlx.Tx, dest *float64, addressID string, warehouseID string) error {
	q := `
SELECT
    ST_Distance_Sphere(
        POINT(w.longitude, w.latitude),
        POINT(a.longitude, a.latitude)
    ) AS distance_m
FROM
    addresses a
JOIN
    warehouses w ON w.id = ?
WHERE
    a.id = ?
	`

	err := tx.Get(dest, q, addressID, warehouseID)
	if err != nil {
		return err
	}

	return nil

}

func (r *AddressRepository) FindExcept(tx *sqlx.Tx, addressID string, addresses *[]entity.Address, user *entity.User) error {
	q := "SELECT * FROM addresses WHERE id != :address_id AND user_id = :user_id"
	param := map[string]any{
		"address_id": addressID,
		"user_id":    user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Select(addresses, param)
	if err != nil {
		return err
	}

	return nil
}

func (r *AddressRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Address) error {
	q := fmt.Sprintf("SELECT * FROM addresses WHERE %s = :value", column)
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

func (r *AddressRepository) FindAddressUser(tx *sqlx.Tx, column string, value string, entity *entity.Address, user *entity.User) error {
	q := fmt.Sprintf("SELECT * FROM addresses WHERE %s = :value AND user_id = :user_id", column)
	param := map[string]any{
		"value":   value,
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Get(entity, param)
	if err != nil {
		return err
	}

	return nil
}
