package main

import (
	"log"
	"net/http"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	UNKNOWN
)

type  Responsibility  interface  {
	decision(*SnakeRequest) int
}
var (
	responsibilities = []Responsibility{
		// NO OPTION
		OnlyOneChoice{},
		// ONLY ONE NOT THREATENED CHOICE
		// HUNGRY (health level?) GO FOR FOOD
		// SHORTEST SNAKE GO FOR FOOD
		// POTENTIAL KILL
		// EAT THEIR LUNCH (force them to starve)
		// ANY NON SOLID MOVE
	}
	directions = []int{UP, DOWN, LEFT, RIGHT}
	directionStrings = map[int]string{
		UP: "up",
		DOWN: "down",
		LEFT: "left",
		RIGHT: "right",
	}
)


func Start(res http.ResponseWriter, req *http.Request) {
	decoded := SnakeRequest{}
	err := DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	sr := SnakeRequest{}
	err := DecodeSnakeRequest(req, &sr)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	for _, r := range responsibilities {
		choice := r.decision(&sr)
		if choice != UNKNOWN {
			respond(res, MoveResponse{Move: directionStrings[choice]})
			return
		}
	}

	respond(res, MoveResponse{Move: "down"})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}
