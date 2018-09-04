package main

type OnlyOneChoice struct {

}

func (ooc OnlyOneChoice) decision(sr *SnakeRequest) int {
	// if there is only one choice, return that choice
	choices := []Coord{}
	for _, a := range sr.You.Head().Adjacent() {
		if sr.Board.IsEmpty(a) {
			choices = append(choices, a)
		}
	}
	if len(choices) == 1 {
		return sr.You.Head().DirectionTo(choices[0])
	}
	return UNKNOWN
}



