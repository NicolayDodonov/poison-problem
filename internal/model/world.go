package model

import (
	"math/rand"
)

type World struct {
	MaxX int
	MaxY int
	Year int
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
	return &World{x, y, 0, Map}
}

func (w *World) spawnFood(count int) {
	for {
		if count <= 0 {
			break
		}
		cords := Cords{
			uint(rand.Intn(w.MaxX)),
			uint(rand.Intn(w.MaxY)),
		}
		w.setValue(1, -1, cords)
		count--
	}
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
// value less then 0 ignored.
func (w *World) setValue(food, poison int, c Cords) {
	if food > -1 {
		w.Map[c.Y][c.X].FoodLevel = food
	}
	if poison > -1 {
		w.Map[c.Y][c.X].PoisonLevel = poison
	}
}
