package jsonDb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type commissionRepo struct {
	fileName string
}

func NewCommissionRepo(fileName string) *commissionRepo {
	return &commissionRepo{
		fileName: fileName,
	}
}

func (c *commissionRepo) AddCommission(req *models.Commission) error {
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		return err
	}

	var commissions []models.Commission
	err = json.Unmarshal(data, &commissions)
	if err != nil {
		return err
	}

	commissions = append(commissions, models.Commission{
		Balance: req.Balance,
	})

	body, err := json.MarshalIndent(commissions, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
