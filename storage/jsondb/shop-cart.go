package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type shopCartRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewShopCartRepo(fileName string, file *os.File) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (s *shopCartRepo) AddShopCart(req *models.AddShopCart) (models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return models.ShopCart{}, err
	}

	// if userId and productId exist replace only count
	newShopCart := models.ShopCart{}

	id := uuid.NewString()

	newShopCart = models.ShopCart{
		Id:        id,
		ProductId: req.ProductId,
		UserId:    req.UserId,
		Count:     req.Count,
	}
	carts = append(carts, newShopCart)

	// stringify struct to json
	body, err := json.MarshalIndent(carts, "", "   ")
	if err != nil {
		return models.ShopCart{}, err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return models.ShopCart{}, err
	}

	return newShopCart, nil

}

func (s *shopCartRepo) RemoveShopCart(req *models.RemoveShopCart) (models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return models.ShopCart{}, err
	}

	deletedShopCart := models.ShopCart{}
	for i, v := range carts {
		if v.UserId == req.UserId && v.ProductId == req.ProductId {
			deletedShopCart = carts[i]
			carts = append(carts[:i], carts[i+1:]...)
		}
	}

	if len(deletedShopCart.Id) <= 0 {
		return models.ShopCart{}, errors.New("shop-cart not found")
	}

	// stringify struct to json
	body, err := json.MarshalIndent(carts, "", "   ")
	if err != nil {
		return models.ShopCart{}, err
	}

	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return models.ShopCart{}, err
	}

	return deletedShopCart, nil
}

func (s *shopCartRepo) GetUserShopCarts(req *models.UserPrimaryKey) ([]models.ShopCart, error) {
	carts := []models.ShopCart{}

	// Read data from file
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return []models.ShopCart{}, err
	}

	// parse json data
	err = json.Unmarshal(data, &carts)
	if err != nil {
		return []models.ShopCart{}, err
	}

	userShopCarts := []models.ShopCart{}

	for _, v := range carts {
		if v.UserId == req.Id {
			userShopCarts = append(userShopCarts, v)
		}
	}

	return userShopCarts, nil
}
