package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceBannerRepository interface {
	//
	GetAllBanner(tx *sqlx.Tx, banners *[]entity.Banner) error
}

type BannerRepository struct {
	//
	DB *sqlx.DB
}

func NewBannerRepository(db *sqlx.DB) InterfaceBannerRepository {
	return &BannerRepository{
		DB: db,
	}
}

func (r *BannerRepository) GetAllBanner(tx *sqlx.Tx, banners *[]entity.Banner) error {
	q := `SELECT * FROM banners`

	err := tx.Select(banners, q)
	if err != nil {
		fmt.Println("error nih: " + err.Error())
		return err
	}

	return nil
}
