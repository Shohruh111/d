package controller

import (
	"app/config"
	"app/models"
)

type Controller struct {
	Cfg   *config.Config
	Users []*models.User
}

func NewController(cfg *config.Config) *Controller {
	return &Controller{
		Cfg: cfg,
	}
}
