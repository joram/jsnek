package main

import "math/rand"

type TrapFood struct {
	hungryHealth int
}

func (ge TrapFood) taunt() string {
	return "ITS A TARP!"
}

func (ge TrapFood) isHoneyPotted(sr *SnakeRequest, food Coord) bool {
	coords := food.SurroundingCoords()
	toFill := []Coord{}
	for _, coord := range coords {
		if sr.Board.IsEmpty(coord) {
			toFill = append(toFill, coord)
		}
	}
	return len(toFill) <= 2
}

func (ge TrapFood) decision(sr *SnakeRequest) int {
	orderedFood := sr.Board.OrderedClosestFood(sr.You.Head())

	for _, food := range orderedFood {
		if ge.isHoneyPotted(sr, food){
			continue
		}
		coords := food.SurroundingCoords()
		toFill := []Coord{}
		foundClosestToFill := false
		closestToFill := Coord{}
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

		validChoices := []Coord{}
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
	return UNKNOWN
}
