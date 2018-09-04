package main

import (
	"math/rand"
)

type ValidDirection struct {

}

func (ge ValidDirection) taunt() string {
	return "ANY VALID DIRECTION"
}


func (ec ValidDirection) decision(sr *SnakeRequest) int {
	choices := sr.MyEmptyAdjacents()
	i := rand.Intn(len(choices))
	nextCoord := choices[i]
	return sr.You.Head().DirectionTo(nextCoord)
}
