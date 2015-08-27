package main

import (
	"github.com/sbogacz/GameOfLife/service"
	"log"
	"net/http"
)

func main() {
	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":8081", router))

}
