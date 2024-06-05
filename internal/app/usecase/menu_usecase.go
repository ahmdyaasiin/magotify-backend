package usecase

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type InterfaceMenuUseCase interface {
	//
}

type MenuUseCase struct {
	//
	DB             *sqlx.DB
	Log            *logrus.Logger
	MenuRepository repository.InterfaceMenuRepository
}

func NewMenuUseCase(db *sqlx.DB, log *logrus.Logger, mr repository.InterfaceMenuRepository) InterfaceMenuUseCase {
	//
	return &MenuUseCase{
		DB:             db,
		Log:            log,
		MenuRepository: mr,
	}
}
