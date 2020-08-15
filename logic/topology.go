package logic

import (
	"github.com/joram/jsnek/api"
	"math"
)

type Topology struct {
	Weight map[api.Coord]float64
}
type ToVisit struct {
	coord api.Coord
	prevValue float64
}

func (t Topology) Taunt() string {
	return "topology"
}

type visitFuncDef func (sr *api.SnakeRequest, curr *api.Coord, prevValue float64) float64

func (t Topology) Traverse(initialVal float64, sr *api.SnakeRequest, initialCoords []api.Coord, visitFunc visitFuncDef) {
	for _, coord := range initialCoords {
		visitFunc(sr, &coord, initialVal)
	}
	toVisit := []ToVisit{}
	for _, coord := range initialCoords {
		toVisit = append(toVisit, ToVisit{coord, initialVal})
	}
	haveVisited := map[api.Coord]bool{}
	for {

		// prep
		if len(toVisit) == 0 {
			return
		}
		visiting := toVisit[0]
		toVisit = toVisit[1:]
		if sr.Board.IsSolid(visiting.coord){
			continue
		}

		// traverse neighbours
		for _, neighbour := range visiting.coord.Adjacent() {
			if sr.Board.IsEmpty(neighbour) && !haveVisited[neighbour]{
				haveVisited[neighbour] = true
				currVal := visitFunc(sr, &neighbour, visiting.prevValue)
				toVisit = append(toVisit, ToVisit{neighbour, currVal})
			}
		}

	}
}

func (t Topology) VisitFoodDistance(sr *api.SnakeRequest, curr *api.Coord, prevVal float64) float64 {
	delta := math.Abs(101.-float64(sr.You.Health)/10.)
	t.Weight[*curr] +=  prevVal - delta
	return t.Weight[*curr]
}

func (t Topology) VisitOpponentDistance(sr *api.SnakeRequest, curr *api.Coord, prevVal float64) float64 {
	t.Weight[*curr] +=  prevVal+10
	return t.Weight[*curr]
}


func (t Topology) Decision(sr *api.SnakeRequest) int {
	t.Weight = map[api.Coord]float64{}
	t.Traverse(100, sr, sr.Board.Food, t.VisitFoodDistance)

	var opponentHeads []api.Coord
	for _, oppponent := range sr.OtherSnakes(){
		opponentHeads = append(opponentHeads, oppponent.GetHead())
	}
	t.Traverse(100, sr, opponentHeads, t.VisitOpponentDistance)

	// choose largest neighbour
	first := true
	maxWeight := -1000.
	maxWeightDirString := api.DirToString(api.UNKNOWN)
	for dirString, coord := range sr.You.GetHead().AdjacentMap() {
		if sr.Board.IsSolid(coord) {
			continue
		}
		//fmt.Printf("%s:%f\n",dirString, t.Weight[coord])
		if t.Weight[coord] > maxWeight || first {
			maxWeight = t.Weight[coord]
			maxWeightDirString = dirString
			first = false
		}
	}
	return api.StringToDir(maxWeightDirString)
}
