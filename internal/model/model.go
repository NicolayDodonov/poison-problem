package model

import (
	"context"
	"math/rand/v2"
	"poison-problem/internal/logger"
)

type Model struct {
	Agents []*Agent
	*World
	*Statistic
	*Parameters
}

type Parameters struct {
	chanceFoodAppearing int
	levelPoisonHalfFood int
}

func New(countAgent, worldX, worldY int, sings ...*Sings) *Model {
	//init []*Agents
	agents := make([]*Agent, countAgent*len(sings))
	for s := 0; s < len(sings); s++ {
		for a := 0; a < countAgent; a++ {
			agents[s*a+a] = NewAgent(worldX, worldY, sings[s])
		}
	}

	//init World
	world := NewWorld(worldX, worldY)

	return &Model{
		agents,
		world,
		&Statistic{
			0,
			0,
			0,
			Sings{
				0, 0, 0, [2]int{0, 0}, 0, 0, 0,
			},
		},
		&Parameters{
			25,
			50,
		},
	}
}

// Run starts one epoch of agent life simulation.
// The epoch ends when the number of live agents is <= targetCountAgent.
func (m *Model) Run(ctx context.Context, logger *logger.Logger, targetCountAgent int) {
	for {
		//update world resources
		m.resourceHandler()

		liveAgent := 0
		//run all agent
		for _, agent := range m.Agents {
			err := agent.Run(m.World)
			if err != nil {
				logger.Error(err.Error())
			}
			if agent.Energy > 0 {
				liveAgent++
			}
		}

		if liveAgent <= targetCountAgent {
			break
		}
	}
	//update model stat
	m.statisticHandler()
}

func (m *Model) resourceHandler() {
	for _, cells := range m.Map {
		for _, cell := range cells {
			//if in cell much food - we half this food
			if cell.PoisonLevel > m.levelPoisonHalfFood && cell.FoodLevel > 15 {
				cell.FoodLevel /= 2
				continue
			}
			//if cell have small food - add it with a certain chance
			if cell.FoodLevel < 10 && rand.IntN(100) >= m.chanceFoodAppearing {
				cell.FoodLevel += 10
			}
		}
	}
}

// statisticHandler update information about
func (m *Model) statisticHandler() {
	food := 0
	poison := 0
	for _, cells := range m.Map {
		for _, cell := range cells {
			food += cell.FoodLevel
			poison += cell.PoisonLevel
		}
	}
	m.Statistic.Food = food
	m.Statistic.Poison = poison

	m.Statistic.Sings = Sings{
		0,
		0,
		0,
		[2]int{0, 0},
		0,
		0,
		0,
	}
	//sum sings parameters from agents
	for _, agent := range m.Agents {
		m.moveOrAction += agent.moveOrAction
		m.turnOrMove += agent.turnOrMove
		m.leftOrRight += agent.leftOrRight
		m.eatOrClear[0] += agent.eatOrClear[0]
		m.eatOrClear[1] += agent.eatOrClear[1]
		m.getFood += agent.getFood
		m.getPoison += agent.getPoison
		m.makePoison += agent.makePoison
	}
	//and make avg count of all sings parameters
	m.moveOrAction /= len(m.Agents)
	m.turnOrMove /= len(m.Agents)
	m.leftOrRight /= len(m.Agents)
	m.eatOrClear[0] /= len(m.Agents)
	m.eatOrClear[1] /= len(m.Agents)
	m.getFood /= len(m.Agents)
	m.getPoison /= len(m.Agents)
	m.makePoison /= len(m.Agents)
}
