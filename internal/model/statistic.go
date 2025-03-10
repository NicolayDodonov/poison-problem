package model

import "strconv"

type Statistic struct {
	Food      int
	Poison    int
	AvgEnergy int
	Sings
}

func (s *Statistic) String() string {
	str := strconv.Itoa(s.Food) + "; " +
		strconv.Itoa(s.Poison) + "; " +
		strconv.Itoa(s.moveOrAction) + "; " +
		strconv.Itoa(s.turnOrMove) + "; " +
		strconv.Itoa(s.leftOrRight) + "; " +
		strconv.Itoa(s.eatOrClear[0]) + "; " +
		strconv.Itoa(s.eatOrClear[1]) + "; " +
		strconv.Itoa(s.getFood) + "; " +
		strconv.Itoa(s.getPoison) + "; " +
		strconv.Itoa(s.makePoison) + "; "
	return str
}
