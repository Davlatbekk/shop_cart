package controller

import "app/models"

func (c *Controller) CreateProduct(req *models.CreateProduct) (string, error) {
	id, err := c.store.Product().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) error {
	err := c.store.Product().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateProduct(req *models.UpdateProduct, productId string) error {
	err := c.store.Product().Update(req, productId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (models.Product, error) {
	product, err := c.store.Product().GetByID(req)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

func (c *Controller) GetAllProduct() (models.GetListProduct, error) {
	products, err := c.store.Product().GetAll()
	if err != nil {
		return models.GetListProduct{}, err
	}
	return products, nil
}
