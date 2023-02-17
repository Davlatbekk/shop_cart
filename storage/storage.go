package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetPkey(*models.UserPrimaryKey) (*models.User, error)
	GitList(*models.GetListRequest) (*models.GetListResponse, error)
	UpdateUser(*models.UpdateUser) error
	Delete(*models.UserPrimaryKey) error
}

type ProductRepoI interface {
	Create(*models.CreateProduct) (string, error)
	UpdateProduct(*models.UpdateProduct) error
	GetPkey(req *models.ProductPrimaryKey) (models.Product, error)
	GetList(req *models.GetListRequestProduct) (models.GetListResponseProduct, error)
	Delete(*models.ProductPrimaryKey) error
}

type ShopCartI interface {
	AddShopCart(*models.Add) error
	RemoveShopCart(*models.Remove) error
}
