package repository

import (
	"github.com/jmoiron/sqlx"
)

type InterfaceMenuRepository interface {
	//
}

type MenuRepository struct {
	//
	DB *sqlx.DB
}

func NewMenuRepository(db *sqlx.DB) InterfaceMenuRepository {
	return &MenuRepository{
		DB: db,
	}
}
