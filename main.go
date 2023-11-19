package main

import (
	"github.com/h00s/raptor"
	"github.com/h00s/raptor-app/app/config"
)

func main() {
	r := raptor.NewRaptor(raptor.Config{
		Address: "localhost",
		Port:    3000,
	})
	r.SetRouter(config.Router())
	r.Start()
}
