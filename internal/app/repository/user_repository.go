package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceUserRepository interface {
	//
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.User) error
	Create(db *sqlx.Tx, user *entity.User) error
}

type UserRepository struct {
	//
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) InterfaceUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.User) error {
	q := fmt.Sprintf("SELECT * FROM users WHERE %s = :value", column)
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

func (r *UserRepository) Create(tx *sqlx.Tx, user *entity.User) error {
	_, err := tx.NamedExec(query.ForCreate(user), user)
	return err
}
