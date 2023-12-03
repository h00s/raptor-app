package controllers

import (
	"github.com/h00s/raptor"
)

type HomeController struct {
	raptor.Controller
}

func (hc *HomeController) Root(c *raptor.Context) error {
	return c.SendString("Hello from Home#Root")
}

func (hc *HomeController) Index(c *raptor.Context) error {
	return c.SendString("Hello from Home#Index")
}

func (hc *HomeController) Example(c *raptor.Context) error {
	return c.SendString("Hello from Home#Example")
}
