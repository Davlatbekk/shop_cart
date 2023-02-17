package models

type ProductPrimaryKey struct {
	Id string `json:"id"`
}

type Product struct {
	Id    string `json:"product_id"`
	Name  string `json:"product_name"`
	Price string `json:"price"`
}

type CreateProduct struct {
	Id    string `json:"product_id"`
	Name  string `json:"product_name"`
	Price string `json:"price"`
}

type UpdateProduct struct {
	Id    string `json:"id"`
	Name  string `json:"product_name"`
	Price string `json:"product_price"`
}

type GetListRequestProduct struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetListResponseProduct struct {
	Count    int       `json:"count"`
	Products []Product `json:"products"`
}
