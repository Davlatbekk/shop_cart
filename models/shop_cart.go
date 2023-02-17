package models

type ShopCart struct {
	UserId    string
	ProductId string
	Count     uint
}

type Add struct {
	UserId    string
	ProductId string
	Count     uint
}

type Remove struct {
	UserId    string
	ProductId string
}
