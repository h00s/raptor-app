package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/", "HomeController", "Root"),
		raptor.Route("GET", "/Home", "HomeController", "Index"),
		raptor.Route("GET", "/Home/Example", "HomeController", "Example"),
		raptor.Route("GET", "/movies", "MoviesController", "Index"),
		raptor.Route("GET", "/movies/:id", "MoviesController", "Show"),
	}
}
