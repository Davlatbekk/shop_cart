package config

type Config struct {
	Path string

	UserFileName     string
	ProductFileName  string
	ShopCartFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/customer.json"
	cfg.ProductFileName = "/products.json"
	cfg.ShopCartFileName = "/shopCart.json"

	return cfg
}
