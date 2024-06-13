package repository

import "github.com/jmoiron/sqlx"

type InterfaceVehiclesRepository interface {
	//
}

type VehiclesRepository struct {
	//
	DB *sqlx.DB
}

func NewVehiclesRepository(db *sqlx.DB) InterfaceMenuRepository {
	return &MenuRepository{
		DB: db,
	}
}
