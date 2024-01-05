package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/controllers"
)

func Controllers() raptor.Controllers {
	return raptor.RegisterControllers(
		&controllers.HomeController{},
		&controllers.MoviesController{
			Ms: Ms,
		},
	)
}
