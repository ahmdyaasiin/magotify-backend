package repository

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/jmoiron/sqlx"
)

type InterfaceVoucherRepository interface {
	//
	TotalVouchers(tx *sqlx.Tx, totalVoucher *int, user *entity.User) error
}

type VoucherRepository struct {
	//
	DB *sqlx.DB
}

func NewVoucherRepository(db *sqlx.DB) InterfaceVoucherRepository {
	return &VoucherRepository{
		DB: db,
	}
}

func (r VoucherRepository) TotalVouchers(tx *sqlx.Tx, totalVoucher *int, user *entity.User) error {
	q := "SELECT COUNT(*) FROM vouchers WHERE user_id = :user_id"
	param := map[string]any{
		"user_id": user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Get(totalVoucher, param)
	if err != nil {
		return err
	}

	return nil
}
