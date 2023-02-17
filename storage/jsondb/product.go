package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
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

func (u *productRepo) CreateProduct(req *models.CreateProduct) (id string, err error) {

	var products []*models.Product
	err = json.NewDecoder(u.file).Decode(&products)
	if err != nil {
		return "", err
	}
	id = uuid.NewString()

	products = append(products, &models.Product{
		Id:    id,
		Name:  req.Name,
		Price: req.Price,
	})

	body, err := json.MarshalIndent(products, "", "   ")

	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Get list of Products
func (u *productRepo) GetListProduct(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return nil, err
	}

	if req.Limit+req.Offset > len(products) {
		if req.Offset > len(products) {
			return &models.GetListProductResponse{
				Count:    len(products),
				Products: []models.Product{},
			}, nil
		}

		return &models.GetListProductResponse{
			Count:    len(products),
			Products: products[req.Offset:],
		}, nil
	}

	response := &models.GetListProductResponse{
		Count:    len(products),
		Products: products[req.Offset : req.Limit+req.Offset],
	}

	return response, nil
}

// Get list by id
func (u *productRepo) GetProductById(req *models.ProductPrimaryKey) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	for _, v := range products {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.Product{}, errors.New("product not found")
}

// Update user by id
func (u *productRepo) UpdateProduct(req *models.UpdateProduct) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	updatedUser := models.Product{}
	for i, v := range products {
		if v.Id == req.Id {
			products[i].Name = req.Name
			products[i].Price = req.Price
			updatedUser = products[i]
		}
	}

	if len(updatedUser.Name) <= 0 {
		return models.Product{}, errors.New("product not found")
	}

	body, err := json.MarshalIndent(products, "", "   ")

	if err != nil {
		return models.Product{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.Product{}, err
	}

	return updatedUser, nil

}

// Delete user by id
func (u *productRepo) DeleteProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	products := make([]models.Product, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.Product{}, err
	}
	err = json.Unmarshal(data, &products)
	if err != nil {
		return models.Product{}, err
	}

	deletedUser := models.Product{}
	for i, v := range products {
		if v.Id == req.Id {
			deletedUser = products[i]
			products = append(products[:i], products[i+1:]...)
		}
	}

	if len(deletedUser.Name) <= 0 {
		return models.Product{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(products, "", "   ")

	if err != nil {
		return models.Product{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.Product{}, err
	}

	return deletedUser, nil

}
