package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type InterfaceMenuRepository interface {
	//
	HotItemsGeneral(tx *sqlx.Tx, dest *[]model.ExploreItems, user_id string) error
	HotItemsSpecific(tx *sqlx.Tx, dest *[]model.ExploreItems, categoryName string, user_id string) error
	TimeToCleanUp(tx *sqlx.Tx, dest *[]model.ExploreTimeToCleanUp, user_id string) error
	ProductBestOffer(tx *sqlx.Tx, products *[]model.ExploreItems) error
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

func (r *MenuRepository) HotItemsGeneral(tx *sqlx.Tx, dest *[]model.ExploreItems, userID string) error {
	q := `
SELECT
    ps.product_id,
    ps.product_image,
    ps.name,
    ps.real_price,
    ps.discount_price,
    COALESCE(pr.average_rating, 0) AS average_rating,
    ps.sold,
    COALESCE(w.is_wishlist, 0) AS is_wishlist
FROM
    (
        SELECT
            ti.product_id,
            (SELECT url_photo FROM media m WHERE m.product_id = ti.product_id ORDER BY m.url_photo LIMIT 1) as product_image,
            p.name,
            p.price AS real_price,
            p.price * (1 - p.discount_percentage / 100) AS discount_price,
            COUNT(ti.quantity) AS sold,
            p.category_id
        FROM
            transaction_items ti
        JOIN
            products p ON ti.product_id = p.id
        GROUP BY
            ti.product_id, p.name, p.price, p.discount_percentage, p.category_id
    ) ps
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
    ) pr ON ps.product_id = pr.product_id
LEFT JOIN
    (
        SELECT
            user_id,
            product_id,
            1 AS is_wishlist
        FROM
            wishlists
    ) w ON ps.product_id = w.product_id AND w.user_id = ?
ORDER BY
    ps.sold DESC;
	`

	err := tx.Select(dest, q, userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *MenuRepository) HotItemsSpecific(tx *sqlx.Tx, dest *[]model.ExploreItems, categoryName string, user_id string) error {
	q := `
SELECT
    ps.product_id,
    ps.product_image,
    ps.name,
    ps.real_price,
    ps.discount_price,
    COALESCE(pr.average_rating, 0) AS average_rating,
    ps.sold,
    COALESCE(w.is_wishlist, 0) AS is_wishlist
FROM
    (
        SELECT
            ti.product_id,
            (SELECT url_photo FROM media m WHERE m.product_id = ti.product_id ORDER BY m.url_photo LIMIT 1) as product_image,
            p.name,
            p.price AS real_price,
            p.price * (1 - p.discount_percentage / 100) AS discount_price,
            COUNT(ti.quantity) AS sold,
            p.category_id
        FROM
            transaction_items ti
        JOIN
            products p ON ti.product_id = p.id
        WHERE
            p.category_id IN (SELECT id FROM categories WHERE name = ?)
        GROUP BY
            ti.product_id, p.name, p.price, p.discount_percentage, p.category_id
    ) ps
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
    ) pr ON ps.product_id = pr.product_id
LEFT JOIN
    (
        SELECT
            user_id,
            product_id,
            1 AS is_wishlist
        FROM
            wishlists
    ) w ON ps.product_id = w.product_id AND w.user_id = ?
ORDER BY
    ps.sold DESC;
	`

	err := tx.Select(dest, q, categoryName, user_id)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}

func (r *MenuRepository) TimeToCleanUp(tx *sqlx.Tx, dest *[]model.ExploreTimeToCleanUp, user_id string) error {
	q := `
SELECT
    v.id,
    v.name,
    v.url_photo,
    v.description,
    v.status,
    w.name as warehouse_name,
    ST_Distance_Sphere(
        POINT(w.longitude, w.latitude),
        POINT(a.longitude, a.latitude)
    ) AS distance_m
FROM
    vehicles v
JOIN
    warehouses w ON v.warehouse_id = w.id
JOIN
    addresses a ON a.user_id = ? AND a.is_primary = 1
ORDER BY
    distance_m;
	`

	err := tx.Select(dest, q, user_id)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}

func (r *MenuRepository) ProductBestOffer(tx *sqlx.Tx, products *[]model.ExploreItems) error {
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
    p.discount_percentage != 0
ORDER BY
    p.discount_percentage DESC;
	`

	err := tx.Select(products, q)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}
