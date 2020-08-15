package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/filters"
	"github.com/joram/jsnek/logic"
)

var (
	orderedLogics = []logic.Responsibility{
		logic.Topology{Weight: map[api.Coord]float64{} },
		logic.OnlyOneChoice{},
		logic.AvoidHeadOnHead{},
		logic.AvoidThreatened{},
		logic.GoEatOrthogonal{IgnoreHazardFood: false, HungryHealth: 40},
		logic.EatWhenShortestSnake{IgnoreHazardFood: false, LengthCompensation: 3},
		logic.GuaranteedKill{},
		//logic.TrapFood{},
		logic.GoMoreRoom{Ratio: 3},
		logic.GoToClosestTail{},
		logic.GoEatOrthogonal{IgnoreHazardFood: true, HungryHealth: 40},
		logic.EatWhenShortestSnake{IgnoreHazardFood: true, LengthCompensation: 3},
		logic.AvoidOthers{},
		logic.TrapFood{},
		logic.ValidDirection{},
	}
	directionStrings = map[int]string{
		api.UP:      "up",
		api.DOWN:    "down",
		api.LEFT:    "left",
		api.RIGHT:   "right",
		api.UNKNOWN: "WFT!",
	}
)

func isGoodDecision(choice int, request api.SnakeRequest, filters []filters.DecisionFilter) bool {
	for _, filter := range filters {
		ok, _ := filter.Allowed(choice, &request)
		if !ok {
			return false
		}
	}
	return true
}


func move(request api.SnakeRequest) string {
	unknown := directionStrings[api.UNKNOWN]

	if len(request.OtherSnakes()) == 0 {
		return moveSingleplayer(request)
	}

	possibleFilters := [][]filters.DecisionFilter{
		{
			filters.IsUnknownFilter{},
			filters.IsSolidFilter{},
			filters.IsThreatenedFilter{},
			filters.IsHazardFilter{},
			filters.IsSmallSpace{},
		},
		{
			filters.IsUnknownFilter{},
			filters.IsThreatenedFilter{},
			filters.IsHazardFilter{},
			filters.IsSmallSpace{},
		},
		{
			filters.IsUnknownFilter{},
			filters.IsThreatenedFilter{},
			filters.IsHazardFilter{},
		},
		{
			filters.IsUnknownFilter{},
			filters.IsThreatenedFilter{},
		},
		{
			filters.IsUnknownFilter{},
		},
		{},
	}
	for i, fs := range possibleFilters {
		s, reason := attemptMove(request, orderedLogics, fs)
		if s != unknown {
			fmt.Printf("[%d]%s\t%s\n", i, s, reason)
			return s
		}
	}


	s := moveRandomEmpty(request)
	fmt.Printf("yolo: %s\n", s)
	return s
}

func moveRandomEmpty(request api.SnakeRequest) string {
	for s, c := range request.You.GetHead().AdjacentMap() {
		if request.Board.IsEmpty(c) {
			return s
		}
	}
	return directionStrings[api.UNKNOWN]
}


func attemptMove(request api.SnakeRequest, logics []logic.Responsibility, filters []filters.DecisionFilter) (string, string) {
	for _, l := range logics {
		choice := l.Decision(&request)
		if choice == api.UNKNOWN {
			//fmt.Printf("not doing %s, unknown\n", l.Taunt())
			continue
		}
		if !isGoodDecision(choice, request, filters) {
			//fmt.Printf("not doing %s, bad decision\n", l.Taunt())
			continue
		}
		return directionStrings[choice], l.Taunt()
	}
	return directionStrings[api.UNKNOWN], "unknown"

}


func moveSingleplayer(request api.SnakeRequest) string {
	head := request.You.GetHead()


	if head.Y == request.Board.Height - 1 {

		if head.X == 0 {
			return "up"
		}
		return "left"
	}

	// zig zag down the right side
	l := request.Board.Width - 2
	r := request.Board.Width - 1
	if head.X == l || head.X == r {
		t := head.X + head.Y
		if t % 2 == 0 {
			return "down"
		}
		if head.X == r {
			return "left"
		}
		return "right"
	}

	if head.Y == 0 && head.X % 2 == 0 {
		return "right"
	}
	if head.Y == request.Board.Height - 2 && head.X % 2 == 1 {
		return "right"
	}
	if head.X % 2 == 0 {
		return "up"
	}
	return "down"
}