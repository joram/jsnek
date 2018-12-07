package logic

import (
	"github.com/joram/jsnek/api"
)

type GoMoreRoom struct {
	Ratio float32
}

func (ge GoMoreRoom) Taunt() string {
	return "More Room"
}

func (ge GoMoreRoom) Decision(sr *api.SnakeRequest) int {
	sr.Board.PopulateDistances()

	biggest := -1
	biggestDir := ""
	secondBiggest := -1
	for dir, count := range sr.Board.AbleToVisitCount {
		if count > biggest {
			secondBiggest = biggest
			biggest = count
			biggestDir = dir
		}
	}

	if float32(secondBiggest)/float32(biggest) < ge.Ratio {
		return map[string]int{
			"up": api.UP,
			"down": api.DOWN,
			"left": api.LEFT,
			"right":api.RIGHT,
			"WFT!": api.UNKNOWN,
		}[biggestDir]
	}

	return api.UNKNOWN
}
