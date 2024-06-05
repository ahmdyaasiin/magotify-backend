package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type Repository[T any] struct {
	DB *sqlx.DB
}

func (r *Repository[T]) Create(db *sqlx.Tx, entity query.TableInterface) error {
	fmt.Println(query.ForCreate(entity))
	return nil
}
