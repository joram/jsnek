package logic

import "github.com/joram/jsnek/api"

type GoEatOrthogonal struct {
	IgnoreHazardFood bool
	HungryHealth int
}

func (ge GoEatOrthogonal) Taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) Decision(sr *api.SnakeRequest) int {
	if sr.You.Health > ge.HungryHealth {
		return api.UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.GetHead(), ge.IgnoreHazardFood)
	if err != nil {
		return api.UNKNOWN
	}

	d := sr.You.GetHead().NearestDirectionTo(*closestFood)
	nextCoord, err := sr.You.GetHead().Offset(d)
	if err != nil {
		return api.UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return api.UNKNOWN
}
