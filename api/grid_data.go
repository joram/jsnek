package api

import "errors"

type DistanceData struct {
	InitialCoords []Coord
	Data map[int]map[int]int `json:"data"`
	Count int `json:"count"`
	board *Board
	calculated bool
}

func (gd *DistanceData) Calculate(initalCoords []Coord, board *Board) {
	if gd.calculated {
		return
	}
	gd.InitialCoords = initalCoords
	gd.board = board
	gd.Count = 0

	type toVisitCoord struct {
		coord Coord
		value int
	}
	toVisit := []toVisitCoord{}
	for _, c := range initalCoords {
		toVisit = append(toVisit, toVisitCoord{c, 1})
	}

	for true {

		// break
		if len(toVisit) == 0 {
			gd.calculated = true
			return
		}

		// pop
		tvc := toVisit[0]
		toVisit = toVisit[1:]

		// skip
		if gd.HasData(tvc.coord){
			continue
		}

		// visit
		gd.AddData(tvc.coord, tvc.value)
		gd.Count += 1

		// iterate
		for _, coord := range tvc.coord.Adjacent(){
			if !gd.HasData(coord){
				toVisit = append(toVisit, toVisitCoord{coord, tvc.value+1})
			}
		}
	}

}

func (b *DistanceData) AddData(c Coord, val int){
	if b.Data == nil {
		b.Data = map[int]map[int]int{}
	}

	_, exists := b.Data[c.X]
	if !exists {
		b.Data[c.X] = map[int]int{}
	}

	b.Data[c.X][c.Y] = val
}

func (b *DistanceData) HasData(c Coord) bool {
	if b.Data == nil {
		return false
	}
	_, exists := b.Data[c.X]
	if !exists {
		return false
	}
	_, exists = b.Data[c.X][c.Y]
	return exists
}

func (b *DistanceData) GetData(c Coord) (int, error){
	err := errors.New("nothing at coord")
	if b.Data == nil {
		return 0, err
	}
	_, exists := b.Data[c.X]
	if !exists {
		return 0, err
	}

	val, exists := b.Data[c.X][c.Y]
	if !exists {
		return 0, err
	}
	return val, nil
}


