package usecase

import (
	"github.com/ahmdyaasiin/magotify-backend/internal/app/entity"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/model"
	"github.com/ahmdyaasiin/magotify-backend/internal/app/repository"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type InterfaceMenuUseCase interface {
	//
	Explore(auth string) (*model.ResponseMenuExplore, error)
	Shop(auth string) (*model.ResponseMenuShop, error)
}

type MenuUseCase struct {
	//
	DB                 *sqlx.DB
	Log                *logrus.Logger
	MenuRepository     repository.InterfaceMenuRepository
	UserRepository     repository.InterfaceUserRepository
	AddressRepository  repository.InterfaceAddressRepository
	VoucherRepository  repository.InterfaceVoucherRepository
	CartRepository     repository.InterfaceCartRepository
	BannerRepository   repository.InterfaceBannerRepository
	CategoryRepository repository.InterfaceCategoryRepository
}

func NewMenuUseCase(db *sqlx.DB, log *logrus.Logger, mr repository.InterfaceMenuRepository, ur repository.InterfaceUserRepository, ar repository.InterfaceAddressRepository, vr repository.InterfaceVoucherRepository, cr repository.InterfaceCartRepository, br repository.InterfaceBannerRepository, car repository.InterfaceCategoryRepository) InterfaceMenuUseCase {
	//
	return &MenuUseCase{
		DB:                 db,
		Log:                log,
		MenuRepository:     mr,
		UserRepository:     ur,
		AddressRepository:  ar,
		VoucherRepository:  vr,
		CartRepository:     cr,
		BannerRepository:   br,
		CategoryRepository: car,
	}
}

func (u *MenuUseCase) Explore(auth string) (*model.ResponseMenuExplore, error) {
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

	var totalVoucher int

	err = u.VoucherRepository.TotalVouchers(tx, &totalVoucher, user)
	if err != nil {
		return nil, err
	}

	// get hot items
	var h []model.HotItemsSlice

	err = u.CategoryRepository.GetALlCategoriesName(tx, &h)
	if err != nil {
		return nil, err
	}

	allCategory := model.HotItemsSlice{
		Name:     "All",
		UrlPhoto: "",
		Products: nil,
	}

	h = append([]model.HotItemsSlice{allCategory}, h...)

	for i, n := range h {
		if n.Name == "All" {
			err = u.MenuRepository.HotItemsGeneral(tx, &h[i].Products, user.ID)
		} else {
			err = u.MenuRepository.HotItemsSpecific(tx, &h[i].Products, h[i].Name, user.ID)
		}

		if err != nil {
			return nil, err
		}
	}

	// get time to clean up
	var ttc []model.ExploreTimeToCleanUp

	err = u.MenuRepository.TimeToCleanUp(tx, &ttc, user.ID)
	if err != nil {
		return nil, err
	}

	var threeVehicles []model.ExploreTimeToCleanUp
	if len(ttc) >= 3 {
		threeVehicles = ttc[0:3]
	} else {
		threeVehicles = ttc
	}

	return &model.ResponseMenuExplore{
		User: model.ExploreUser{
			ID:           user.ID,
			FullName:     user.Name,
			PhotoProfile: user.UrlPhoto,
			Balance:      user.Balance,
			Voucher:      totalVoucher,
		},
		HotItems:      h,
		TimeToCleanUp: threeVehicles,
	}, nil
}

func (u *MenuUseCase) Shop(auth string) (*model.ResponseMenuShop, error) {
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

	address := new(entity.Address)

	err = u.AddressRepository.GetPrimaryAddress(tx, address, user)
	if err != nil {
		return nil, err
	}

	// cart
	var totalCart int

	err = u.CartRepository.CountCart(tx, &totalCart, user)
	if err != nil {
		return nil, err
	}

	// banner
	var banners []entity.Banner

	err = u.BannerRepository.GetAllBanner(tx, &banners)
	if err != nil {
		return nil, err
	}

	// get hot items
	var h []model.HotItemsSlice

	err = u.CategoryRepository.GetALlCategoriesName(tx, &h)
	if err != nil {
		return nil, err
	}

	allCategory := model.HotItemsSlice{
		Name:     "All",
		UrlPhoto: "",
		Products: nil,
	}

	h = append([]model.HotItemsSlice{allCategory}, h...)

	for i, n := range h {
		if n.Name == "All" {
			err = u.MenuRepository.HotItemsGeneral(tx, &h[i].Products, user.ID)
		} else {
			err = u.MenuRepository.HotItemsSpecific(tx, &h[i].Products, h[i].Name, user.ID)
		}

		if err != nil {
			return nil, err
		}
	}

	// products discount
	var productDiscount []model.ExploreItems

	err = u.MenuRepository.ProductBestOffer(tx, &productDiscount)
	if err != nil {
		return nil, err
	}

	return &model.ResponseMenuShop{
		User: model.ShopUser{
			ID:    user.ID,
			City:  address.City,
			State: address.State,
		},
		TotalCart:       totalCart,
		Banner:          banners,
		HotItems:        h,
		ProductDiscount: productDiscount,
	}, nil
}
