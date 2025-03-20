package simulation

import (
	"context"
	"poison-problem/internal/logger"
	"poison-problem/internal/model"
)

// todo: sim struct
type Simulation struct {
	l *logger.Logger
	//todo: other parameters
}

func (s Simulation) Run() {
	//todo: change this stuff
	sing := &model.Sings{
		50,
		50,
		50,
		[2]int{25, 75},
		50,
		50,
		50,
	}
	if true {
		s.train(300, sing)
	} else {
		s.experiment(300)
	}
}

func (s Simulation) train(targetAge int, sings *model.Sings) {
	for {
		// make model to train sings
		m := model.New(
			100,
			20,
			20,
			sings)

		// run one epoch model
		// epoch end if all agent ded
		m.Run(context.TODO(), s.l, 8)

		// after end epoch - save statistic in file
		//todo: make special struct (not in model) or func to save this data
		m.SaveStatistic()

		// check the exit conditions for the target age
		if m.CheckTargetAge(targetAge) {
			break
		}

		// If the conditions are not met,
		// we start mutation and select the best agents by age.
		m.Fitness()
	}

	//todo: save best sing in file
	// exit
}

func (s Simulation) experiment(maxEpoch int) {
	for epoch := 0; epoch < maxEpoch; epoch++ {
		//todo: load sings from file
		// sings := s.loadSings()

		// make model to experiment
		m := model.New(
			100,
			20,
			20,
		)
		_ = m //todo: delete

		m.Run(context.TODO(), s.l, 0)

		//get ifo about of all sings group
		//todo: m.GetCountStatistic
	}
}
