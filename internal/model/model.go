package model

type Model struct {
	Agents []*Agent
	World
	Statistic
}

func New() *Model {
	//todo: init []*Agents

	//todo: init World

	//todo: init Statistic

	return &Model{}
}

func (m *Model) Run() {
	for _, agent := range m.Agents {
		agent.Run()
	}

	//todo: m.spawnFood()

	//todo: clear() ?

	m.updateStat()
}

func (m *Model) updateStat() {
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
