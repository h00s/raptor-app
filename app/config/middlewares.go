package config

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/middlewares"
)

func Middlewares() raptor.Middlewares {
	return raptor.Middlewares{
		&middlewares.ActionLogger{},
	}
}
