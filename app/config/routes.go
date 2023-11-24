package config

import (
	"github.com/h00s/raptor"
)

func Routes() raptor.Routes {
	return raptor.Routes{
		raptor.Route("GET", "/", "Home", "Root"),
		raptor.Route("GET", "/Home", "Home", "Index"),
		raptor.Route("GET", "/Home/Example", "Home", "Example"),
	}
}
