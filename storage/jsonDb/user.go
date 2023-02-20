package jsonDb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type userRepo struct {
	fileName string
}

func NewUserRepo(fileName string) *userRepo {
	return &userRepo{
		fileName: fileName,
	}
}

func (u *userRepo) Read() ([]models.User, error) {
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return []models.User{}, err
	}

	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

func (u *userRepo) Create(req *models.CreateUser) (string, error) {
	users, err := u.Read()
	if err != nil {
		return "", err
	}

	uuid := uuid.New().String()
	users = append(users, models.User{
		Id: uuid,
		Name: req.Name,
		Surname: req.Surname,
		Balance: req.Balance,
	})

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (u *userRepo) Delete(req *models.UserPrimaryKey) error {
	users, err := u.Read()
	if err != nil {
		return err
	}
	flag := true
	for i, v := range users {
		if v.Id == req.Id {
			users = append(users[:i], users[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) Update(req *models.UpdateUser, userId string) error {
	users, err := u.Read()
	if err != nil {
		return err
	}

	flag := true
	for i, v := range users {
		if v.Id == userId {
			users[i].Name = req.Name
			users[i].Surname = req.Surname
			users[i].Balance = req.Balance
			flag = false
		}
	}

	if flag {
		return errors.New("There is no user with this id")
	}

	body, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err	
	}
	return nil
}

func (u *userRepo) GetByID(req *models.UserPrimaryKey) (models.User, error) {
	users, err := u.Read()
	if err != nil {
		return models.User{}, err
	}

	for _, v := range users {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.User{}, errors.New("There is no user with this id")
}

func (u *userRepo) GetAll(req *models.GetListRequest) (models.GetListResponse, error) {
	users, err := u.Read()
	if err != nil {
		return models.GetListResponse{}, err
	}

	if req.Limit + req.Offset > len(users) {
		return models.GetListResponse{}, errors.New("out of range")
	}
	
	fUsers := []models.User{}
	for i:=req.Offset; i<req.Offset+req.Limit; i++ {
		fUsers = append(fUsers, users[i])
	}
	return models.GetListResponse{
		Users: fUsers,
		Count: len(fUsers),
	}, nil
}