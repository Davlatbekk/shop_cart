package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type shopCartRepo struct {
	fileName string
	file     *os.File
}

func NewShopCartRepo(fileName string, file *os.File) *shopCartRepo {
	return &shopCartRepo{
		fileName: fileName,
		file:     file,
	}
}

func (s *shopCartRepo) AddShopCart(req *models.Add) error {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return err
	}
	var shopCart []models.ShopCart
	err = json.Unmarshal(data, &shopCart)
	if err != nil {
		return err
	}

	shopCart = append(shopCart, models.ShopCart{
		UserId:    req.UserId,
		ProductId: req.ProductId,
		Count:     req.Count,
	})

	body, err := json.MarshalIndent(shopCart, "", " ")
	err = ioutil.WriteFile(s.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (s *shopCartRepo) RemoveShopCart(req *models.Remove) error {
	data, err := ioutil.ReadFile(s.fileName)
	if err != nil {
		return err
	}
	var shopCart []models.ShopCart
	err = json.Unmarshal(data, &shopCart)
	if err != nil {
		return err
	}

	for in, val := range shopCart {
		if val.UserId == req.UserId && val.ProductId == req.ProductId {
			shopCart = append(shopCart[:in], shopCart[in+1:]...)
		}
	}
	return nil
}
