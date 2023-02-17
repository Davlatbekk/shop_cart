package jsondb

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
	file     *os.File
}

// Constructor
func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) Create(req *models.CreateUser) (id string, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return "", err
	}
	id = uuid.NewString()

	users = append(users, &models.User{
		Id:      id,
		Name:    req.Name,
		Surname: req.Surname,
		Balance: req.Balance,
	})

	body, err := json.MarshalIndent(users, "", "   ")

	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Get list of Users
func (u *userRepo) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {
	users := make([]models.User, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, err
	}

	if req.Limit+req.Offset > len(users) {
		if req.Offset > len(users) {
			return &models.GetListResponse{
				Count: len(users),
				Users: []models.User{},
			}, nil
		}

		return &models.GetListResponse{
			Count: len(users),
			Users: users[req.Offset:],
		}, nil
	}

	response := &models.GetListResponse{
		Count: len(users),
		Users: users[req.Offset : req.Limit+req.Offset],
	}

	return response, nil
}

// Get list by id
func (u *userRepo) GetUserById(req *models.UserPrimaryKey) (models.User, error) {
	users := make([]models.User, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.User{}, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return models.User{}, err
	}

	for _, v := range users {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

// Update user by id
func (u *userRepo) UpdateUser(req *models.UpdateUser) (models.User, error) {
	users := make([]models.User, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.User{}, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return models.User{}, err
	}

	updatedUser := models.User{}
	for i, v := range users {
		if v.Id == req.Id {
			users[i].Name = req.Name
			users[i].Surname = req.Surname
			users[i].Balance = req.Balance
			updatedUser = users[i]
		}
	}

	if len(updatedUser.Name) <= 0 {
		return models.User{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(users, "", "   ")

	if err != nil {
		return models.User{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

// Delete user by id
func (u *userRepo) DeleteUser(req *models.UserPrimaryKey) (models.User, error) {
	users := make([]models.User, 0)

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return models.User{}, err
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		return models.User{}, err
	}

	deletedUser := models.User{}
	for i, v := range users {
		if v.Id == req.Id {
			deletedUser = users[i]
			users = append(users[:i], users[i+1:]...)
		}
	}

	if len(deletedUser.Name) <= 0 {
		return models.User{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(users, "", "   ")

	if err != nil {
		return models.User{}, err
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return models.User{}, err
	}

	return deletedUser, nil
}
