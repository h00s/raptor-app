package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/controllers"
	"github.com/h00s/raptor-app/app/services"
)

func Controllers() raptor.Controllers {
	ms := services.NewMoviesService()

	return raptor.RegisterControllers(
		&controllers.HomeController{},
		&controllers.MoviesController{
			Ms: ms,
		},
	)
}
