package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type InterfaceMediaRepository interface {
	//
	GetAllMediaProduct(tx *sqlx.Tx, media *[]model.MediaProduct, productID string) error
}

type MediaRepository struct {
	//
	DB *sqlx.DB
}

func NewMediaRepository(db *sqlx.DB) InterfaceMediaRepository {
	return &MediaRepository{
		DB: db,
	}
}

func (r *MediaRepository) GetAllMediaProduct(tx *sqlx.Tx, media *[]model.MediaProduct, productID string) error {
	q := `SELECT id, url_photo FROM media WHERE product_id = :product_id`

	param := map[string]any{
		"product_id": productID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		fmt.Println("error 1: " + err.Error())
		return err
	}

	err = stmt.Select(media, param)
	if err != nil {
		fmt.Println("error 2: " + err.Error())
		return err
	}

	return nil
}
