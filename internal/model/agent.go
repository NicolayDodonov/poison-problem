package model

import (
	"fmt"
	"math/rand/v2"
)

type Agent struct {
	ID     int
	Energy uint
	Look   turn  //0 to 8
	Cords  Cords //x y
	Sings        //small gen model
}

type Sings struct {
	moveOrAction int //determinate gone of cell or make some action
	turnOrMove   int //determinate make turns or go ahead
	leftOrRight  int //determinate type of turns
	getFood      int //determinate count of eat food
	getPoison    int //determinate count of eat poison
	makePoison   int //determinate count of produce poison
}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) Run(w *World) error {
	cellHaveItem, err := a.look(w)
	if err != nil {
		return fmt.Errorf("can't look in cell, because ", err)
	}

	if cellHaveItem {
		//action or go out?
		if rand.IntN(100) > a.moveOrAction {
			//Action
			err := a.action()
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

	return nil
}

// look can search nearby object and change vector
func (a *Agent) look(w *World) (bool, error) {
	//get data of cell
	cell, err := w.getCell(*multiplyCordOnTurns(a.Look, a.Cords))
	if err != nil {
		return false, fmt.Errorf("cant get cell. ", err)
	}
	//if cell have any item - moveOrAction
	if cell.FoodLevel > 10 || cell.PoisonLevel > 100 {
		return true, nil
	}
	//if cell don't have any item - pass
	return false, nil
}

// move change cords agent to vector
func (a *Agent) move(w *World) {
	if rand.IntN(100) > a.turnOrMove {
		// move to view cell
		a.Cords.setCordsToTurnsWithWorld(a.Look, w)
	} else {
		if rand.IntN(100) > a.leftOrRight {
			a.Look.right()
		} else {
			a.Look.left()
		}
	}
}

func (a *Agent) action() error {

	return nil
}

// clean destroy poison
func (a *Agent) clean() {

}

// death kill this Agent
func (a *Agent) death() {

}
