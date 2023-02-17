package controller

import (
	"app/models"
	"app/pkg/utils"
	"errors"
)

func (c *Controller) CreateProduct(req *models.CreateProduct) (id string, err error) {

	id, err = c.store.Product().CreateProduct(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetListProducts(req *models.GetListProductRequest) (*models.GetListProductResponse, error) {

	products, err := c.store.Product().GetListProduct(req)
	if err != nil {
		return &models.GetListProductResponse{}, err
	}

	return products, nil
}

func (c *Controller) GetProductByIdController(req *models.ProductPrimaryKey) (models.Product, error) {

	if !utils.IsValidUUID(req.Id) {
		return models.Product{}, errors.New("invalid uuid id")
	}

	product, err := c.store.Product().GetProductById(req)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}

func (c *Controller) UpdateProductController(req *models.UpdateProduct) (models.Product, error) {
	product, err := c.store.Product().UpdateProduct(req)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}
func (c *Controller) DeleteProductController(req *models.ProductPrimaryKey) (models.Product, error) {
	product, err := c.store.Product().DeleteProduct(req)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil

}
