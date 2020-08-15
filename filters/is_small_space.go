package filters

import (
	"github.com/joram/jsnek/api"
)

type IsSmallSpace struct {
	accessibleCoords map[int]api.Coord
}

func (iss IsSmallSpace) Description() string { return "is small space" }

func (iss IsSmallSpace) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {

	emptySurroundingCoords := 0
	for _, a := range sr.You.GetHead().SurroundingCoords() {
		if sr.Board.IsEmpty(a){
			emptySurroundingCoords += 1
		}
	}
	if emptySurroundingCoords >= 7 {
		return true, nil
	}

	coords := map[int]int{
		api.UP: len(accessibleCoords(sr.You.GetHead().Up(), []api.Coord{}, sr)),
		api.DOWN: len(accessibleCoords(sr.You.GetHead().Down(), []api.Coord{}, sr)),
		api.LEFT: len(accessibleCoords(sr.You.GetHead().Left(), []api.Coord{}, sr)),
		api.RIGHT: len(accessibleCoords(sr.You.GetHead().Right(), []api.Coord{}, sr)),
	}

	sum := 0
	goodPaths := 0
	for _, c := range coords {
		if c > 0 {
			sum += c
			goodPaths += 1
		}
	}
	if goodPaths == 0 {
		goodPaths = 1
	}
	avg := sum/goodPaths
	//fmt.Printf("up: %d, down:%d, left:%d, right:%d, avg:%d\n", coords[api.UP], coords[api.DOWN], coords[api.LEFT], coords[api.RIGHT], avg)
	if float64(coords[direction])/float64(avg) < float64(0.8) {
		return false, nil
	}
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