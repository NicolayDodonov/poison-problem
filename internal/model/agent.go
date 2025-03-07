package model

// todo: struct agent
type Agent struct {
	ID     int
	Energy uint
	Look   uint
	Cords  Cords
	Sings  Sings
}

type Cords struct {
	X uint
	Y uint
}

type Sings struct {
	//todo: sing parameters
}

func (a *Agent) Run() {

}

// move change cords agent to vector
func (a *Agent) move() {

}

// look can search nearby object and change vector
func (a *Agent) look() {

}

// eat destroy resource
func (a *Agent) eat() {

}

// clean destroy poison
func (a *Agent) clean() {

}

// death kill this Agent
func (a *Agent) death() {

}
