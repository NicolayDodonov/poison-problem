package model

import (
	"fmt"
	"poison-problem/internal/config"
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

func NewWorld(conf *config.World) *World {
	Map := make(Map, conf.MaxY)
	for u := 0; u < conf.MaxY; u++ {
		Map[u] = make([]*Cell, conf.MaxX)
		for v := 0; v < conf.MaxX; v++ {
			Map[u][v] = &Cell{
				FoodLevel:   0,
				PoisonLevel: conf.PoisonLevel,
			}
		}
	}
	return &World{conf.MaxX, conf.MaxY, Map}
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
