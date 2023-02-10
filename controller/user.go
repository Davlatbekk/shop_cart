package controller

import (
	"app/models"
)

var Users []models.User

func CreateUser(data models.User) {
	Users = append(Users, data)
}

func GetListUser() []models.User {
	return Users
}

// getbyid
// update
// delete
