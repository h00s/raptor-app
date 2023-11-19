package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/controllers"
)

func Router() *raptor.Router {
	hc := controllers.HomeController{}

	r := raptor.NewRouter()

	r.AddControllerRoute("Home", &hc,
		raptor.NewRoute("GET", "/", hc.Root),
		raptor.NewRoute("GET", "/Home", hc.Index),
		raptor.NewRoute("GET", "/Home/Example", hc.Example),
	)

	return r
}
