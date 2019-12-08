package logic

import "github.com/joram/jsnek/api"

type GoEatOrthogonal struct {
	HungryHealth int
}

func (ge GoEatOrthogonal) Taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) Decision(sr *api.SnakeRequest) int {
	if sr.You.Health > ge.HungryHealth {
		return api.UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.Head())
	if err != nil {
		return api.UNKNOWN
	}

	d := sr.You.Head().NearestDirectionTo(*closestFood)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return api.UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return api.UNKNOWN
}
