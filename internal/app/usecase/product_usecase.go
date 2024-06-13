package usecase

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type InterfaceProductUseCase interface {
	//
	GetDetails(auth string, productID string) (*model.ProductDetails, error)
}

type ProductUseCase struct {
	//
	DB                *sqlx.DB
	Log               *logrus.Logger
	ProductRepository repository.InterfaceProductRepository
	UserRepository    repository.InterfaceUserRepository
	CartRepository    repository.InterfaceCartRepository
	MediaRepository   repository.InterfaceMediaRepository
	RatingRepository  repository.InterfaceRatingRepository
}

func NewProductUseCase(db *sqlx.DB, log *logrus.Logger, pr repository.InterfaceProductRepository, ur repository.InterfaceUserRepository, cr repository.InterfaceCartRepository, mr repository.InterfaceMediaRepository, rr repository.InterfaceRatingRepository) InterfaceProductUseCase {
	//
	return &ProductUseCase{
		DB:                db,
		Log:               log,
		ProductRepository: pr,
		UserRepository:    ur,
		CartRepository:    cr,
		MediaRepository:   mr,
		RatingRepository:  rr,
	}
}

func (u *ProductUseCase) GetDetails(auth string, productID string) (*model.ProductDetails, error) {
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

	// total cart
	var totalCart int

	err = u.CartRepository.CountCart(tx, &totalCart, user)
	if err != nil {
		return nil, err
	}

	// PD
	var productDetails model.PD

	err = u.ProductRepository.ProductDetails(tx, &productDetails, user, productID)
	if err != nil {
		return nil, err
	}

	// media
	var media []model.MediaProduct

	err = u.MediaRepository.GetAllMediaProduct(tx, &media, productID)
	if err != nil {
		return nil, err
	}

	productDetails.Media = media

	// review
	var reviews []model.ReviewProduct

	err = u.RatingRepository.GetAllRatings(tx, &reviews, productID)
	if err != nil {
		return nil, err
	}

	productDetails.Review = ""  // please fix this
	productDetails.Discuss = "" // please fix this

	// product discounts
	var productsDiscount []model.ExploreItems

	err = u.ProductRepository.ProductBestOfferWithout(tx, &productsDiscount, productID)
	if err != nil {
		return nil, err
	}

	return &model.ProductDetails{
		PD:              productDetails,
		TotalCart:       totalCart,
		ProductDiscount: productsDiscount,
	}, nil
}
