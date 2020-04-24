package main

import (
	"fmt"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/filters"
	"github.com/joram/jsnek/logic"
	"sort"
)

var (
	logics = map[logic.Responsibility]int{
		logic.OnlyOneChoice{}:                      1000,
		logic.AvoidHeadOnHead{}:                    10,
		logic.AvoidThreatened{}:                    10,
		logic.GoEatOrthogonal{HungryHealth: 25}:    1,
		logic.ShortestSnake{LengthCompensation: 3}: 1,
		logic.KillOnlyOneChoice{}:                  1,
		logic.GoMoreRoom{Ratio: 3}:                 20,
		logic.TrapFood{}:                           1,
		logic.ValidDirection{}:                     1,
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
	return move_weighted(request)
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
	return "up"
}

func move_sequential_check(request api.SnakeRequest) string {

	for l, _ := range logics {
		choice := l.Decision(&request)
		if choice == api.UNKNOWN {
			continue
		}
		if !isGoodDecision(choice, request) {
			println("skipping choice " + directionStrings[choice] + " by " + l.Taunt())
			continue
		}
		fmt.Println(request.Game.ID, l.Taunt())
		return directionStrings[choice]
	}
	return "up"
}