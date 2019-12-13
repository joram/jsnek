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
