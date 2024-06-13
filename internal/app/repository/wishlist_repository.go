package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceWishlistRepository interface {
	//
	Create(tx *sqlx.Tx, wishlist *entity.Wishlist) error
	Delete(tx *sqlx.Tx, wishlist *entity.Wishlist) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Wishlist, user *entity.User) error
	MyWishlist(tx *sqlx.Tx, wishlists *[]model.ExploreItems, user *entity.User) error
}

type WishlistRepository struct {
	//
	DB *sqlx.DB
}

func NewWishlistRepository(db *sqlx.DB) InterfaceWishlistRepository {
	return &WishlistRepository{
		DB: db,
	}
}

func (r *WishlistRepository) Create(tx *sqlx.Tx, wishlist *entity.Wishlist) error {
	_, err := tx.NamedExec(query.ForCreate(wishlist), wishlist)
	return err
}

func (r *WishlistRepository) Delete(tx *sqlx.Tx, wishlist *entity.Wishlist) error {
	_, err := tx.NamedExec(query.ForDelete(wishlist), wishlist)
	return err
}

func (r *WishlistRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Wishlist, user *entity.User) error {
	q := fmt.Sprintf("SELECT * FROM wishlists WHERE %s = :value AND user_id = :user_id", column)
	param := map[string]any{
		"value":   value,
		"user_id": user.ID,
	}

	fmt.Println(q)

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

func (r *WishlistRepository) MyWishlist(tx *sqlx.Tx, wishlists *[]model.ExploreItems, user *entity.User) error {
	q := `
SELECT
    p.id as product_id,
    p.name,
    p.price as real_price,
    p.price * (1 - p.discount_percentage / 100) AS discount_price,
    1 as is_wishlist,
    (SELECT url_photo FROM media m WHERE m.product_id = p.id ORDER BY m.url_photo LIMIT 1) AS product_image,
    COALESCE(AVG(r.star), 0) AS average_rating,
    COALESCE(SUM(ti.quantity), 0) AS sold
FROM
    wishlists w
JOIN
    products p ON w.product_id = p.id
LEFT JOIN
    transaction_items ti ON ti.product_id = p.id
LEFT JOIN
    ratings r ON r.transaction_item_id = ti.id
WHERE
    w.user_id = :user_id
GROUP BY
    p.id, p.name, p.price
    `
	param := map[string]any{
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Select(wishlists, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
