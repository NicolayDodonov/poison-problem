package simulation

import (
	"log"
	"os"
	"poison-problem/internal/model"
	"strconv"
	"strings"
)

// foodPoisonCounter update model.Statistic fields: Food, Poison, AvgEnergy
func foodPoisonCounter(m *model.Model) {
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
}

// avgCounter update model.Statistic fields: Sing
func avgCounter(m *model.Model) {
	m.Statistic.Sing = model.Sing{
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

// todo: make save date to file more standard
// singsCounter counts live agents all Sing-group
func singsCounter(m *model.Model) {
	//use base handler
	foodPoisonCounter(m)

	//make map
	s := make(map[string]int)
	for i := 0; i < m.Parameters.CountSings; i++ {
		s[strconv.Itoa(i)] = 0
	}
	s["food"] = m.Food
	s["poison"] = m.Poison
	//range all agent in model
	for _, agent := range m.Agents {
		//if agent live - save data about this agent
		if agent.Energy > 0 {
			key := strings.Split(agent.ID, "-")
			if _, ok := s[key[0]]; ok {
				//if ok = true - increment value
				s[key[0]]++
			} else {
				//if ok = false - append key
				s[key[0]] = 1
			}
		}
	}
	//convert s to string
	sb := strings.Builder{}
	for i := 0; i < m.Parameters.CountSings; i++ {
		sb.WriteString(strconv.Itoa(s[strconv.Itoa(i)]) + "; ")
	}
	sb.WriteString(strconv.Itoa(s["food"]) + "; ")
	sb.WriteString(strconv.Itoa(s["poison"]) + "; ")
	sb.WriteString("\n")

	//open file
	f, _ := os.OpenFile("saves/experiment.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	//save data
	msg := sb.String()

	if _, err := f.WriteString(msg); err != nil {
		//it not stable!!!!!!!!!!!!!!!
		log.Printf("singsCounter " + err.Error())
	}
}

func firstString(m *model.Model) {
	sb := strings.Builder{}
	for i := 0; i < m.Parameters.CountSings; i++ {
		sb.WriteString("s-" + strconv.Itoa(i) + "; ")
	}
	sb.WriteString(" food;")
	sb.WriteString(" poison;\n")

	f, _ := os.OpenFile("saves/experiment.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	//save data
	msg := sb.String()

	if _, err := f.WriteString(msg); err != nil {
		//it not stable!!!!!!!!!!!!!!!
		log.Printf("singsCounter " + err.Error())
	}
}

func todoHandler(m *model.Model) {}
