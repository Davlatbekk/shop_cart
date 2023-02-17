package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"fmt"
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

func (u *userRepo) Create(req *models.CreateUser) (string, error) {

	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return "", err
	}

	id := uuid.New().String()
	if len(users) > 0 {
		// id = users[len(users)-1].Id + 1
		users = append(users, &models.User{
			Id:      id,
			Name:    req.Name,
			Surname: req.Surname,
		})
	}

	body, err := json.MarshalIndent(users, "", "   ")

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) GetPkey(req *models.UserPrimaryKey) (res *models.User, err error) {

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return &models.User{}, err
	}
	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		return &models.User{}, err
	}

	for _, user := range users {
		if user.Id == req.Id {
			return &models.User{}, nil
		}
	}

	return nil, errors.New("no found user")
}

func (u *userRepo) GitList(req *models.GetListRequest) (*models.GetListResponse, error) {

	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return &models.GetListResponse{}, err
	}
	var users []models.User

	err = json.Unmarshal(data, &users)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	if req.Limit+req.Offset > len(users) {
		return &models.GetListResponse{}, err

	}
	return &models.GetListResponse{
		Count: req.Limit,
		Users: users[req.Offset : req.Limit+req.Offset],
	}, nil
}

func (u *userRepo) UpdateUser(req *models.UpdateUser) error {
	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return err
	}

	flag := false

	for i, val := range users {

		if val.Id == req.Id {

			users[i].Name = req.Name
			users[i].Surname = req.Surname

			flag = true
		}

	}
	if flag {
		return errors.New("BUnday uzgaruvchi yuq")
	}

	fmt.Println("O'zgartirildi")
	body, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/users.json", body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}

func (u *userRepo) Delete(req *models.UserPrimaryKey) error {

	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		return err
	}
	for i, val := range users {

		if req.Id == val.Id {

			users = append(users[:i], users[i+1:]...)
		}
	}
	fmt.Println("Ochirildi")

	body, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("data/users.json", body, os.ModePerm)
	if err != nil {
		return err
	}
	return nil

}

// package storage

// import (
// 	"app/models"
// 	"encoding/json"
// 	"errors"
// 	"io/ioutil"

// 	"os"
// )

// type userRepo struct {
// 	fileName string
// 	file     *os.File
// }

// // Constructor
// func NewUserRepo(fileName string, file *os.File) *userRepo {
// 	return &userRepo{
// 		fileName: fileName,
// 		file:     file,
// 	}
// }

// func (u *userRepo) Create(req *models.CreateUser) (id int, err error) {

// 	var users []*models.User
// 	err = json.NewDecoder(u.file).Decode(&users)
// 	if err != nil {
// 		return 0, err
// 	}

// 	if len(users) > 0 {
// 		id = users[len(users)-1].Id + 1
// 		users = append(users, &models.User{
// 			Id:      id,
// 			Name:    req.Name,
// 			Surname: req.Surname,
// 		})
// 	} else {
// 		id = 1
// 		users = append(users, &models.User{
// 			Id:      id,
// 			Name:    req.Name,
// 			Surname: req.Surname,
// 		})
// 	}

// 	body, err := json.MarshalIndent(users, "", "   ")
// 	if err != nil {
// 		return id, err
// 	}

// 	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

// func (u *userRepo) GitList(req *models.GetListRequest) (models.GetListResponse, error) {

// 	// data, err := ioutil.ReadFile(u.fileName)
// 	data, err := ioutil.ReadFile(u.fileName)
// 	if err != nil {
// 		return models.GetListResponse{}, err
// 	}
// 	var users []models.User

// 	err = json.Unmarshal(data, &users)
// 	if err != nil {
// 		return models.GetListResponse{}, err
// 	}

// 	if req.Limit+req.Offset > len(users) {
// 		return models.GetListResponse{}, err

// 	}
// 	return models.GetListResponse{
// 		Count: req.Limit,
// 		Users: users[req.Offset : req.Limit+req.Offset],
// 	}, nil
// }

// func (u *userRepo) UpdateUser(req *models.UpdateUser) error {
// 	var users []*models.User
// 	err := json.NewDecoder(u.file).Decode(&users)
// 	if err != nil {
// 		return err
// 	}

// 	flag := false

// 	for i, val := range users {

// 		if val.Id == req.Id {

// 			users[i].Name = req.Name
// 			users[i].Surname = req.Surname

// 			flag = true
// 		}

// 	}
// 	if flag {
// 		return errors.New("BUnday uzgaruvchi yuq")
// 	}

// 	// fmt.Println("O'zgartirildi")
// 	body, err := json.MarshalIndent(users, "", "   ")
// 	if err != nil {
// 		return err
// 	}

// 	err = ioutil.WriteFile("data/users.json", body, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
