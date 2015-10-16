package grid

import (
	"bytes"
	"fmt"
	"strconv"
)

//Grid is the struct that holds the a grids information
type Grid struct {
	Grid   [][]bool `json:"grid"`
	Width  int      `json:"width"`
	Height int      `json:"height"`
}

//NewGrid creates a new multidimensional slice to hold the grid
func NewGrid(width, height int) Grid {
	grid := make([][]bool, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return Grid{Grid: grid, Width: width, Height: height}
}

//Set allows us to set specific locations to different values
func (g Grid) Set(x, y int, val bool) {
	g.Grid[y][x] = val
	return
}

//State returns the state of the requested location
func (g Grid) State(x, y int) (state bool) {
	state = g.Grid[y][x]
	return
}

// GetNewYCoord is because Go doesn't do array index checking for multidimensional arrays (I think)
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

//GetNewXCoord see explanation above
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

//GetBoardString is to convert the board's CurrentGrid to a string
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

//Display is a helper function to display the board to stdout
//irrelevant for webapp
func (g Grid) Display() {
	fmt.Print(g.GetBoardString())
}

func getChar(state bool) string {
	if state {
		return "@"
	}
	return "."
}

//Encode is a function to encode a grid to JSON. This is probably superfluous
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

//Copy is to copy the values of src grid to the grid we are acting on
func (g Grid) Copy(src Grid) {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.Grid[j][i] = src.Grid[j][i]
		}
	}
}

//Clear is to reset grid
func (g Grid) Clear() {
	for j := 0; j < g.Height; j++ {
		for i := 0; i < g.Width; i++ {
			g.Set(i, j, false)
		}
	}
}
