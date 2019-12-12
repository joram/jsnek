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
	sr.Board.PopulateDistances(sr.You)

	biggest := -1
	biggestDir := ""
	for dir, count := range sr.Board.AbleToVisitCount {
		if count > biggest {
			biggest = count
			biggestDir = dir
		}
	}
	return api.StringToDir(biggestDir)
}
