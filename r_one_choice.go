package main

type OnlyOneChoice struct {

}

func (ooc OnlyOneChoice) decision(sr *SnakeRequest) int {
	// if there is only one choice, return that choice
	choices := sr.MyEmptyAdjacents()
	if len(choices) == 1 {
		return sr.You.Head().DirectionTo(choices[0])
	}
	return UNKNOWN
}



