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

	switch strings.ToLower(s.Conf.Type) {
	case "train":
		s.Log.Info("Start train")

		//load one sing from file
		sing, err := s.loadSing()
		if err != nil {
			s.Log.Error(err.Error())
		} else {
			s.Log.Info("Sing loaded")
		}
		//if load sing is empty, set base sing
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
		//start train model
		s.train(s.Conf.TargetAge, sing)

	case "experiment":
		s.Log.Info("Start experiment")

		//load many sings from file
		s.Log.Info("Load sings from file")
		sings, err := s.loadSings()
		if err != nil {
			s.Log.Error(err.Error())
			return
		}
		//start make experiment
		s.experiment(sings)
	}

	s.Log.Info("Simulation end")
}

func (s Simulation) train(targetAge int, sing *model.Sing) {
	s.Log.Debug("Init model")
	// make model to train sings
	m := model.New(
		s.Conf.StartAgent,
		20,
		20,
		[]*model.Sing{sing})

	for {
		// run one epoch model
		// epoch end if all agent ded
		m.Run(context.TODO(), s.Log, s.Conf.EndAgent, foodPoisonCounter, avgCounter)

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

func (s Simulation) experiment(sings []*model.Sing) {
	//init model
	m := model.New(
		100,
		20,
		20,
		sings,
	)
	//run model
	m.Run(context.TODO(), s.Log, 0, singsCounter, todoHandler)
	s.Log.Info("Finished experiment")
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
		return nil, fmt.Errorf("cannot load sing, path is empty")
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

func (s Simulation) loadSings() ([]*model.Sing, error) {
	if s.Conf.LoadSings == "" {
		return nil, fmt.Errorf("cannot load sings, path is empty!")
	}

	data, err := os.ReadFile(s.Conf.LoadSings)
	if err != nil {
		return nil, err
	}

	type array struct {
		Sings []*model.Sing `json:"sings"`
	}

	sings := array{}

	err = json.Unmarshal(data, &sings)
	if err != nil {
		return nil, err
	}

	return sings.Sings, nil
}
