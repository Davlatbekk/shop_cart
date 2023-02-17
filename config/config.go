package config

type Config struct {
	Path string

	UserFileName string

	ProductFileName string

	ShopcartFileName string
}

func Load() Config {

	cfg := Config{}

	cfg.Path = "./data"
	cfg.UserFileName = "/customer.json"
	cfg.ProductFileName = "/product.json"
	cfg.ShopcartFileName = "/shop_cart.json"

	return cfg
}
