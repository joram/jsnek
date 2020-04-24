package logic

import "github.com/joram/jsnek/api"

type GoToClosestTail struct {
	HungryHealth int
}

func (ge GoToClosestTail) Taunt() string {
	return "EATING"
}

func (ge GoToClosestTail) Decision(sr *api.SnakeRequest) int {
	closestTail := sr.Board.ClosestTail(sr.You.Head())

	d := sr.You.Head().NearestDirectionTo(closestTail)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return api.UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return api.UNKNOWN
}
