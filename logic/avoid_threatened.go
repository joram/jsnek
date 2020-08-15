package logic

import "github.com/joram/jsnek/api"

type AvoidThreatened struct {
	HungryHealth int
}

func (at AvoidThreatened) Taunt() string {
	return "Avoiding Threatened squares"
}

func (at AvoidThreatened) Decision(sr *api.SnakeRequest) int {
	var threatenedSquares []api.Coord
	for _, snake := range sr.OtherSnakes() {
		if len(snake.Body) < len(sr.You.Body) {
			continue
		}
		threatenedSquares = append(threatenedSquares, snake.GetHead().Adjacent()...)
	}

	var potentialCoordinates []api.Coord
	for _, potentialCoord := range sr.You.GetHead().Adjacent() {
		threatened := false
		for _, threatenedSquare := range threatenedSquares {
			if potentialCoord.Equal(threatenedSquare) {
				threatened = true
				break
			}
		}
		if threatened {
			continue
		}
		potentialCoordinates = append(potentialCoordinates, potentialCoord)
	}

	if len(potentialCoordinates) == 1 {
		return sr.You.GetHead().DirectionTo(potentialCoordinates[0])
	}

	return api.UNKNOWN
}
