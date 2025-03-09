package model

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
// todo: add context
func (m *Model) Run() {
	for _, agent := range m.Agents {
		err := agent.Run(m.World)
		if err != nil {
			//todo: l.ERROR(err)
		}
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
