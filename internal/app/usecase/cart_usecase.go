package usecase

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type InterfaceCartUseCase interface {
	//
	GetCart(auth string) (*model.MyCart, error)
	AddCart(auth string, request *model.RequestAddCart) (*model.ResponseAddCart, error)
}

type CartUseCase struct {
	//
	DB                *sqlx.DB
	Log               *logrus.Logger
	MenuRepository    repository.InterfaceMenuRepository
	UserRepository    repository.InterfaceUserRepository
	CartRepository    repository.InterfaceCartRepository
	ProductRepository repository.InterfaceProductRepository
}

func NewCartUseCase(db *sqlx.DB, log *logrus.Logger, mr repository.InterfaceMenuRepository, ur repository.InterfaceUserRepository, cr repository.InterfaceCartRepository, pr repository.InterfaceProductRepository) InterfaceCartUseCase {
	//
	return &CartUseCase{
		DB:                db,
		Log:               log,
		MenuRepository:    mr,
		UserRepository:    ur,
		CartRepository:    cr,
		ProductRepository: pr,
	}
}

func (u *CartUseCase) GetCart(auth string) (*model.MyCart, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	// product cart
	var productCart []model.ProductCart

	err = u.CartRepository.MyCart(tx, &productCart, user)
	if err != nil {
		return nil, err
	}

	// total cart
	var totalCart int

	err = u.CartRepository.CountCart(tx, &totalCart, user)
	if err != nil {
		return nil, err
	}

	// products discount
	var productDiscount []model.ExploreItems

	err = u.MenuRepository.ProductBestOffer(tx, &productDiscount)
	if err != nil {
		return nil, err
	}

	return &model.MyCart{
		Product:         productCart,
		TotalCart:       totalCart,
		ProductDiscount: productDiscount,
	}, nil
}

func (u *CartUseCase) AddCart(auth string, request *model.RequestAddCart) (*model.ResponseAddCart, error) {
	tx, err := u.DB.Beginx()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}

	// get data user
	user := new(entity.User)

	err = u.UserRepository.FindBy(tx, "id", auth, user)
	if err != nil {
		return nil, err
	}

	// check request product_id
	product := new(entity.Product)

	err = u.ProductRepository.FindBy(tx, "id", request.ProductID, product)
	if err != nil {
		return nil, err
	}

	// add cart
	cart := new(entity.Cart)

	err = u.CartRepository.FindBy(tx, "product_id", request.ProductID, cart, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	if request.Quantity == 0 {
		err = u.CartRepository.Delete(tx, cart)
		if err != nil {
			return nil, err
		}
	} else {
		if cart.ID == "" { // add
			err = u.CartRepository.Create(tx, &entity.Cart{
				ID:        uuid.NewString(),
				Quantity:  request.Quantity,
				CreatedAt: time.Now().Local().Unix(),
				UpdatedAt: time.Now().Local().Unix(),
				UserID:    user.ID,
				ProductID: request.ProductID,
			})
			if err != nil {
				return nil, err
			}
		} else { // update
			cart.Quantity = request.Quantity
			cart.UpdatedAt = time.Now().Local().Unix()
			err = u.CartRepository.Update(tx, cart)
			if err != nil {
				return nil, err
			}
		}
	}

	// after add cart
	var totalCart int

	err = u.CartRepository.CountCart(tx, &totalCart, user)
	if err != nil {
		return nil, err
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseAddCart{
		TotalCart: totalCart,
	}, nil
}
