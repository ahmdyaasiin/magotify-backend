package repository

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/jmoiron/sqlx"
)

type InterfaceCategoryRepository interface {
	//
	GetALlCategoriesName(tx *sqlx.Tx, categoriesName *[]model.HotItemsSlice) error
}

type CategoryRepository struct {
	//
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) InterfaceCategoryRepository {
	return &CategoryRepository{
		DB: db,
	}
}

func (r *CategoryRepository) GetALlCategoriesName(tx *sqlx.Tx, categoriesName *[]model.HotItemsSlice) error {

	q := `
		SELECT c.name, c.url_photo FROM categories c;
	`

	err := tx.Select(categoriesName, q)
	if err != nil {
		return err
	}

	return nil
}
