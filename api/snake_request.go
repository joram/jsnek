package api

func (sr SnakeRequest) MyEmptyAdjacents() []Coord {
	choices := []Coord{}
	for _, a := range sr.You.Head().Adjacent() {
		if sr.Board.IsEmpty(a) {
			choices = append(choices, a)
		}
	}
	return choices
}
