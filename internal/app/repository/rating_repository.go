package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type InterfaceRatingRepository interface {
	//
	GetAllRatings(tx *sqlx.Tx, ratings *[]model.ReviewProduct, productID string) error
}

type RatingRepository struct {
	//
	DB *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) InterfaceRatingRepository {
	return &RatingRepository{
		DB: db,
	}
}

func (r *RatingRepository) GetAllRatings(tx *sqlx.Tx, ratings *[]model.ReviewProduct, productID string) error {
	q := `SELECT
    u.name,
    u.url_photo,
    r.star,
    r.content,
    r.created_at
FROM
    ratings r
JOIN
    transaction_items ti ON r.transaction_item_id = ti.id
JOIN
    products p ON ti.product_id = p.id
JOIN
    users u ON r.user_id = u.id
WHERE
    p.id = :product_id
ORDER BY
    r.created_at DESC
    `

	param := map[string]any{
		"product_id": productID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Select(ratings, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
