package initializers

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/controllers"
	"github.com/h00s/raptor-app/app/middlewares"
	"github.com/h00s/raptor-app/app/services"
)

func App() *raptor.AppInitializer {
	ms := services.NewMoviesService()

	return &raptor.AppInitializer{
		Services: raptor.Services{
			ms,
		},

		Middlewares: raptor.Middlewares{
			&middlewares.ActionLogger{},
		},

		Controllers: raptor.Controllers{
			&controllers.HomeController{},
			&controllers.MoviesController{
				Ms: ms,
			},
		},
	}
}
