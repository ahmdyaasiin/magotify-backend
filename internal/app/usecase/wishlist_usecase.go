package usecase

import (
	"fmt"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type InterfaceWishlistUseCase interface {
	//
	AddWishlist(auth string, request *model.RequestManageWishlist) (*model.ResponseManageWishList, error)
	GetWishlist(auth string) (*model.MyWishlist, error)
}

type WishlistUseCase struct {
	//
	DB                 *sqlx.DB
	Log                *logrus.Logger
	UserRepository     repository.InterfaceUserRepository
	WishlistRepository repository.InterfaceWishlistRepository
	CartRepository     repository.InterfaceCartRepository
	ProductRepository  repository.InterfaceProductRepository
}

func NewWishlistUseCase(db *sqlx.DB, log *logrus.Logger, ur repository.InterfaceUserRepository, wr repository.InterfaceWishlistRepository, cr repository.InterfaceCartRepository, pr repository.InterfaceProductRepository) InterfaceWishlistUseCase {
	//
	return &WishlistUseCase{
		DB:                 db,
		Log:                log,
		UserRepository:     ur,
		WishlistRepository: wr,
		CartRepository:     cr,
		ProductRepository:  pr,
	}
}

func (u *WishlistUseCase) AddWishlist(auth string, request *model.RequestManageWishlist) (*model.ResponseManageWishList, error) {
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

	// check wishlist
	wishlist := new(entity.Wishlist)

	err = u.WishlistRepository.FindBy(tx, "product_id", request.ProductID, wishlist, user)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, err
	}

	fmt.Println(wishlist)

	var message string
	if wishlist.ID == "" { // add
		message = "success add product"
		err = u.WishlistRepository.Create(tx, &entity.Wishlist{
			ID:        uuid.NewString(),
			CreatedAt: time.Now().Local().Unix(),
			UserID:    user.ID,
			ProductID: request.ProductID,
		})
		if err != nil {
			return nil, err
		}
	} else { // delete
		message = "success delete product"
		err = u.WishlistRepository.Delete(tx, wishlist)
		if err != nil {
			return nil, err
		}
	}

	// commit
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &model.ResponseManageWishList{
		Message: message,
	}, nil
}

func (u *WishlistUseCase) GetWishlist(auth string) (*model.MyWishlist, error) {
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

	// cart
	var totalCart int

	err = u.CartRepository.CountCart(tx, &totalCart, user)
	if err != nil {
		return nil, err
	}

	// my wishlist
	var myWishlistProducts []model.ExploreItems

	err = u.WishlistRepository.MyWishlist(tx, &myWishlistProducts, user)
	if err != nil {
		return nil, err
	}

	return &model.MyWishlist{
		TotalCart: totalCart,
		Product:   myWishlistProducts,
	}, nil
}
