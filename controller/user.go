package controller

import (
	"app/models"
	"errors"
)

func (c *Controller) CreateUser(req *models.CreateUser) (string, error) {
	id, err := c.store.User().Create(req)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (c *Controller) DeleteUser(req *models.UserPrimaryKey) error {
	err := c.store.User().Delete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) UpdateUser(req *models.UpdateUser, userId string) error {
	err := c.store.User().Update(req, userId)
	if err != nil {
		return err
	}
	return nil
} 

func (c *Controller) GetByIdUser(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetByID(req)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (c *Controller) GetAllUser(req *models.GetListRequest) (models.GetListResponse, error) {
	users, err := c.store.User().GetAll(req)
	if err != nil {
		return models.GetListResponse{}, err
	}
	return users, nil
}

func (c *Controller) WithdrawCheque(total float64, userId string) error {
	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})
	if err != nil {
		return err
	}
	if user.Balance >= total {
		user.Balance -= total
	} else {
		return errors.New("You don't have enough money")
	}

	err = c.store.User().Update(&models.UpdateUser{
		Balance: user.Balance,
	}, userId)
	if err != nil {
		return err
	}

	err = c.store.ShopCart().UpdateShopCart(userId)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) MoneyTransfer(sender string, receiver string, money float64) error {
	send, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: sender}) 
	if err != nil {
		return err
	}

	receive, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: receiver}) 
	if err != nil {
		return err
	}

	comMoney := 0.1 * float64(money)
	if money+comMoney > send.Balance {
		return errors.New("Sender doesn't have enough money")
	}
	send.Balance -= money + comMoney
	err = c.store.User().Update(&models.UpdateUser{
		Name: send.Name,
		Surname: send.Surname,
		Balance: send.Balance,
	}, sender)
	if err != nil {
		return err
	}

	err = c.store.Commission().AddCommission(&models.Commission{
		Balance: comMoney,
	})

	receive.Balance += money
	err = c.store.User().Update(&models.UpdateUser{
		Name: receive.Name,
		Surname: receive.Surname,
		Balance: receive.Balance,
	}, receiver)
	if err != nil {
		return err
	}
	return nil
}