package logic

import "github.com/joram/jsnek/api"

type GuaranteedKill struct {
}

func (kooc GuaranteedKill) Taunt() string {
	return "Killing Other Snakes with only one choice"
}

func (kooc GuaranteedKill) Decision(sr *api.SnakeRequest) int {

	for _, snake := range sr.OtherSnakes() {
		if len(snake.Body) >= len(sr.You.Body) {
			continue
		}

		var potentialCoordinates []api.Coord
		for _, c := range snake.GetHead().Adjacent() {
			if sr.Board.IsEmpty(c) {
				potentialCoordinates = append(potentialCoordinates, c)
			}
		}

		if len(potentialCoordinates) == 1 {
			for _, c := range sr.You.GetHead().Adjacent() {
				if c.IsAdjacent(potentialCoordinates[0]) {
					return c.DirectionTo(potentialCoordinates[0])
				}
			}
		}

	}
	return api.UNKNOWN
}
