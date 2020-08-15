package logic

import "github.com/joram/jsnek/api"

type GoToClosestTail struct {
}

func (ge GoToClosestTail) Taunt() string {
	return "EATING"
}

func (ge GoToClosestTail) Decision(sr *api.SnakeRequest) int {
	closestTail := sr.Board.ClosestTail(sr.You.GetHead())

	d := sr.You.GetHead().NearestDirectionTo(closestTail)
	nextCoord, err := sr.You.GetHead().Offset(d)
	if err != nil {
		return api.UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return api.UNKNOWN
}
