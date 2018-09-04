package main

type GoEatOrthogonal struct {
	hungryHealth int
}

func (ge GoEatOrthogonal) taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) decision(sr *SnakeRequest) int {
	if sr.You.Health > ge.hungryHealth {
		return UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.Head())
	if err != nil {
		return UNKNOWN
	}

	d := sr.You.Head().NearestDirectionTo(*closestFood)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return UNKNOWN
}
