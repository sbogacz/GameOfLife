package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/sbogacz/GameOfLife/grid"
)

//for this POC just maintain 20 grids in memory at most.
//for further development would add db representation of games
var (
	activeGames = make(grid.LifeGrids, 20)
	numGames    = 0
)

//InitGame initialize a new game, either random or with the classic glider
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
	gameID := numGames % 20
	activeGames[gameID] = *l
	//max 20 games
	numGames++

	gameIDJSON, err := json.Marshal(gameID)
	if err != nil {
		log.WithField("Error", err).Error("Error trying to marshal gameID to JSON")
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.WriteHeader(http.StatusOK)
	writer.Write(gameIDJSON)
}

//GetGameBoard simply prints the state of the board to stdout
func GetGameBoard(writer http.ResponseWriter, request *http.Request) {
	var gameID int
	var err error
	vars := mux.Vars(request)
	if gameID, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameID < 0 || gameID >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	log.WithField("Game ID", gameID).Info("Looking up Game of Life")
	layout := activeGames[gameID].CurrentGrid.GetBoardString()
	fmt.Fprintf(writer, layout)

}

//StepGame takes a game id and evolves the state of the game by one step
//eventually could maybe also pass # steps to evolve it
func StepGame(writer http.ResponseWriter, request *http.Request) {
	var gameID int
	var err error
	vars := mux.Vars(request)
	if gameID, err = strconv.Atoi(vars["gameId"]); err != nil {
		panic(err)
	}
	if gameID < 0 || gameID > numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	activeGames[gameID].Step()
	var layout []byte
	if layout, err = json.Marshal(activeGames[gameID].CurrentGrid); err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.WriteHeader(http.StatusOK)
	writer.Write(layout)
}

//GetNumActiveGames returns the number of games in the activeGames array
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

//GetGameBoardJSON returns the status of the board as a JSON arrau
func GetGameBoardJSON(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	gameID, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if gameID < 0 || gameID >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.WriteHeader(http.StatusOK)
	log.WithField("Game ID", gameID).Info("Looking up Game of Life")

	layout, err := json.Marshal(activeGames[gameID].CurrentGrid)
	if err != nil {
		panic(err)
	}
	writer.Write(layout)
}

//UpdateGameBoard updates the board with the given JSON
func UpdateGameBoard(writer http.ResponseWriter, request *http.Request) {
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
	gameID, err := strconv.Atoi(vars["gameId"])
	if err != nil {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if gameID < 0 || gameID >= numGames {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(strings.NewReader(jsonRequest))
	if err = decoder.Decode(&updatedGrid); err != nil {
		log.Error(err)
	}

	activeGames[gameID].CurrentGrid.Copy(updatedGrid)

	writer.Header().Set("Content-Type", "application/json;charset-UTF-8")
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	writer.WriteHeader(http.StatusOK)
}
