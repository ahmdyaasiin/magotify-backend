package usecase

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/ahmdyaasiin/magotify-backend/internal/pkg/firebase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type InterfaceUserUseCase interface {
	//
	Create(request *model.RequestUserRegister) (*model.ResponseUser, error)
	Login(request *model.RequestUserLogin) (*model.ResponseUser, error)
	Verify(request *model.UUIDMiddleware) error
}

type UserUseCase struct {
	//
	DB                *sqlx.DB
	Log               *logrus.Logger
	UserRepository    repository.InterfaceUserRepository
	AddressRepository repository.InterfaceAddressRepository
}

func NewUserUseCase(db *sqlx.DB, log *logrus.Logger, ur repository.InterfaceUserRepository, ar repository.InterfaceAddressRepository) InterfaceUserUseCase {
	//
	return &UserUseCase{
		DB:                db,
		Log:               log,
		UserRepository:    ur,
		AddressRepository: ar,
	}
}

func (u *UserUseCase) Create(request *model.RequestUserRegister) (*model.ResponseUser, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	user := new(entity.User)

	// check duplicate email
	err = u.UserRepository.FindBy(tx, "email", request.Email, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	if user.ID != "" {
		fmt.Println("wowo1")
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("wowo4")
		return nil, err
	}

	now := time.Now().Local().Unix()

	user.ID = uuid.NewString()
	user.Name = request.FullName
	user.Email = request.Email
	user.Password = string(password)
	user.UrlPhoto = "default.jpg"
	user.CreatedAt = now
	user.UpdatedAt = now

	err = u.UserRepository.Create(tx, user)
	if err != nil {
		fmt.Println("create: " + err.Error())
		return nil, err
	}

	// add primary address
	address := &entity.Address{
		ID:          uuid.NewString(),
		Name:        "Main Address",
		Address:     request.Address,
		District:    request.District,
		City:        request.City,
		State:       request.State,
		PostalCode:  request.PostalCode,
		Latitude:    request.Latitude,
		Longitude:   request.Longitude,
		IsPrimary:   true,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   now,
		UpdatedAt:   now,
		UserID:      user.ID,
	}

	err = u.AddressRepository.Create(tx, address)
	if err != nil {
		return nil, err
	}

	token, err := firebase.CreateCustomToken(user.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseUser{Token: token}, nil
}

func (u *UserUseCase) Login(request *model.RequestUserLogin) (*model.ResponseUser, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	user := new(entity.User)

	// check user exists
	err = u.UserRepository.FindBy(tx, "email", request.Email, user)
	if err != nil {
		return nil, err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return nil, fiber.ErrUnauthorized
	}

	token, err := firebase.CreateCustomToken(user.ID)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseUser{Token: token}, nil
}

func (u *UserUseCase) Verify(request *model.UUIDMiddleware) error {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	user := new(entity.User)
	err = u.UserRepository.FindBy(tx, "id", request.UUID, user)
	if err != nil {
		return err
	}

	return nil

}
