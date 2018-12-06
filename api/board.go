package api

import (
	"errors"
	"fmt"
)

func (b *Board) IsEmpty(c Coord) bool {
	if c.X >= b.Width {
		return false
	}
	if c.Y >= b.Height {
		return false
	}
	if c.X < 0 {
		return false
	}
	if c.Y < 0 {
		return false
	}

	for _, snake := range b.Snakes {
		for _, coord := range snake.Body {
			if coord.Equal(c) {
				return false
			}
		}
	}
	return true
}

func (b *Board) AddData(c Coord, key string, val int){
	if b.Data == nil {
		b.Data = map[string]map[int]map[int]int{}
	}
	_, exists := b.Data[key]
	if !exists {
		b.Data[key] = map[int]map[int]int{}
	}
	_, exists = b.Data[key][c.X]
	if !exists {
		b.Data[key][c.X] = map[int]int{}
	}

	b.Data[key][c.X][c.Y] = val
}

func (b *Board) GetData(c Coord, key string) (int, error){
	err := errors.New("nothing at coord")
	if b.Data == nil {
		return 0, err
	}
	_, exists := b.Data[key]
	if !exists {
		return 0, err
	}
	_, exists = b.Data[key][c.X]
	if !exists {
		return 0, err
	}

	val, exists := b.Data[key][c.X][c.Y]
	if !exists {
		return 0, err
	}
	return val, nil
}

func (b *Board) ClosestFood(c Coord) (*Coord, error) {
	foundFood := false
	closestFood := Coord{}
	closestDist := float64(-1)

	for _, food := range b.Food {
		dist := c.OrthogonalDistance(food)
		if !foundFood || dist < closestDist {
			closestFood = food
			closestDist = dist
			foundFood = true
		}
	}

	if !foundFood {
		return nil, errors.New("no food")
	}
	return &closestFood, nil
}

func closest(c Coord, coords []Coord) Coord {
	foundFood := false
	closestFood := Coord{}
	closestDist := float64(-1)
	for _, food := range coords {
		dist := c.OrthogonalDistance(food)
		if !foundFood || dist < closestDist {
			closestFood = food
			closestDist = dist
			foundFood = true
		}
	}
	return closestFood
}

func (b *Board) OrderedClosestFood(c Coord) []Coord {
	toSort := b.Food
	sorted := []Coord{}
	for len(toSort) > 0 {
		nextClosestFood := closest(c, toSort)
		sorted = append(sorted, nextClosestFood)
		newToSort := []Coord{}
		for _, ts := range toSort {
			if !ts.Equal(nextClosestFood) {
				newToSort = append(newToSort, ts)
			}
		}
		toSort = newToSort
	}
	return sorted
}

func (b *Board) GetTimeTo(c Coord, snake_id string) (int, error) {
	key := fmt.Sprintf("time_to_%s", snake_id)
	_, err := b.GetData(c, key)
	if err != nil {
		b.PopulateDistances()
	}
	return b.GetData(c, key)
}


func (b* Board) PopulateDistances(){
	type DistToCoord struct {
		coord Coord
		distance int
		snake_id string
	}
	edge := []DistToCoord{}
	for _, snake := range b.Snakes {
		for _, coord := range snake.Head().Adjacent() {
			if b.IsEmpty(coord){
				edge = append(edge, DistToCoord{coord, 1, snake.ID,})
			}
		}
	}

	for true {
		if len(edge) == 0 {
			break
		}
		dtc := edge[0]
		edge = edge[1:]
		key := fmt.Sprintf("time_to_%s", dtc.snake_id)
		b.AddData(dtc.coord, key, dtc.distance)


		//// get closest dist
		//minDist := dtc.distance
		//val, err := b.GetData(dtc.coord, key)
		//if err == nil && val < minDist {
		//	continue
		//}

		// delve further
		for _, adjCoord := range dtc.coord.Adjacent() {
			if !b.IsEmpty(adjCoord) {
				continue
			}
			_, err := b.GetData(adjCoord, key)
			if err == nil {
				continue
			}

			edge = append(edge, DistToCoord{adjCoord, dtc.distance+1, dtc.snake_id,})

		}
	}
}