package grid

import (
	"bytes"
	"fmt"
	"strconv"
)

type Grid struct {
	Grid   [][]bool `json:"grid"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

func NewGrid(width, height int) Grid {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return Grid{Grid: grid, Width: width, Height: height}
}

func (g Grid) Set(x, y int, val bool) {
	g.Grid[y][x] = val
}

func (g Grid) State(x, y int) (state bool) {
	state = g.Grid[y][x]
	return
}

// Go doesn't do array index checking for multidimensional arrays (I think)
// so make helper functions to use later on
func (g Grid) GetNewYCoord(coord int) (newY int) {
	if coord < 0 {
		newY = g.Height - 1
	} else if coord >= g.Height {
		newY = 0
	} else {
		newY = coord
	}
	return
}

func (g Grid) GetNewXCoord(coord int) (newX int) {
	if coord < 0 {
		newX = g.Width - 1
	} else if coord >= g.Width {
		newX = 0
	} else {
		newX = coord
	}
	return
}

func (g Grid) GetBoardString() string {
	var buffer bytes.Buffer
	//first do the outer loop (height)
	for i := 0; i < g.Height; i++ {
		//start with a tab, to make it look better
		buffer.WriteString("\t")
		//next do inner loop, for rows
		for j := 0; j < g.Width; j++ {
			buffer.WriteString(getChar(g.Grid[i][j]))
		}
		//end line
		buffer.WriteString("\n")
	}
	//create some separation between next display
	buffer.WriteString("\n")
	buffer.WriteString("\n")
	return buffer.String()
}

func (g Grid) Display() {
	fmt.Print(g.GetBoardString())
}

func getChar(state bool) string {
	if state {
		return "@"
	}
	return "."
}

func (g Grid) Encode() string {
	var buffer bytes.Buffer
	buffer.WriteString("{\"grid\":{\n")
	buffer.WriteString("\t\"rows\":[\n")
	//first do the outer loop
	for i := 0; i < g.Height; i++ {
		buffer.WriteString("\t\t[")
		for j := 0; j < g.Width; j++ {
			//buffer.WriteString("\"")
			buffer.WriteString(strconv.FormatBool(g.Grid[i][j]))
			//buffer.WriteString("\"")
			buffer.WriteString(",")
		}
		//don't want to end up with an extra comma
		buffer.Truncate(buffer.Len() - 1)
		buffer.WriteString("],\n")
	}
	buffer.Truncate(buffer.Len() - 2)
	buffer.WriteString("]}}")
	return buffer.String()
}

func (g Grid) Copy(src Grid) {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.Grid[j][i] = src.Grid[j][i]
		}
	}
}

//reset grid
func (g Grid) Clear() {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.Set(i, j, false)
		}
	}
}
