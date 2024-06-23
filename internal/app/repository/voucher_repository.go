package repository

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/query"
	"github.com/jmoiron/sqlx"
)

type InterfaceVoucherRepository interface {
	//
	Update(tx *sqlx.Tx, voucher *entity.Voucher) error
	TotalVouchers(tx *sqlx.Tx, totalVoucher *int, user *entity.User) error
	FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Voucher, user *entity.User) error
	FindExcept(tx *sqlx.Tx, voucherID string, vouchers *[]entity.Voucher, user *entity.User) error
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

func (r *VoucherRepository) Update(tx *sqlx.Tx, voucher *entity.Voucher) error {
	_, err := tx.NamedExec(query.ForUpdate(voucher), voucher)
	return err
}

func (r *VoucherRepository) TotalVouchers(tx *sqlx.Tx, totalVoucher *int, user *entity.User) error {
	q := "SELECT COUNT(*) FROM vouchers WHERE user_id = :user_id AND status = 1"
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

func (r *VoucherRepository) FindBy(tx *sqlx.Tx, column string, value string, entity *entity.Voucher, user *entity.User) error {
	q := fmt.Sprintf("SELECT * FROM vouchers WHERE %s = :value AND user_id = :user_id AND status = 1", column)
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

func (r *VoucherRepository) FindExcept(tx *sqlx.Tx, voucherID string, vouchers *[]entity.Voucher, user *entity.User) error {
	q := "SELECT * FROM vouchers WHERE id != :voucher_id AND user_id = :user_id AND status = 1"
	param := map[string]any{
		"voucher_id": voucherID,
		"user_id":    user.ID,
	}

	stmt, err := tx.PrepareNamed(q)
	if err != nil {
		return err
	}

	err = stmt.Select(vouchers, param)
	if err != nil {
		return err
	}

	return nil
}
