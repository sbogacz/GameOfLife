package main

import (
	//    "fmt"
	//    "time"
	//    "github.com/sbogacz/game_of_life/grid"
	"github.com/sbogacz/game_of_life/service"
	"log"
	"net/http"
)

/*func main() {
	l := grid.NewLifeGrid(30, 12)
	//start with the classic glider
	//l.RandomInit()
	l.GetGlider(20,10)
	l.CurrentGrid.Display()

	for i := 0; i < 2; i++ {
		l.Step()
		fmt.Print("\x0c")
		//l.CurrentGrid.Display()
		fmt.Print(l.CurrentGrid.Encode())
		time.Sleep(time.Second/4)
	}
}*/

func main() {
	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
