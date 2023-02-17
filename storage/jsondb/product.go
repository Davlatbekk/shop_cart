package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type productRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewProductRepo(fileName string, file *os.File) *productRepo {
	return &productRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *productRepo) Create(req *models.CreateProduct) (string, error) {
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return "", err
	}
	var product []*models.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()
	product = append(product, &models.Product{
		Id:    id,
		Name:  req.Name,
		Price: req.Price,
	})

	body, err := json.MarshalIndent(product, "", "   ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *productRepo) UpdateProduct(req *models.UpdateProduct) error {
	var products []*models.Product
	err := json.NewDecoder(u.file).Decode(&products)
	if err != nil {
		return err
	}

	flag := true

	for i, val := range products {

		if val.Id == req.Id {

			products[i].Name = req.Name
			products[i].Price = req.Price

			flag = false
		}

	}
	if flag {
		return errors.New("BUnday uzgaruvchi yuq")
	}

	fmt.Println("O'zgartirildi")
	body, err := json.MarshalIndent(products, "", "   ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/product.json", body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}

func (p *productRepo) Delete(req *models.ProductPrimaryKey) error {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return err
	}
	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return err
	}

	flag := true
	for i, val := range products {
		if val.Id == req.Id {
			products = append(products[:i], products[i+1:]...)
			flag = false
		}
	}
	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(products, "", " ")
	err = ioutil.WriteFile(p.fileName, body, os.ModePerm)

	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) GetPkey(req *models.ProductPrimaryKey) (models.Product, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return models.Product{}, err
	}
	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	for _, val := range products {
		if val.Id == req.Id {
			return val, nil
		}
	}
	return models.Product{}, errors.New("There is no product with this id")
}

func (p *productRepo) GetList(req *models.GetListRequestProduct) (models.GetListResponseProduct, error) {
	data, err := ioutil.ReadFile(p.fileName)
	if err != nil {
		return models.GetListResponseProduct{}, err
	}
	var products []models.Product
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.GetListResponseProduct{}, err
	}

	return models.GetListResponseProduct{
		Products: products,
		Count:    len(products),
	}, nil
}
