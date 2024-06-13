package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceCartRepository interface {
	//
	Update(tx *sqlx.Tx, cart *entity.Cart) error
	Create(tx *sqlx.Tx, cart *entity.Cart) error
	Delete(tx *sqlx.Tx, cart *entity.Cart) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Cart, user *entity.User) error
	CountCart(tx *sqlx.Tx, total *int, user *entity.User) error
	MyCart(tx *sqlx.Tx, myCart *[]model.ProductCart, user *entity.User) error
}

type CartRepository struct {
	//
	DB *sqlx.DB
}

func NewCartRepository(db *sqlx.DB) InterfaceCartRepository {
	return &CartRepository{
		DB: db,
	}
}

func (r *CartRepository) Create(tx *sqlx.Tx, cart *entity.Cart) error {
	_, err := tx.NamedExec(query.ForCreate(cart), cart)
	return err
}

func (r *CartRepository) Update(tx *sqlx.Tx, cart *entity.Cart) error {
	_, err := tx.NamedExec(query.ForUpdate(cart), cart)
	return err
}

func (r *CartRepository) Delete(tx *sqlx.Tx, cart *entity.Cart) error {
	_, err := tx.NamedExec(query.ForDelete(cart), cart)
	return err
}

func (r *CartRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Cart, user *entity.User) error {
	q := fmt.Sprintf("SELECT * FROM carts WHERE %s = :value AND user_id = :user_id", column)
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

	return err
}

func (r *CartRepository) CountCart(tx *sqlx.Tx, total *int, user *entity.User) error {
	q := `SELECT COUNT(*) FROM carts WHERE user_id = :user_id`
	param := map[string]any{
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Get(total, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}

func (r *CartRepository) MyCart(tx *sqlx.Tx, myCart *[]model.ProductCart, user *entity.User) error {
	q := `
    SELECT
        p.id,
        p.name as name,
        p.price,
        c.quantity,
        (SELECT url_photo FROM media m WHERE m.product_id = p.id ORDER BY m.url_photo LIMIT 1) as url_photo,
        cat.name as cat_name
    FROM
        carts c
    JOIN
        products p ON c.product_id = p.id
    JOIN
        categories cat ON p.category_id = cat.id
    WHERE
        c.user_id = :user_id
    `

	param := map[string]any{
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Select(myCart, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
