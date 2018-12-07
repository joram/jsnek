package logic

import "github.com/joram/jsnek/api"

import (
	"math/rand"
)

type ValidDirection struct {
}

func (ge ValidDirection) Taunt() string {
	return "ANY VALID DIRECTION"
}

func (ec ValidDirection) Decision(sr *api.SnakeRequest) int {
	choices := sr.MyEmptyAdjacents()
	if len(choices) == 0 {
		return api.UNKNOWN
	}
	i := rand.Intn(len(choices))
	nextCoord := choices[i]
	return sr.You.Head().DirectionTo(nextCoord)
}
