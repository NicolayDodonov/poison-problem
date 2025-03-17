package simulation

import (
	"context"
	"poison-problem/internal/logger"
	"poison-problem/internal/model"
)

// todo: sim struct
type Simulation struct {
	l logger.Logger
	//todo: other parameters
}

func (s Simulation) Run() {
	for {
		//model = New()
		m := model.New(100, 20, 20,
			&model.Sings{
				50,
				50,
				50,
				[2]int{25, 75},
				30,
				30,
				5,
			},
		)
		//model.Run
		m.Run(context.TODO(), &s.l, 8)
		//model.SaveStat
		m.SaveStatistic()
		//model.Fitness
		m.Fitness()

		//todo: Exit conditions
	}

}
