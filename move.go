package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/filters"
	"github.com/joram/jsnek/logic"
	"sort"
)

var (
	orderedLogics = []logic.Responsibility{
		logic.OnlyOneChoice{},
		logic.AvoidHeadOnHead{},
		logic.AvoidThreatened{},
		logic.GoEatOrthogonal{HungryHealth: 25},
		logic.ShortestSnake{LengthCompensation: 3},
		logic.GoToClosestTail{},
		logic.KillOnlyOneChoice{},
		logic.TrapFood{},
		logic.GoMoreRoom{Ratio: 3},
		logic.ShortestSnake{},
		logic.AvoidOthers{},
		logic.TrapFood{},
		logic.ValidDirection{},
	}
	logics = map[logic.Responsibility]int{
		logic.OnlyOneChoice{}:                      1000,
		logic.AvoidHeadOnHead{}:                    10,
		logic.AvoidThreatened{}:                    10,
		logic.GoEatOrthogonal{HungryHealth: 25}:    1,
		logic.ShortestSnake{LengthCompensation: 3}: 1,
		logic.GoToClosestTail{}: 					10,
		logic.KillOnlyOneChoice{}:                  1,
		logic.GoMoreRoom{Ratio: 3}:                 20,
		logic.TrapFood{}:                           1,
		logic.ValidDirection{}:                     1,
		logic.ShortestSnake{}: 10,
		logic.AvoidOthers{}: 10,
		logic.TrapFood{}:50,
		// EAT THEIR LUNCH (force them to starve)
	}
	directionStrings = map[int]string{
		api.UP:      "up",
		api.DOWN:    "down",
		api.LEFT:    "left",
		api.RIGHT:   "right",
		api.UNKNOWN: "WFT!",
	}
	decisionFilters = []filters.DecisionFilter{
		filters.IsUnknownFilter{},
		filters.IsSolidFilter{},
		filters.IsThreatenedFilter{},
		filters.IsHazardFilter{},
		filters.IsSmallSpace{},
	}
)

func isGoodDecision(choice int, request api.SnakeRequest) bool {
	for _, filter := range decisionFilters {
		ok, _ := filter.Allowed(choice, &request)
		if !ok {
			return false
		}
	}
	return true
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}


func move(request api.SnakeRequest) string {
	unknown := directionStrings[api.UNKNOWN]

	if len(request.OtherSnakes()) == 0 {
		return move_singleplayer(request)
	}

	s := move_weighted(request)
	if s != unknown {
		fmt.Printf("weighted: %s\n", s)
		return s
	}

	s = move_safe_sequential_check(request)
	if s != unknown {
		fmt.Printf("safe: %s\n", s)
		return s
	}

	s = move_unsafe_sequential_check(request)
	if s != unknown {
		fmt.Printf("unsafe: %s\n", s)
		return s
	}

	s = move_random_empty(request)
	fmt.Printf("yolo: %s\n", s)
	return s
}

func move_random_empty(request api.SnakeRequest) string {
	for s, c := range request.You.Head().AdjacentMap() {
		if request.Board.IsEmpty(c) {
			return s
		}
	}
	return directionStrings[api.UNKNOWN]
}

func move_weighted(request api.SnakeRequest) string {
	directions := map[string]int{
		directionStrings[api.UP]:    0,
		directionStrings[api.DOWN]:  0,
		directionStrings[api.LEFT]:  0,
		directionStrings[api.RIGHT]: 0,
	}

	for l, weight := range logics {
		choice := l.Decision(&request)
		direction := directionStrings[choice]
		directions[direction] += weight
	}

	weights := []int{}
	weightMap := map[int]string{}
	for k, v := range directions {
		weightMap[v] = k
		weights = append(weights, v)
	}
	sort.Ints(weights)
	reverseInts(weights)

	for _, weight := range weights {
		direction := weightMap[weight]
		directionInt := api.StringToDir(direction)
		if !isGoodDecision(directionInt, request) {
			continue
		}
		return direction
	}
	return directionStrings[api.UNKNOWN]
}

func move_safe_sequential_check(request api.SnakeRequest) string {
	for _, l := range orderedLogics {
		choice := l.Decision(&request)
		if choice == api.UNKNOWN {
			continue
		}
		if !isGoodDecision(choice, request) {
			continue
		}
		return directionStrings[choice]
	}
	return directionStrings[api.UNKNOWN]

}

func move_unsafe_sequential_check(request api.SnakeRequest) string {
	for _, l := range orderedLogics {
		choice := l.Decision(&request)
		if choice == api.UNKNOWN {
			continue
		}
		return directionStrings[choice]
	}
	return "up"
}

func move_singleplayer(request api.SnakeRequest) string {
	head := request.You.Head()


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