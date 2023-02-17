package controller

import (
	"app/models"
)

func (c *Controller) CreateProduct(req *models.CreateProduct) (id string, err error) {

	id, err = c.store.Product().Create(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct) error {
	err := c.store.Product().UpdateProduct(req)
	if err != nil {
		return err
	}
	return nil
}
