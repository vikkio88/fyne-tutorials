package ui

import (
	"strings"
)

type AppRoute uint8

const (
	main    string = "MAIN"
	list    string = "LIST"
	newGame string = "NEW_GAME"
	quit    string = "QUIT"

	invalid string = "INVALID_ROUTE"
)

const (
	Main AppRoute = iota
	List
	NewGame

	Quit
)

func getMapping() map[AppRoute]string {
	return map[AppRoute]string{
		Main:    main,
		List:    list,
		NewGame: newGame,
		Quit:    quit,
	}
}

func getReverseMapping() map[string]AppRoute {
	return map[string]AppRoute{
		main:    Main,
		list:    List,
		newGame: NewGame,
		quit:    Quit,
	}
}

func RouteFromString(route string) AppRoute {
	route = strings.ToUpper(route)
	mapping := getReverseMapping()
	if val, ok := mapping[route]; ok {
		return val
	}

	return Main
}

func (a AppRoute) String() string {
	mapping := getMapping()
	if val, ok := mapping[a]; ok {
		return val
	}

	return invalid
}
