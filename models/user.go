package models

type User struct {
	Id int
	Name string
	Surname string
}

type GetListRequest struct {
	Offset int
	Limit int
}