package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()

	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	defer jsondb.CloseDB()

	c := controller.NewController(&cfg, jsondb)

	// id, err := c.CreateUser(
	// 	&models.CreateUser{
	// 		Name:    "Abduqodir",
	// 		Surname: "Musayev",
	// 	},
	// )

	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	user, err := c.GetPkeyUser(&models.UserPrimaryKey{Id: "c985ae75-e6ff-4dd5-82a7-1fe5d6076890"})
	if err != nil {
		log.Println("error while GetPkeyUser:", err.Error())
		return
	}

	fmt.Println(*user)
}
