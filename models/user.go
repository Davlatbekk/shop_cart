package models

type UserPrimaryKey struct {
	Id string `json:"id"`
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type UpdateUser struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type CreateUser struct {
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Balance float64 `json:"balance"`
}

type GetListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListResponse struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}
