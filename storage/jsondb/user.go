package jsondb

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"

	"app/models"
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

	id = uuid.New().String()

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return "", err
	}

	users = append(users, &models.User{
		Id:      id,
		Name:    req.Name,
		Surname: req.Surname,
	})

	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) GetPkey(req *models.UserPrimaryKey) (res *models.User, err error) {

	var users []*models.User
	err = json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Id == req.Id {
			return user, nil
		}
	}

	return nil, errors.New("No found user")
}
