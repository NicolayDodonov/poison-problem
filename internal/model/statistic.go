package model

import "strconv"

type Statistic struct {
	Food      int
	Poison    int
	AvgEnergy int
	Year      int
	Sing
}

func (s *Statistic) String() string {
	str := strconv.Itoa(s.Year) + "; " +
		strconv.Itoa(s.Food) + "; " +
		strconv.Itoa(s.Poison) + "; " +
		strconv.Itoa(s.MoveOrAction) + "; " +
		strconv.Itoa(s.TurnOrMove) + "; " +
		strconv.Itoa(s.LeftOrRight) + "; " +
		strconv.Itoa(s.EatOrClear[0]) + "; " +
		strconv.Itoa(s.EatOrClear[1]) + "; " +
		strconv.Itoa(s.GetFood) + "; " +
		strconv.Itoa(s.GetPoison) + "; " +
		strconv.Itoa(s.MakePoison) + "; "
	return str
}
