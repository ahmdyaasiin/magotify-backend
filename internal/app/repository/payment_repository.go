package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
	"strings"
)

type InterfacePaymentRepository interface {
	//
	ProductsPaymentShop(tx *sqlx.Tx, products *[]model.ProductCart, productIds []string) error
}

type PaymentRepository struct {
	//
	DB *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) InterfacePaymentRepository {
	return &PaymentRepository{
		DB: db,
	}
}

func (r *PaymentRepository) ProductsPaymentShop(tx *sqlx.Tx, products *[]model.ProductCart, productIds []string) error {
	q := fmt.Sprintf(`
SELECT
    p.id,
    p.name as name,
    p.price,
    p.weight,
    (SELECT url_photo FROM media m WHERE m.product_id = p.id ORDER BY m.url_photo LIMIT 1) as url_photo,
    p.price * (1 - p.discount_percentage / 100) AS discount_price,
    c.name as cat_name
FROM
    products p
JOIN
    categories c ON p.category_id = c.id
WHERE
    p.id IN ('%s')
	`, strings.Join(productIds, "', '"))

	err := tx.Select(products, q)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
