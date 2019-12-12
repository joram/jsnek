package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/filters"
	"github.com/joram/jsnek/logic"
)

var (
	logics = map[logic.Responsibility]int{
		logic.OnlyOneChoice{}: 1,
		logic.AvoidHeadOnHead{}: 1,
		logic.AvoidThreatened{}: 1,
		logic.GoEatOrthogonal{HungryHealth: 25}: 1,
		logic.ShortestSnake{LengthCompensation: 3}: 1,
		logic.KillOnlyOneChoice{}: 1,
		// EAT THEIR LUNCH (force them to starve)
		logic.GoMoreRoom{Ratio: 3}: 1,
		logic.TrapFood{}: 1,
		logic.ValidDirection{}: 1,
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

func move(request api.SnakeRequest) string {
	directions := map[string]int{
		directionStrings[api.UP]:    0,
		directionStrings[api.DOWN]:  0,
		directionStrings[api.LEFT]:  0,
		directionStrings[api.RIGHT]: 0,
	}

	for logic, weight := range logics {
		choice := logic.Decision(&request)
		direction := directionStrings[choice]
		fmt.Printf("Logic says: %s\t%d\n", direction, weight)
		directions[direction] += weight
	}

	bestDirection := directionStrings[api.UP]
	bestWeight := -1
	for direction, weight := range directions {
		if weight > bestWeight && direction != directionStrings[api.UNKNOWN] {
			bestWeight = weight
			bestDirection = direction
		}
		fmt.Printf("Weight says: %s\t%d\n", direction, weight)
	}
	return bestDirection

		//okChoice := true
		//for _, filter := range decisionFilters {
		//	ok, _ := filter.Allowed(choice, &sr)
		//	if !ok {
		//		okChoice = false
		//		break
		//	}
		//}
		//if choice == api.UNKNOWN {
		//	continue
		//}
		//if !okChoice {
		//	println("skipping choice "+directionStrings[choice]+" by "+l.Taunt())
		//	continue
		//}
		//fmt.Println(sr.Game.ID, l.Taunt())
		//respond(res, api.MoveResponse{
		//	Move:  directionStrings[choice],
		//	Taunt: l.Taunt(),
		//})
		//return

}
