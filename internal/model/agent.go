package model

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Agent struct {
	ID     int
	Energy int
	Age    int
	Look   turn  //0 to 8
	Cords  Cords //x y
	Sing         //small gen model
}

type Sing struct {
	MoveOrAction int    `json:"move_or_action"` //determinate gone of cell or make some action
	TurnOrMove   int    `json:"turn_or_move"`   //determinate make turns or go ahead
	LeftOrRight  int    `json:"left_or_right"`  //determinate type of turns
	EatOrClear   [2]int `json:"eat_or_clear"`   //determinate the range (0 to [0]) ([0] to [1]) and ([1] to 100)
	GetFood      int    `json:"get_food"`       //determinate count of eat food
	GetPoison    int    `json:"get_poison"`     //determinate count of eat poison
	MakePoison   int    `json:"make_poison"`    //determinate count of produce poison
}

func NewAgent(MaxX, MaxY int, s *Sing) *Agent {
	return &Agent{
		time.Now().Nanosecond(),
		100,
		0,
		0,
		Cords{
			rand.IntN(MaxX),
			rand.IntN(MaxY),
		},
		*s,
	}
}

func (a *Agent) Run(w *World) error {
	if a.Energy <= 0 {
		return nil
	}
	//take energy price
	a.Energy--
	a.Age++

	//take the cell that the agent is looking at
	cell, err := w.getCell(a.Cords.getCordsOnViewWithWorld(a.Look, w))
	if err != nil {
		return fmt.Errorf("can't get cell on look, because %e", err)
	}

	if a.look(cell) {
		//action or go out?
		if rand.IntN(100) > a.MoveOrAction {
			//Action
			err := a.action(cell)
			if err != nil {
				return fmt.Errorf("can't some action, because %e", err)
			}
		} else {
			//go out
			a.move(w)
		}
	} else {
		//leave from empty
		a.move(w)
	}

	err = a.pollute(w)
	if err != nil {
		return fmt.Errorf("can't pollute, because %e", err)
	}

	return nil
}

// look can search nearby object and change vector
func (a *Agent) look(cell *Cell) bool {
	//if cell have any item - MoveOrAction
	if cell.FoodLevel > 10 || cell.PoisonLevel > 10 {
		return true
	}
	//if cell don't have any item - pass
	return false
}

// move change cords agent to vector
func (a *Agent) move(w *World) {
	if rand.IntN(100) > a.TurnOrMove {
		// move to view cell
		a.Cords = *a.Cords.getCordsOnViewWithWorld(a.Look, w)
		a.Energy--
	} else {
		if rand.IntN(100) > a.LeftOrRight {
			a.Look.right()
		} else {
			a.Look.left()
		}
	}
}

func (a *Agent) action(cell *Cell) error {
	n := rand.IntN(100)
	if n <= a.EatOrClear[0] {
		//eat
		a.eat(cell)
	} else if n <= a.EatOrClear[1] {
		//eat & clear
		a.eat(cell)
		a.clean(cell)
	} else {
		//clear
		a.clean(cell)
	}
	return nil
}

// eat destroy food
func (a *Agent) eat(cell *Cell) {
	eatenFood := cell.FoodLevel * a.GetFood / 100
	cell.FoodLevel -= eatenFood
	a.Energy += eatenFood
}

// clean destroy poison
func (a *Agent) clean(cell *Cell) {
	eatenPoison := cell.PoisonLevel * a.GetPoison / 100
	cell.PoisonLevel -= eatenPoison
	a.Energy += eatenPoison / 2
}

func (a *Agent) pollute(w *World) error {
	if a.Cords.Y >= 0 && a.Cords.Y < w.MaxY &&
		a.Cords.X >= 0 && a.Cords.X < w.MaxX {
		w.Map[a.Cords.Y][a.Cords.X].PoisonLevel += a.MakePoison

		if w.Map[a.Cords.Y][a.Cords.X].PoisonLevel > 50 {
			a.Energy -= 1
		} else if w.Map[a.Cords.Y][a.Cords.X].PoisonLevel > 75 {
			a.Energy -= 5
		} else if w.Map[a.Cords.Y][a.Cords.X].PoisonLevel > 100 {
			a.Energy -= 25
		}

		return nil
	} else {
		return fmt.Errorf("out of world! X: %q Y: %q", a.Cords.X, a.Cords.Y)
	}
}

func (s *Sing) mutation(mutation int) {
	/*
		0 	MoveOrAction int    //determinate gone of cell or make some action
		1 	TurnOrMove   int    //determinate make turns or go ahead
		2 	LeftOrRight  int    //determinate type of turns
		3,4 EatOrClear   [2]int //determinate the range (0 to [0]) ([0] to [1]) and ([1] to 100)
		5 	GetFood      int    //determinate count of eat food
		6 	GetPoison    int    //determinate count of eat poison
		7 	MakePoison   int    //determinate count of produce poison
	*/
	for i := 0; i < mutation; i++ {
		var n int
		if rand.IntN(2) == 1 {
			n = 1
		} else {
			n = -1
		}

		switch rand.IntN(8) {
		case 0:
			s.MoveOrAction += n
			if s.MoveOrAction < 0 {
				s.MoveOrAction = 0
			}
			if s.MoveOrAction > 100 {
				s.MoveOrAction = 100
			}
		case 1:
			s.TurnOrMove += n
			if s.TurnOrMove < 0 {
				s.TurnOrMove = 0
			}
			if s.TurnOrMove > 99 {
				s.TurnOrMove = 99
			}
		case 2:
			s.LeftOrRight += n
			if s.LeftOrRight < 0 {
				s.LeftOrRight = 0
			}
			if s.LeftOrRight > 99 {
				s.LeftOrRight = 99
			}
		case 3:
			s.EatOrClear[0] += n
			if s.EatOrClear[0] < 0 {
				s.EatOrClear[0] = 0
			}
			if s.EatOrClear[0] > s.EatOrClear[1] {
				s.EatOrClear[0] = s.EatOrClear[1] - 1
			}
		case 4:
			s.EatOrClear[1] += n
			if s.EatOrClear[1] < s.EatOrClear[0] {
				s.EatOrClear[1] = s.EatOrClear[0] + 1
			}
			if s.EatOrClear[1] > 99 {
				s.EatOrClear[1] = 99
			}
		case 5:
			s.GetFood += n
			if s.GetFood < 10 {
				s.GetFood = 10
			}
			if s.GetFood > 100 {
				s.GetFood = 100
			}
		case 6:
			s.GetPoison += n
			if s.GetPoison < 10 {
				s.GetPoison = 10
			}
			if s.GetPoison > 100 {
				s.GetPoison = 100
			}
		case 7:
			s.MakePoison += n
			if s.MakePoison < 1 {
				s.MakePoison = 1
			}
			if s.MakePoison > 10 {
				s.MakePoison = 10
			}
		}
	}
}
