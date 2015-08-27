package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sbogacz/GameOfLife/grid"
	"net/http"
	"strconv"
	"strings"
)

//for this POC just maintain 20 grids in memory at most.
//for further development would add db representation of games
var activeGames grid.LifeGrids = make(grid.LifeGrids, 20)
var numGames = 0

//initialize a new game, either random or with the classic glider
func InitGame(writer http.ResponseWriter, request *http.Request) {
	l := grid.NewLifeGrid(40, 20)
	vars := mux.Vars(request)
	gameType := vars["gameType"]

	if gameType == "classic" {
		//start with the classic glider
		l.GetGlider(20, 10)
	} else if gameType == "random" {
		l.RandomInit()
	} else {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	//the modulus is so that we never try to write beyond the allocated size of the slice
	gameId := numGames % 20
	activeGames[gameId] = *l
	//max 20 games
	numGames++

	var err error
	var gameIdJSON []byte
	if gameIdJSON, err = json.Marshal(gameId); err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.WriteHeader(http.StatusOK)
	writer.Write(gameIdJSON)
	//fmt.Fprintf(writer, "{ gameID : %d}", gameId)
}

//right now we're just printing the state of the board, but once we
//start integrating the front-end, we can use the grid.Encode() function
//to pass it to Angular
func GetGameBoard(writer http.ResponseWriter, request *http.Request) {
	var gameId int
	var err error
	vars := mux.Vars(request)
	if gameId, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameId < 0 || gameId >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(writer, "Looking up Game of Life with id = %d\n", gameId)
	layout := activeGames[gameId].CurrentGrid.GetBoardString()
	fmt.Fprintf(writer, layout)

}

//takes a game id and evolves the state of the game by one step
//eventually could maybe also pass # steps to evolve it
func StepGame(writer http.ResponseWriter, request *http.Request) {
	var gameId int
	var err error
	vars := mux.Vars(request)
	if gameId, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameId < 0 || gameId > numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	activeGames[gameId].Step()
	var layout []byte
	if layout, err = json.Marshal(activeGames[gameId].CurrentGrid); err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.WriteHeader(http.StatusOK)
	writer.Write(layout)
}

//returns the number of games in the activeGames array
func GetNumActiveGames(writer http.ResponseWriter, request *http.Request) {
	if numGames != 1 {
		if numGames > 20 {
			fmt.Fprintf(writer, "There are %d active games", 20)
		} else {
			fmt.Fprintf(writer, "There are %d active games", numGames)
		}
	} else {
		fmt.Fprintf(writer, "There is %d active game\n", numGames)
	}
}

func GetGameBoardJSON(writer http.ResponseWriter, request *http.Request) {
	var gameId int
	var err error
	vars := mux.Vars(request)
	if gameId, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameId < 0 || gameId >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	//fmt.Fprintf(writer, "Looking up Game of Life with id = %d\n", gameId)
	//layout := activeGames[gameId].CurrentGrid.Encode()
	var layout []byte
	if layout, err = json.Marshal(activeGames[gameId].CurrentGrid); err != nil {
		panic(err)
	}
	writer.Write(layout)
	//fmt.Fprintf(writer, layout)
}

func UpdateGameBoard(writer http.ResponseWriter, request *http.Request) {
	var gameId int
	var err error
	updatedGrid := grid.Grid{}
	buffer := new(bytes.Buffer)
	buffer.ReadFrom(request.Body)
	jsonRequest := strings.TrimSpace(buffer.String())

	if request.Method == "OPTIONS" {
		fmt.Println("no json data received")
		writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "PUT")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		writer.WriteHeader(http.StatusOK)
		return
	}
	vars := mux.Vars(request)
	if gameId, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameId < 0 || gameId >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(strings.NewReader(jsonRequest))
	if err = decoder.Decode(&updatedGrid); err != nil {
		fmt.Printf(err.Error())
	}

	activeGames[gameId].CurrentGrid.Copy(updatedGrid)

	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.WriteHeader(http.StatusOK)
}
