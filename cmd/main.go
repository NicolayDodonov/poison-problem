package main

import (
	c "poison-problem/internal/config"
	l "poison-problem/internal/logger"
	s "poison-problem/internal/simulation"
)

func main() {

	Config := c.MustInit("configs/config.yaml")

	Logger := l.New(
		Config.Logger.Path,
		Config.Logger.Type)

	Logger.Info("Application start")
	Simulation := s.New(Logger, &Config.Simulation)

	Simulation.Run()

	Logger.Info("Application shut down without errors")
}
