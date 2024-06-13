package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type InterfaceProductRepository interface {
	//
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Product) error
	ProductDetails(tx *sqlx.Tx, details *model.PD, user *entity.User, productID string) error
	ProductBestOfferWithout(tx *sqlx.Tx, products *[]model.ExploreItems, productID string) error
}

type ProductRepository struct {
	//
	DB *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) InterfaceProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (r *ProductRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Product) error {
	q := fmt.Sprintf("SELECT * FROM products WHERE %s = :value", column)
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

func (r *ProductRepository) ProductDetails(tx *sqlx.Tx, details *model.PD, user *entity.User, productID string) error {
	q := `SELECT
    p.id,
    p.name,
    p.description,
    p.price,
    p.price * (1 - p.discount_percentage / 100) AS discount_price,
    COALESCE(c.quantity, 0) AS cart_quantity,
    COALESCE(pr.average_rating, 0) as average_rating,
    COALESCE(w.is_wishlist, 0) AS is_wishlist,
    COALESCE(ti.total_sold, 0) AS total_sold
FROM
    products p
LEFT JOIN
    (
        SELECT
            ti.product_id,
            AVG(r.star) AS average_rating
        FROM
            ratings r
        JOIN
            transaction_items ti ON r.transaction_item_id = ti.id
        GROUP BY
            ti.product_id
    ) pr ON p.id = pr.product_id
LEFT JOIN
    (
        SELECT
            user_id,
            product_id,
            1 AS is_wishlist
        FROM
            wishlists
    ) w ON p.id = w.product_id AND w.user_id = :user_id
LEFT JOIN
    carts c ON p.id = c.product_id AND c.user_id = :user_id
LEFT JOIN
    (
        SELECT
            ti.product_id,
            SUM(ti.quantity) AS total_sold
        FROM
            transaction_items ti
        GROUP BY
            ti.product_id
    ) ti ON p.id = ti.product_id
WHERE
    p.id = :product_id;
`
	param := map[string]any{
		"user_id":    user.ID,
		"product_id": productID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Get(details, param)
	if err != nil {
		fmt.Println("error 2x: " + err.Error())
		return err
	}

	return nil
}

func (r *ProductRepository) ProductBestOfferWithout(tx *sqlx.Tx, products *[]model.ExploreItems, productID string) error {
	q := `
SELECT
    p.id as product_id,
    p.name,
    p.price as real_price,
    p.price * (1 - p.discount_percentage / 100) AS discount_price,
    (SELECT url_photo FROM media m WHERE m.product_id = p.id ORDER BY m.url_photo LIMIT 1) as product_image,
    COALESCE(pr.average_rating, 0) AS average_rating,
    COALESCE(ti.total_sold, 0) AS sold
FROM
    products p
LEFT JOIN
    (
        SELECT
            ti.product_id,
            AVG(r.star) AS average_rating
        FROM
            ratings r
        JOIN
            transaction_items ti ON r.transaction_item_id = ti.id
        GROUP BY
            ti.product_id
    ) pr ON p.id = pr.product_id
LEFT JOIN
    (
        SELECT
            ti.product_id,
            SUM(ti.quantity) AS total_sold
        FROM
            transaction_items ti
        GROUP BY
            ti.product_id
    ) ti ON p.id = ti.product_id
WHERE
    p.discount_percentage != 0 AND p.id != :product_id
ORDER BY
    p.discount_percentage DESC;
	`

	param := map[string]any{
		"product_id": productID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Select(products, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
