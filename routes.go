package main

import (
	"log"
	"net/http"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/logic"
	"github.com/joram/jsnek/filters"
	"fmt"
)

var (
	responsibilities = []logic.Responsibility{
		// NO OPTION
		logic.OnlyOneChoice{},
		// ONLY ONE NOT THREATENED CHOICE
		// HUNGRY (health level?) GO FOR FOOD
		logic.GoEatOrthogonal{25},
		// SHORTEST SNAKE GO FOR FOOD
		// POTENTIAL KILL
		// EAT THEIR LUNCH (force them to starve)
		logic.TrapFood{},
		logic.ValidDirection{},
	}
	directionStrings = map[int]string{
		api.UP: "up",
		api.DOWN: "down",
		api.LEFT: "left",
		api.RIGHT: "right",
		api.UNKNOWN: "WFT!",
	}
	decisionFilters = []filters.DecisionFilter{
		filters.IsUnknownFilter{},
		filters.IsSolidFilter{},
	}

)

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	sr := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &sr)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	for _, r := range responsibilities {
		choice := r.Decision(&sr)
		okChoice := true
		for _, filter := range decisionFilters {
			ok, _ := filter.Allowed(choice, &sr)
			if !ok {
				fmt.Printf("%s %s", directionStrings[choice], filter.Description())
				okChoice = false
				break
			}
		}
		if !okChoice {
			break
		}
		fmt.Sprintf("%s\n", directionStrings[choice])
		respond(res, api.MoveResponse{
			Move: directionStrings[choice],
			Taunt: r.Taunt(),
		})
		return
	}

	respond(res, api.MoveResponse{Move: "down"})
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}
