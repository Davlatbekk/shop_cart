package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsonDb"
	"fmt"
	"log"
)

func main() {
	cfg := config.Load()

	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		log.Fatal("error while connecting to database")
	}
	defer jsonDb.CloseDb()

	c := controller.NewController(&cfg, jsonDb)

	// c.CreateProduct(&models.CreateProduct{
	// 	Name:       "Smartfon vivo V25 8/256 GB",
	// 	Price:      4_860_000,
	// 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
	// })

	product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

	if err != nil {
		log.Println(err)
		return
	}

	// c.GetAllProduct(
	// 	offset:
	// 	limit:
	// 	categoryid: "38292285-4c27-497b-bc5f-dfe418a9f959"
	// )

	fmt.Printf("%+v\n", product)

}

func Category(c *controller.Controller) {
	// c.CreateCategory(&models.CreateCategory{
	// 	Name:     "Smartfonlar va telefonlar",
	// 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
	// })

	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(category)

}

func User(c *controller.Controller) {

	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
	err := c.MoneyTransfer(sender, receiver, 500_000)
	if err != nil {
		log.Println(err)
	}
}
