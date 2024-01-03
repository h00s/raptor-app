package controllers

import (
	"net/http"

	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/services"
)

type MoviesController struct {
	raptor.Controller

	Ms *services.MoviesService
}

func (mc *MoviesController) Index(c *raptor.Context) error {
	return c.Render("movies/index", raptor.Map{
		"Movies": mc.Ms.All(),
	})
}

func (mc *MoviesController) Show(c *raptor.Context) error {
	id, err := c.ParamsInt("id")
	if err == nil {
		if movie := mc.Ms.Find(id); movie != nil {
			return c.Render("movies/show", raptor.Map{
				"Movie": movie,
			})
		}
	}

	return c.SendStatus(http.StatusNotFound)
}
