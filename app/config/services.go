package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/services"
)

var (
	Ms = services.NewMoviesService()
)

func Services() raptor.Services {
	return raptor.Services{
		Ms,
	}
}
