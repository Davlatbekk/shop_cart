package controller

import (
	"app/config"
	"app/storage"
)

type Controller struct {
	cfg   *config.Config
	store storage.StorageI
}

// Controller Constructor
func NewController(cfg *config.Config, store storage.StorageI) *Controller {
	return &Controller{
		cfg:   cfg,
		store: store,
	}
}
