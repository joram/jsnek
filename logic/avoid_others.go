package logic

import "github.com/joram/jsnek/api"

type AvoidOthers struct {
	HungryHealth int
}

func (ao AvoidOthers) Taunt() string {
	return "Avoid Others"
}

func (ao AvoidOthers) Decision(sr *api.SnakeRequest) int {
	if sr.You.Health > ao.HungryHealth {
		return api.UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.GetHead(), false)
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
