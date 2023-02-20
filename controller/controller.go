package controller

import (
	"app/config"
	"app/storage"
)

type Controller struct {
	store storage.StorageI
	cfg *config.Config
}

func NewController(cfg *config.Config, store storage.StorageI) *Controller {
	return &Controller{
		store: store,
		cfg: cfg,
	}
}