package main

import "github.com/joram/jsnek/api"

var (
	exampleBoard = &api.Board{
		Width:20,
		Height:20,
		Food: []api.Coord{{0,0},{19,19},},
		Snakes: []api.Snake{
			{"snake1", "snake 1", 50, []api.Coord{{5,5},{5,6}}},
			{"snake2", "snake 2", 50, []api.Coord{{8,5},{8,6}}},
		},
	}
)