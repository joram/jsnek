package util

import "github.com/joram/jsnek/api"

type ToVisit struct {
	coord api.Coord
	weight int
}

func CreateHeatmap(startingCoords []api.Coord) map[api.Coord]int {
	heat := map[api.Coord]int{}
	var toVisit []ToVisit
	for _, coord := range startingCoords {
		tv := ToVisit{coord, 0}
		toVisit = append(toVisit, tv)
	}
	for {
		// done
		if len(toVisit) == 0 {
			break
		}

		// pop
		tv := toVisit[0]
		toVisit = toVisit[1:]

		// visit
		heat[tv.coord] = tv.weight

		// visit next
		for _, next := range tv.coord.Adjacent() {
			if _, ok := heat[next]; !ok {
				tvNext := ToVisit{next, tv.weight+1}
				toVisit = append(toVisit, tvNext)
			}
		}
	}
	return heat
}
