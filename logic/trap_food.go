package logic

import "github.com/joram/jsnek/api"

import "math/rand"

type TrapFood struct {
	hungryHealth int
}

func (ge TrapFood) Taunt() string {
	return "ITS A TARP!"
}

func (ge TrapFood) isHoneyPotted(sr *api.SnakeRequest, food api.Coord) bool {
	coords := food.SurroundingCoords()
	toFill := []api.Coord{}
	for _, coord := range coords {
		if sr.Board.IsEmpty(coord) {
			toFill = append(toFill, coord)
		}
	}
	return len(toFill) <= 2
}

func (ge TrapFood) Decision(sr *api.SnakeRequest) int {
	orderedFood := sr.Board.OrderedClosestFood(sr.You.Head())

	for _, food := range orderedFood {
		if ge.isHoneyPotted(sr, food){
			continue
		}
		coords := food.SurroundingCoords()
		toFill := []api.Coord{}
		foundClosestToFill := false
		closestToFill := api.Coord{}
		closestDist := float64(-1)
		for _, coord := range coords {
			if sr.Board.IsEmpty(coord) {
				dist := sr.You.Head().OrthogonalDistance(coord)
				if !foundClosestToFill || dist <= closestDist {
					if sr.Board.IsEmpty(coord) {
						closestToFill = coord
						closestDist = dist
						foundClosestToFill = true
					}
				}
				toFill = append(toFill, coord)
			}
		}

		if len(toFill) > 2 {
			return sr.You.Head().DirectionTo(closestToFill)
		}

		validChoices := []api.Coord{}
		for _, coord := range sr.You.Head().Adjacent() {
			toAvoid := false
			for _, tfCoord := range toFill {
				if tfCoord.Equal(coord) {
					toAvoid = true
					break
				}
			}
			if !toAvoid {
				validChoices = append(validChoices, coord)
			}
		}
		if len(validChoices) >= 1 {
			i := rand.Intn(len(validChoices))
			nextCoord := validChoices[i]
			return sr.You.Head().DirectionTo(nextCoord)
		}
	}
	return api.UNKNOWN
}
