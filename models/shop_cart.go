package models

type ShopCartPrimaryKey struct {
	Id string `json:"id"`
}

type ShopCart struct {
	Id        string `json:"id"`
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
}

type Add struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
	Count     int    `json:"count"`
}

type Remove struct {
	ProductId string `json:"productId"`
	UserId    string `json:"userID"`
}
