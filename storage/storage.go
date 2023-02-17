package storage

import "app/models"

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Product() ProductRepoI
	ShopCart() ShopCartRepoI
}

type UserRepoI interface {
	Create(*models.CreateUser) (string, error)
	GetUserById(req *models.UserPrimaryKey) (models.User, error)
	GetList(req *models.GetListRequest) (*models.GetListResponse, error)
	UpdateUser(req *models.UpdateUser) (models.User, error)
	DeleteUser(req *models.UserPrimaryKey) (models.User, error)
}

type ProductRepoI interface {
	CreateProduct(req *models.CreateProduct) (id string, err error)
	GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error)
	GetProductById(req *models.ProductPrimaryKey) (models.Product, error)
	UpdateProduct(req *models.UpdateProduct) (models.Product, error)
	DeleteProduct(req *models.ProductPrimaryKey) (models.Product, error)
}

type ShopCartRepoI interface {
	AddShopCart(req *models.AddShopCart) (models.ShopCart, error)
	RemoveShopCart(req *models.RemoveShopCart) (models.ShopCart, error)
	GetUserShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error)
}
