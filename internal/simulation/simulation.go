package simulation

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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

type Sings struct {
	Sings []*model.Sing
}

func New(logger *logger.Logger, Conf *config.Simulation) *Simulation {
	return &Simulation{
		Conf,
		logger,
	}
}

func (s Simulation) Run() {
	//todo: todo: sings := s.LoadSings()
	sing := &model.Sing{
		50,
		50,
		50,
		[2]int{25, 75},
		50,
		50,
		2,
	}

	switch strings.ToLower(s.Conf.Type) {
	case "train":
		//todo: and run len(sings) train with 1 sing
		s.train(s.Conf.TargetAge, sing)
	case "experiment":
		//todo: and run 1 experiment with all sings
		s.experiment(s.Conf.MaxEpoch)
	}
}

func (s Simulation) train(targetAge int, sings *model.Sing) {

	// make model to train sings
	m := model.New(
		s.Conf.StartAgent,
		20,
		20,
		sings)

	for {
		//todo: m.ClearModel

		// run one epoch model
		// epoch end if all agent ded
		m.Run(context.TODO(), s.Log, s.Conf.EndAgent)

		// after end epoch - save statistic in file
		if err := s.SaveStatistic(m.String()); err != nil {
			s.Log.Error(err.Error())
		}

		// check the exit conditions for the target age
		if m.CheckTargetAge(targetAge) {
			break
		}

		// If the conditions are not met,
		// we start mutation and select the best agents by age.
		m.Fitness()
	}

	s.SaveSing(&m.BestSing().String())
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

func (s Simulation) SaveStatistic(stat string) error {
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

func (s Simulation) SaveSing(sing string) error {
	file, err := os.OpenFile(s.Conf.SaveSing, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err = file.WriteString(sing); err != nil {
		return err
	}
	return nil
}

func (s Simulation) LoadStatistic() (*[]*model.Sing, error) {
	if s.Conf.LoadSing == "" {
		return nil, fmt.Errorf("cannot load sing file")
	}

	var data Sings
	if err := cleanenv.ReadConfig(s.Conf.LoadSing, &data); err != nil {
		log.Fatalf("cannot read sing file: %s", err)
	}

	return &data.Sings, nil
}
