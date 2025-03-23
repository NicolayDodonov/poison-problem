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

func New(countAgent, worldX, worldY int, sings ...*Sing) *Model {
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
			0,
			Sing{
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
		//update model stat
		m.statisticHandler(false)
		m.Statistic.Year++
	}
	m.statisticHandler(true)
}

func (m *Model) Reset() {
	//clear map from food and poison
	for _, cells := range m.Map {
		for _, cell := range cells {
			cell.FoodLevel = 0
			cell.PoisonLevel = 0
		}
	}
	//todo: make all agent to base count
	for _, agent := range m.Agents {
		agent.Cords = Cords{Y: rand.IntN(m.MaxY), X: rand.IntN(m.MaxX)}
		agent.Look = turn(rand.IntN(8))
		agent.Energy = 100
		agent.Age = 0

	}
}

func (m *Model) Fitness() {
	//todo: get sings best agent.age
	//todo: sort Agent to age
	//todo: delete old agent
	//todo: make new agent with best sings
	//todo: mutation sings
}

func (m *Model) CheckTargetAge(target int) bool {
	for _, agent := range m.Agents {
		if agent.Age >= target {
			return true
		}
	}
	return false
}

func (m *Model) BestSing() *Sing {
	//todo!
	return nil
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
func (m *Model) statisticHandler(updateGlobal bool) {
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

	sum := 0
	count := 0
	for _, agent := range m.Agents {
		if agent.Energy > 0 {
			sum += agent.Energy
			count++
		}

	}
	m.Statistic.AvgEnergy = sum / count

	if updateGlobal {
		m.Statistic.Sing = Sing{
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
			m.MoveOrAction += agent.MoveOrAction
			m.TurnOrMove += agent.TurnOrMove
			m.LeftOrRight += agent.LeftOrRight
			m.EatOrClear[0] += agent.EatOrClear[0]
			m.EatOrClear[1] += agent.EatOrClear[1]
			m.GetFood += agent.GetFood
			m.GetPoison += agent.GetPoison
			m.MakePoison += agent.MakePoison
		}
		//and make avg count of all sings parameters
		m.MoveOrAction /= len(m.Agents)
		m.TurnOrMove /= len(m.Agents)
		m.LeftOrRight /= len(m.Agents)
		m.EatOrClear[0] /= len(m.Agents)
		m.EatOrClear[1] /= len(m.Agents)
		m.GetFood /= len(m.Agents)
		m.GetPoison /= len(m.Agents)
		m.MakePoison /= len(m.Agents)
	}
}
