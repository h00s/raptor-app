package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/controllers"
)

func Controllers() raptor.Controllers {
	hc := &controllers.HomeController{}

	return raptor.NewControllers(
		raptor.NewController("Home", &hc.Controller,
			raptor.Action("Root", hc.Root),
			raptor.Action("Index", hc.Index),
			raptor.Action("Example", hc.Example),
		),
	)
}
