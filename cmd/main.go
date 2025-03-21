package main

import (
	c "poison-problem/internal/config"
	l "poison-problem/internal/logger"
	s "poison-problem/internal/simulation"
)

func main() {

	Config := c.MustInit("config/config.yaml")

	Logger := l.New(
		Config.Logger.Path,
		Config.Logger.Type)

	Simulation := s.New(Logger, &Config.Simulation)

	Simulation.Run()
}
