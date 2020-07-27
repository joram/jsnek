package filters

import (
	"fmt"
	"github.com/joram/jsnek/api"
)

type IsSmallSpace struct {
	accessibleCoords map[int]api.Coord
}

func (iss IsSmallSpace) Description() string { return "is small space" }

func (iss IsSmallSpace) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {

	//emptySurroundingCoords := 0
	//for _, a := range sr.You.Head().SurroundingCoords() {
	//	if sr.Board.IsEmpty(a){
	//		emptySurroundingCoords += 1
	//	}
	//}
	//if emptySurroundingCoords >= 7 {
	//	fmt.Println("skipping small space check")
	//	return true, nil
	//}

	coords := map[int]int{
		api.UP: len(accessibleCoords(sr.You.Head().Up(), []api.Coord{}, sr)),
		api.DOWN: len(accessibleCoords(sr.You.Head().Down(), []api.Coord{}, sr)),
		api.LEFT: len(accessibleCoords(sr.You.Head().Left(), []api.Coord{}, sr)),
		api.RIGHT: len(accessibleCoords(sr.You.Head().Right(), []api.Coord{}, sr)),
	}

	sum := 0
	goodPaths := 0
	for _, c := range coords {
		if c > 0 {
			sum += c
			goodPaths += 1
		}
	}
	avg := sum/goodPaths
	fmt.Printf("up: %d, down:%d, left:%d, right:%d, avg:%d\n", coords[api.UP], coords[api.DOWN], coords[api.LEFT], coords[api.RIGHT], avg)
	if avg - coords[direction] > 10 {
		fmt.Println("is too small")
		return false, nil
	}
	fmt.Println("is big enough")
	return true, nil
}


func accessibleCoords(from api.Coord, alreadyVisitedCoords []api.Coord, sr *api.SnakeRequest) []api.Coord {
	visited := map[api.Coord]bool{}

	// solid
	if !sr.Board.IsEmpty(from) {
		return []api.Coord{}
	}

	toVisit := []api.Coord{from}
	for {

		// done
		if len(toVisit) == 0 {
			break
		}

		// pop
		c := toVisit[0]
		toVisit = toVisit[1:]

		if visited[c] || !sr.Board.IsEmpty(c) {
			continue
		}

		visited[c] = true
		for _, next := range c.Adjacent(){
			if !visited[next] && sr.Board.IsEmpty(next) {
				toVisit = append(toVisit, next)
			}
		}
	}

	accessible := []api.Coord{}
	for c, b := range visited {
		if b {
			accessible = append(accessible, c)
		}
	}
	return accessible

}

func accessibleCoordsOld(from api.Coord, alreadyVisitedCoords []api.Coord, sr *api.SnakeRequest) []api.Coord {

	// solid
	if !sr.Board.IsEmpty(from) {
		return []api.Coord{}
	}

	// already visited
	for _, visitedCoord := range alreadyVisitedCoords {
		if from.Equal(visitedCoord) {
			return []api.Coord{}
		}
	}

	nextAlreadyVisitedCoords := append(alreadyVisitedCoords, from)
	alreadyVisitedCoords = accessibleCoords(from.Up(), nextAlreadyVisitedCoords, sr)
	alreadyVisitedCoords = accessibleCoords(from.Down(), nextAlreadyVisitedCoords, sr)
	alreadyVisitedCoords = accessibleCoords(from.Left(), nextAlreadyVisitedCoords, sr)
	alreadyVisitedCoords = accessibleCoords(from.Right(), nextAlreadyVisitedCoords, sr)
	alreadyVisitedCoords = unique(alreadyVisitedCoords)
	return alreadyVisitedCoords
}

func unique(intSlice []api.Coord) []api.Coord {
	keys := make(map[api.Coord]bool)
	var list []api.Coord
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
