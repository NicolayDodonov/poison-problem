package model

import (
	"context"
	"fmt"
	"poison-problem/internal/logger"
)

type Model struct {
	Agents []*Agent
	*World
	*Statistic
}

func New() *Model {
	//todo: init []*Agents

	//todo: init World

	//todo: init Statistic

	return &Model{}
}

// Run проводит исследование отдельной группы агентов на заданные условия мира в
// структуре модели. После завершения моделирования, выводит какие то результаты
// во вне. //todo: определить что будем возвращать
func (m *Model) Run(ctx context.Context, logger *logger.Logger) {
	for {
		//update world resources
		m.resourceHandler()

		//run all agent
		for _, agent := range m.Agents {
			err := agent.Run(m.World)
			if err != nil {
				logger.Error(fmt.Errorf("Agent ID", err).Error())
			}
		}

		//update model stat
		m.statHandler()

		//todo: ch <- some info

		//handle ctx event
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func (m *Model) resourceHandler() {
	//todo: make smart spawn resource system
}

func (m *Model) statHandler() {
	food := 0
	poison := 0
	for _, cells := range m.Map {
		for _, cell := range cells {
			food += cell.FoodLevel
			poison += cell.PoisonLevel
		}
	}
	m.Food = food
	m.Poison = poison
}
