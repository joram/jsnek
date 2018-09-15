package logic

import "github.com/joram/jsnek/api"

type OnlyOneChoice struct{}

func (ge OnlyOneChoice) Taunt() string {
	return "ONLY ONE OPTION"
}

func (ooc OnlyOneChoice) Decision(sr *api.SnakeRequest) int {
	// if there is only one choice, return that choice
	choices := sr.MyEmptyAdjacents()
	if len(choices) == 1 {
		return sr.You.Head().DirectionTo(choices[0])
	}
	return api.UNKNOWN
}
