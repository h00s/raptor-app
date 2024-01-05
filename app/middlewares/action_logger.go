package middlewares

import "github.com/h00s/raptor"

type ActionLogger struct {
	raptor.Middleware
}

func (m *ActionLogger) New(c *raptor.Context) error {
	// logging action
	return c.Next()
}
