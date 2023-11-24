package controllers

import (
	"github.com/h00s/raptor"
)

type HomeController struct {
	raptor.Controller
}

func NewHomeController() *HomeController {
	hc := &HomeController{}
	hc.Name = "Home"
	hc.RegisterActions(
		raptor.Action("Root", hc.Root),
		raptor.Action("Index", hc.Index),
		raptor.Action("Example", hc.Example),
	)
	return hc
}

func (hc *HomeController) Root(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Root")
	return c.SendString("Hello from Home#Root")
}

func (hc *HomeController) Index(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Index")
	return c.SendString("Hello from Home#Index")
}

func (hc *HomeController) Example(c *raptor.Context) error {
	hc.Services.Log.Info("HomeController.Example")
	return c.SendString("Hello from Home#Example")
}
