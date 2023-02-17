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

	id, err := c.CreateUser(
		&models.CreateUser{
			Name:    "Abduqodir",
			Surname: "Musayev",
		},
	)
	if err != nil {
		log.Println("error while CreateUser:", err.Error())
		return
	}
	fmt.Println(id)

	id, err = c.CreateProduct(
		&models.CreateProduct{
			Name:  "olma",
			Price: "14000",
		},
	)

	if err != nil {
		log.Println("error while CreateProduct:", err.Error())
		return
	}
	fmt.Println(id)

	err = c.AddShopCart(&models.Add{
		UserId:    "c80f9fa8-e01c-4262-ab1d-e6557b248f3e",
		ProductId: "d070ecc9-3e17-4b14-9598-f1296a38004e",
		Count:     3,
	})

	if err != nil {
		log.Println("error while AddShopCart:", err.Error())
		return
	}

	/////////////////////UpdateUser///////////////////////////
	// err = c.UpdateProduct(&models.UpdateProduct{
	// 	Id:    "b4e14d79-483c-47f3-b10e-5721f7b85160",
	// 	Name:  "Davlat",
	// 	Price: "Jalolov",
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// if err != nil {
	// 	log.Println("error while CreateUser:", err.Error())
	// 	return
	// }

	// user, err := c.GetPkeyUser(&models.UserPrimaryKey{Id: "7"})
	// if err != nil {
	// 	log.Println("error while GetPkeyUser:", err.Error())
	// 	return
	// }

	// fmt.Println(*user)

	// // 	//////////////////////////GITListUser////////////////////////////////
	// users, e := c.GitListUser(&models.GetListRequest{
	// 	Limit:  1,
	// 	Offset: 1,
	// })

	// if e != nil {
	// 	log.Println(e)
	// }
	// fmt.Println(users)

}
