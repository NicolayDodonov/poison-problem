package model

// turn change in range (0 to 7)
/*
0 → x
	1 2 3
	0 * 4
	7 6 5
*/
type turn int

func (t *turn) left() {
	*t--
}

func (t *turn) right() {
	*t++
}

type Cords struct {
	X int
	Y int
}

// setCoordsToTurns change Cords with turn factor
func (c *Cords) setCoordsToTurns(t turn) {
	switch t % 8 {
	case 0:
		c.X--
	case 1:
		c.X--
		c.Y--
	case 2:
		c.Y--
	case 3:
		c.X++
		c.Y--
	case 4:
		c.X++
	case 5:
		c.X++
		c.Y++
	case 6:
		c.Y++
	case 7:
		c.X--
		c.Y++
	}
}

func (c *Cords) setCordsToTurnsWithWorld(t turn, w *World) {
	//copy cords to nCords
	nCord := &Cords{
		c.X,
		c.Y,
	}

	nCord.setCoordsToTurns(t)
	//nCords fix with world.MaxN
	if nCord.X > w.MaxX {
		nCord.X = nCord.X % w.MaxX
	}
	if nCord.X < 0 {
		nCord.X = w.MaxX + nCord.X
	}
	if nCord.Y > w.MaxY {
		nCord.Y = nCord.Y % w.MaxY
	}
	if nCord.Y < 0 {
		nCord.Y = w.MaxY + nCord.Y
	}

	c.X = nCord.X
	c.Y = nCord.Y
}

// multiplyCordOnTurns copy Cords with turn factor
func multiplyCordOnTurns(t turn, c Cords) *Cords {
	c.setCoordsToTurns(t)
	return &Cords{
		c.X,
		c.Y,
	}
}
