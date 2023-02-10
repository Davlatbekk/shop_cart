package main

import (
	"fmt"

	"app/controller"
	"app/models"
)

func main() {

	controller.GenerateUser(455)

	users, err := controller.GetListUser(models.GetListRequest{
		Offset: 40,
		Limit: 10,
	})

	if err {
		fmt.Println("users out of range")
		return
	}

	for _, user := range users {
		fmt.Println(user)
	}
}

// Page Soni 46:
// Page yozing:
// 4

