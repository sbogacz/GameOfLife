package service

import (
	"net/http"
)

type Route struct {
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetNumGames",
		"GET",
		"/getNumGames",
		GetNumActiveGames,
	},
	Route{
		"InitGame",
		"PUT",
		"/initGame/{gameType}",
		InitGame,
	},
	Route{
		"GetGame",
		"GET",
		"/getGame/{gameId:[0-9]+}",
		GetGameBoard,
	},
	Route{
		"StepGame",
		"PUT",
		"/stepGame/{gameId:[0-9]+}",
		StepGame,
	},
	Route{
        "GetGameJSON",
        "GET",
        "/getGameJSON/{gameId:[0-9]+}",
        GetGameBoardJSON,
    },
}
