package grid

import (
	"math/rand"
)

//		Game of Life Rules:
//    Any live cell with fewer than two live neighbours dies, as if caused by under-population.
//    Any live cell with two or three live neighbours lives on to the next generation.
//    Any live cell with more than three live neighbours dies, as if by overcrowding.
//    Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

//Next is a function to compute the next state for the grid according to the rules
func (g Grid) Next(x, y int) (nextState bool) {
	neighborCount := 0
	for i := -1; i <= 1; i++ {
		newY := g.GetNewYCoord(y + i)
		for j := -1; j <= 1; j++ {
			newX := g.GetNewXCoord(x + j)
			if j == 0 && i == 0 {
				continue
			}
			if g.State(newX, newY) {
				neighborCount++
			}
		}
	}
	nextState = neighborCount == 3 || neighborCount == 2 && g.State(g.GetNewXCoord(x), g.GetNewYCoord(y))
	return
}

type LifeGrid struct {
	CurrentGrid, nextGrid Grid
	width, height         int
}

type LifeGrids []LifeGrid

func NewLifeGrid(width, height int) *LifeGrid {
	a := NewGrid(width, height)
	return &LifeGrid{
		CurrentGrid: a,
		nextGrid:    NewGrid(width, height),
		width:       width, height: height,
	}
}

func (lg *LifeGrid) RandomInit() {
	for i := 0; i < (lg.width * lg.height / 2); i++ {
		lg.CurrentGrid.Set(rand.Intn(lg.width), rand.Intn(lg.height), true)
	}
}

func (lg *LifeGrid) Step() {
	for j := 0; j < lg.height; j++ {
		for i := 0; i < lg.width; i++ {
			nextState := lg.CurrentGrid.Next(i, j)
			lg.nextGrid.Set(i, j, nextState)
		}
	}
	lg.CurrentGrid, lg.nextGrid = lg.nextGrid, lg.CurrentGrid
}

//...@@@...
//...@.....
//....@....
func (lg LifeGrid) GetGlider(x, y int) {
	lg.CurrentGrid.Clear()
	lg.CurrentGrid.Set(lg.CurrentGrid.GetNewXCoord(x-1), lg.CurrentGrid.GetNewYCoord(y-1), true)
	lg.CurrentGrid.Set(lg.CurrentGrid.GetNewXCoord(x), lg.CurrentGrid.GetNewYCoord(y-1), true)
	lg.CurrentGrid.Set(lg.CurrentGrid.GetNewXCoord(x+1), lg.CurrentGrid.GetNewYCoord(y-1), true)
	lg.CurrentGrid.Set(lg.CurrentGrid.GetNewXCoord(x-1), lg.CurrentGrid.GetNewYCoord(y), true)
	lg.CurrentGrid.Set(lg.CurrentGrid.GetNewXCoord(x), lg.CurrentGrid.GetNewYCoord(y+1), true)
}
