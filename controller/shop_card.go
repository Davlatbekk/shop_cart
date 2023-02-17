package controller

import (
	"app/models"
)

func (c *Controller) AddShopCart(req *models.Add) error {
	_, err := c.store.User().GetPkey(&models.UserPrimaryKey{Id: req.UserId})
	if err != nil {
		return err
	}

	_, err = c.store.Product().GetPkey(&models.ProductPrimaryKey{Id: req.ProductId})
	if err != nil {
		return err
	}

	err = c.store.ShopCart().AddShopCart(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) RemoveShopCart(req *models.Remove) error {
	err := c.store.ShopCart().RemoveShopCart(req)
	if err != nil {
		return err
	}
	return nil
}
