package controller

import (
	"app/models"

	"github.com/bxcodec/faker/v3"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser(req models.GetListRequest) (resp []models.User, err bool){

	if req.Limit + req.Offset > len(Users) {
		return []models.User{}, true
	}


	return Users[req.Offset : req.Limit + req.Offset], false
}

func GenerateUser(count int) {
	for i := 0; i < count; i++ {
		Users = append(Users, models.User{
			Id: i+1,
			Name: faker.FirstName(),
			Surname: faker.LastName(),
		})
	}
}
