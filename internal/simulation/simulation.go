package simulation

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"poison-problem/internal/config"
	"poison-problem/internal/logger"
	"poison-problem/internal/model"
	"strings"
)

type Simulation struct {
	Conf *config.Simulation
	Log  *logger.Logger
}

func New(logger *logger.Logger, Conf *config.Simulation) *Simulation {
	return &Simulation{
		Conf,
		logger,
	}
}

func (s Simulation) Run() {

	s.Log.Info("Simulation init")

	//todo: rewrite this to many sings loader!
	sing, err := s.loadSing()
	if err != nil {
		s.Log.Error(err.Error())
	} else {
		s.Log.Info("Sings loaded")
	}
	if sing == nil {
		s.Log.Error("sing is empty! Load base sing.")
		sing = &model.Sing{
			50,
			50,
			50,
			[2]int{25, 75},
			50,
			50,
			2,
		}
	}

	switch strings.ToLower(s.Conf.Type) {
	case "train":
		s.Log.Info("Start train")
		//todo: and run len(sings) train with 1 sing
		s.train(s.Conf.TargetAge, sing)
	case "experiment":
		s.Log.Info("Start experiment")
		//todo: and run 1 experiment with all sings
		s.experiment(s.Conf.MaxEpoch)
	}
	s.Log.Info("Simulation end")
}

func (s Simulation) train(targetAge int, sings *model.Sing) {
	s.Log.Debug("Init model")
	// make model to train sings
	m := model.New(
		s.Conf.StartAgent,
		20,
		20,
		sings)

	for {
		// run one epoch model
		// epoch end if all agent ded
		m.Run(context.TODO(), s.Log, s.Conf.EndAgent)

		// after end epoch - save statistic in file
		if err := s.saveStat(m.String()); err != nil {
			s.Log.Error(err.Error())
		}

		// check the exit conditions for the target age
		if m.CheckTargetAge(targetAge) {
			break
		}

		s.Log.Debug("Start fitness function")
		// If the conditions are not met,
		// we start mutation and select the best agents by age.
		m.Fitness(targetAge)

		s.Log.Debug("Reset model")
		//clear world and reset agent
		m.Reset()
	}

	if err := s.saveSing(m.BestSing()); err != nil {
		s.Log.Error(err.Error())
	}
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

		m.Run(context.TODO(), s.Log, 0)

		//get ifo about of all sings group
		//todo: m.GetCountStatistic
	}
}

func (s Simulation) saveStat(stat string) error {
	file, err := os.OpenFile(s.Conf.SaveStat, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(stat); err != nil {
		return err
	}
	return nil
}

func (s Simulation) saveSing(sing *model.Sing) error {
	data, err := json.Marshal(sing)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(s.Conf.SaveSing, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, _ = file.Write(data)

	return nil
}

func (s Simulation) loadSing() (*model.Sing, error) {
	if s.Conf.LoadSing == "" {
		return nil, fmt.Errorf("cannot load sing, path is empty!")
	}

	data, err := os.ReadFile(s.Conf.LoadSing)
	if err != nil {
		return nil, err
	}

	var sing model.Sing

	err = json.Unmarshal(data, &sing)
	if err != nil {
		return nil, err
	}

	return &sing, nil
}
