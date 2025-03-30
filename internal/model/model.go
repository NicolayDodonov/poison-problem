package model

import (
	"context"
	"math/rand/v2"
	"poison-problem/internal/config"
	"poison-problem/internal/logger"
	"strconv"
)

type Model struct {
	Agents []*Agent
	*World
	*Statistic
	*Parameters
}

type Parameters struct {
	chanceFoodAppearing int
	maxPoisonLevel      int
	CountSings          int
}

func New(countAgent int, conf *config.World, sings []*Sing) *Model {
	//make empty slice
	agents := make([]*Agent, 0)

	//fill agents
	for s, sing := range sings {
		group := make([]*Agent, countAgent)
		//fill group agent
		for a := 0; a < countAgent; a++ {
			group[a] = NewAgent(
				conf.MaxX,
				conf.MaxY,
				strconv.Itoa(s)+"-"+strconv.Itoa(a),
				sing,
			)
		}
		agents = append(agents, group...)
	}

	//init World
	world := NewWorld(conf)

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
			conf.ChanceFood,
			conf.MaxLevel,
			len(sings),
		},
	}
}

// Run starts one epoch of agent life simulation.
// The epoch ends when the number of live agents is <= targetCountAgent.
func (m *Model) Run(
	ctx context.Context,
	logger *logger.Logger,
	maxAge,
	targetCountAgent int,
	handlerStat func(*Model),
	handlerUpdate func(*Model),
) {
	logger.Info("Model started")
	for {
		logger.Debug("age №" + strconv.Itoa(m.Year))
		//update world resources
		m.resourceHandler()

		liveAgent := 0
		//run all agent
		for _, agent := range m.Agents {
			err := agent.Run(m.World)
			if err != nil {
				logger.Error("Agent №" + agent.ID + " have error " + err.Error())
			}
			if agent.Energy > 0 {
				logger.Debug("Agent №" + agent.ID + " is live!")
				liveAgent++
			}
		}

		if liveAgent <= targetCountAgent {
			logger.Debug("All agent is dead")
			break
		}
		logger.Debug("Update stat")
		//update model stat
		handlerStat(m)
		m.Statistic.Year++
		if m.Statistic.Year > maxAge && maxAge > 0 {
			break
		}
	}
	handlerUpdate(m)
	logger.Info("Model stopped")
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

func (m *Model) Fitness(countAgent int) {

	//я думаю эта функция будет не раз и не два меня обманывать и плодить баги

	sings := make([]*Sing, 8)
	//sort agent to age
	m.sort()
	//get best agent age
	for i := 0; i < countAgent; i++ {
		sings = append(sings, &m.Agents[i].Sing)
	}
	//delete old agent
	m.Agents = nil
	//make new agent with best sings
	agents := make([]*Agent, countAgent*len(sings))
	for s := 0; s < len(sings); s++ {
		for a := 0; a < countAgent; a++ {
			agents[s*a+a] = NewAgent(
				m.MaxX,
				m.MaxY,
				strconv.Itoa(s)+"-"+strconv.Itoa(a),
				sings[s])
		}
	}
	m.Agents = agents
	//todo: mutation sings
	for i := 0; i < countAgent; i++ {
		m.Agents[i].mutation(1)
	}
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
	best := m.Agents[0]
	for _, agent := range m.Agents {
		if best.Age < agent.Age {
			best = agent
		}
	}
	return &best.Sing
}

func (m *Model) resourceHandler() {
	for _, cells := range m.Map {
		for _, cell := range cells {

			//если в клетке есть max отравление - убрать всё
			if cell.PoisonLevel > m.maxPoisonLevel {
				cell.FoodLevel = 0
				continue
			}
			//если в клетке есть отравление /2 от максимального - оставить четверь
			if cell.PoisonLevel > m.maxPoisonLevel/2 {
				cell.FoodLevel /= 4
				continue
			}
			//если в клетке есть отравление /4 от максимального, оставить половину
			if cell.PoisonLevel > m.maxPoisonLevel/4 {
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

func (m *Model) sort() {
	for i := 0; i < len(m.Agents)-2; i++ {
		for j := 0; j < len(m.Agents)-2-i; j++ {
			if m.Agents[j+1].Age < m.Agents[j].Age {
				m.Agents[j+1], m.Agents[j] = m.Agents[j], m.Agents[j+1]
			}
		}
	}
}
