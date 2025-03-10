package model

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type Agent struct {
	ID     int
	Energy int
	Look   turn  //0 to 8
	Cords  Cords //x y
	Sings        //small gen model
}

type Sings struct {
	moveOrAction int    //determinate gone of cell or make some action
	turnOrMove   int    //determinate make turns or go ahead
	leftOrRight  int    //determinate type of turns
	eatOrClear   [2]int //determinate the range (0 to [0]) ([0] to [1]) and ([1] to 100)
	getFood      int    //determinate count of eat food
	getPoison    int    //determinate count of eat poison
	makePoison   int    //determinate count of produce poison
}

func NewAgent(MaxX, MaxY int, s *Sings) *Agent {
	return &Agent{
		time.Now().Nanosecond(),
		100,
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

	//take the cell that the agent is looking at
	cell, err := w.getCell(a.Cords.getCordsOnViewWithWorld(a.Look, w))
	if err != nil {
		return fmt.Errorf("can't get cell on look, because ", err)
	}

	if a.look(cell) {
		//action or go out?
		if rand.IntN(100) > a.moveOrAction {
			//Action
			err := a.action(cell)
			if err != nil {
				return fmt.Errorf("can't some action, because ", err)
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
		return fmt.Errorf("can't pollute, because ", err)
	}

	return nil
}

// look can search nearby object and change vector
func (a *Agent) look(cell *Cell) bool {
	//if cell have any item - moveOrAction
	if cell.FoodLevel > 10 || cell.PoisonLevel > 10 {
		return true
	}
	//if cell don't have any item - pass
	return false
}

// move change cords agent to vector
func (a *Agent) move(w *World) {
	if rand.IntN(100) > a.turnOrMove {
		// move to view cell
		a.Cords = *a.Cords.getCordsOnViewWithWorld(a.Look, w)
		a.Energy--
	} else {
		if rand.IntN(100) > a.leftOrRight {
			a.Look.right()
		} else {
			a.Look.left()
		}
	}
}

func (a *Agent) action(cell *Cell) error {
	n := rand.IntN(100)
	if n <= a.eatOrClear[0] {
		//eat
		a.eat(cell)
	} else if n <= a.eatOrClear[1] {
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
	eatenFood := cell.FoodLevel * a.getFood / 100
	cell.FoodLevel -= eatenFood
	a.Energy += eatenFood
}

// clean destroy poison
func (a *Agent) clean(cell *Cell) {
	eatenPoison := cell.PoisonLevel * a.getPoison / 100
	cell.PoisonLevel -= eatenPoison
	a.Energy += eatenPoison / 2
}

func (a *Agent) pollute(w *World) error {
	if a.Cords.Y >= 0 && a.Cords.Y < w.MaxY &&
		a.Cords.X >= 0 && a.Cords.X < w.MaxX {
		w.Map[a.Cords.Y][a.Cords.X].PoisonLevel += a.makePoison

		if w.Map[a.Cords.Y][a.Cords.X].PoisonLevel > 50 {
			a.Energy--
		}

		return nil
	} else {
		return fmt.Errorf("out of world! X", a.Cords.X, "Y:", a.Cords.Y)
	}
}

func (a *Agent) mutation(mutation int) {
	/*
		0 	moveOrAction int    //determinate gone of cell or make some action
		1 	turnOrMove   int    //determinate make turns or go ahead
		2 	leftOrRight  int    //determinate type of turns
		3,4 eatOrClear   [2]int //determinate the range (0 to [0]) ([0] to [1]) and ([1] to 100)
		5 	getFood      int    //determinate count of eat food
		6 	getPoison    int    //determinate count of eat poison
		7 	makePoison   int    //determinate count of produce poison
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
			a.moveOrAction += n
			if a.moveOrAction < 0 {
				a.moveOrAction = 0
			}
			if a.moveOrAction > 100 {
				a.moveOrAction = 100
			}
		case 1:
			a.turnOrMove += n
			if a.turnOrMove < 0 {
				a.turnOrMove = 0
			}
			if a.turnOrMove > 99 {
				a.turnOrMove = 99
			}
		case 2:
			a.leftOrRight += n
			if a.leftOrRight < 0 {
				a.leftOrRight = 0
			}
			if a.leftOrRight > 99 {
				a.leftOrRight = 99
			}
		case 3:
			a.eatOrClear[0] += n
			if a.eatOrClear[0] < 0 {
				a.eatOrClear[0] = 0
			}
			if a.eatOrClear[0] > a.eatOrClear[1] {
				a.eatOrClear[0] = a.eatOrClear[1] - 1
			}
		case 4:
			a.eatOrClear[1] += n
			if a.eatOrClear[1] < a.eatOrClear[0] {
				a.eatOrClear[1] = a.eatOrClear[0] + 1
			}
			if a.eatOrClear[1] > 99 {
				a.eatOrClear[1] = 99
			}
		case 5:
			a.getFood += n
			if a.getFood < 10 {
				a.getFood = 10
			}
			if a.getFood > 100 {
				a.getFood = 100
			}
		case 6:
			a.getPoison += n
			if a.getPoison < 10 {
				a.getPoison = 10
			}
			if a.getPoison > 100 {
				a.getPoison = 100
			}
		case 7:
			a.makePoison += n
			if a.makePoison < 10 {
				a.makePoison = 10
			}
			if a.makePoison > 100 {
				a.makePoison = 100
			}
		}
	}
}
