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
	"time"
)

type InterfaceUserUseCase interface {
	//
	Create(request *model.RequestUserRegister) (*model.ResponseUser, error)
	Login(request *model.RequestUserLogin) (*model.ResponseUser, error)
}

type UserUseCase struct {
	//
	DB             *sqlx.DB
	Log            *logrus.Logger
	UserRepository repository.InterfaceUserRepository
}

func NewUserUseCase(db *sqlx.DB, log *logrus.Logger, ur repository.InterfaceUserRepository) InterfaceUserUseCase {
	//
	return &UserUseCase{
		DB:             db,
		Log:            log,
		UserRepository: ur,
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
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if user.ID != "" {
		return nil, fiber.ErrConflict
	}

	// check duplicate phone_number
	err = u.UserRepository.FindBy(tx, "phone_number", request.Email, user)
	if err != nil {
		return nil, err
	}

	if user.ID != "" {
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now().Local().Unix()

	user.ID = uuid.NewString()
	user.Name = request.FullName
	user.Email = request.Email
	user.PhoneNumber = request.PhoneNumber
	user.Password = string(password)
	user.PhotoProfile = "default.jpg"
	user.CreatedAt = now
	user.UpdatedAt = now

	err = u.UserRepository.Create(tx, user)
	if err != nil {
		fmt.Println(err)
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
