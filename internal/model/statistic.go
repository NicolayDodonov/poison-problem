package model

import "strconv"

type Statistic struct {
	//Count food in world now
	Food int
	//Count poison in world now
	Poison int
}

func (s *Statistic) String() string {
	str := "Food: " + strconv.Itoa(s.Food) + " Poison: " + strconv.Itoa(s.Poison)
	return str
}
