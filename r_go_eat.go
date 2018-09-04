package main

type GoEatOrthogonal struct {
	hungryHealth int
}

func (ge GoEatOrthogonal) Taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) decision(sr *SnakeRequest) int {
	if sr.You.Health > ge.hungryHealth {
		return UNKNOWN
	}
	foundFood := false
	closestFood := Coord{}
	closestDist := float64(-1)
	for _, food := range sr.Board.Food {
		dist := sr.You.Head().OrthogonalDistance(food)
		if !foundFood || dist < closestDist {
			closestFood = food
			closestDist = dist
			foundFood = true
		}
	}

	if !foundFood {
		return UNKNOWN
	}

	d := sr.You.Head().NearestDirectionTo(closestFood)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return UNKNOWN
}
