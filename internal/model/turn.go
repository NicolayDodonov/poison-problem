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
	*t = (*t - 1) % 8
}

func (t *turn) right() {
	*t = (*t + 1) % 8
}

type Cords struct {
	X int
	Y int
}

// getCordsOnView base change coords
func (c *Cords) getCordsOnView(t turn) *Cords {
	n := Cords{
		c.X,
		c.Y,
	}
	switch t % 8 {
	case 0:
		n.X--
	case 1:
		n.X--
		n.Y--
	case 2:
		n.Y--
	case 3:
		n.X++
		n.Y--
	case 4:
		n.X++
	case 5:
		n.X++
		n.Y++
	case 6:
		n.Y++
	case 7:
		n.X--
		n.Y++
	}
	return &n
}

func (c *Cords) getCordsOnViewWithWorld(t turn, w *World) *Cords {
	n := Cords{
		c.X,
		c.Y,
	}
	switch t % 8 {
	case 0:
		n.X--
	case 1:
		n.X--
		n.Y--
	case 2:
		n.Y--
	case 3:
		n.X++
		n.Y--
	case 4:
		n.X++
	case 5:
		n.X++
		n.Y++
	case 6:
		n.Y++
	case 7:
		n.X--
		n.Y++
	}

	if n.X > w.MaxX {
		n.X = n.X % w.MaxX
	}
	if n.X < 0 {
		n.X = w.MaxX + n.X
	}
	if n.Y > w.MaxY {
		n.Y = n.Y % w.MaxY
	}
	if n.Y < 0 {
		n.Y = w.MaxY + n.Y
	}

	return &n
}
