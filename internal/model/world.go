package model

import (
	"fmt"
)

type World struct {
	MaxX int
	MaxY int
	Map  Map //y, x
}

type Cell struct {
	FoodLevel   int
	PoisonLevel int
}

type Map [][]*Cell

func NewWorld(x, y int) *World {
	Map := make(Map, y)
	for u := 0; u < y; u++ {
		Map[u] = make([]*Cell, x)
		for v := 0; v < x; v++ {
			Map[u][v] = &Cell{
				FoodLevel:   0,
				PoisonLevel: 0,
			}
		}
	}
	return &World{x, y, Map}
}

func (w *World) clear() {
	for _, cells := range w.Map {
		for _, cell := range cells {
			cell.FoodLevel = 0
			cell.PoisonLevel = 0
		}
	}
}

// setValue get food and poison count in range [0,maxInt).
// value less than 0 ignored.
func (w *World) setValue(food, poison int, c Cords) {
	if food > -1 {
		w.Map[c.Y][c.X].FoodLevel = food
	}
	if poison > -1 {
		w.Map[c.Y][c.X].PoisonLevel = poison
	}
}

func (w *World) getCell(c *Cords) (*Cell, error) {
	if c.Y > w.MaxY-1 || c.Y < 0 {
		return nil, fmt.Errorf("out of range: Y coord's")
	}
	if c.X > w.MaxX-1 || c.X < 0 {
		return nil, fmt.Errorf("out of range: X coord's")
	}
	return w.Map[c.Y][c.X], nil
}
