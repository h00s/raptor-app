package main

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/config"
)

func main() {
	r := raptor.NewRaptorMVC(raptor.Config{
		Address: "localhost",
		Port:    3000,
	})

	r.Controllers(config.Controllers())
	r.Routes(config.Routes())

	r.Listen()
}
