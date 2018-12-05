package api

import "errors"

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
	if !b.timeToBuilt {
		type DistToCoord struct {
			coord Coord
			distance int
		}
		edge := []DistToCoord{}
		for _, snake := range b.Snakes {
			d := DistToCoord{snake.Head(), 0}
			edge = append(edge, d)
		}

		for true {
			dtc, edge := edge[0], edge[1:]

			if !b.IsEmpty(dtc.coord){
				continue
			}

			// map exists
			_, exists := b.timeTo[dtc.coord]
			if !exists {
				b.timeTo[dtc.coord] = map[string]int{}
			}

			shortestDist := dtc.distance
			oldDist, exists := b.timeTo[dtc.coord][snake_id]
			if exists && oldDist < shortestDist {
				shortestDist = oldDist
				continue
			}

			if len(edge) == 0 {
				break
			}
			for _, n := range dtc.coord.Adjacent() {
				if b.IsEmpty(n) {
					_, exists = b.timeTo[dtc.coord][snake_id]
					if !exists {
						next_d := DistToCoord{n, dtc.distance+1}
						edge = append(edge, next_d)
					}
				}
			}
		}

		b.timeToBuilt = true
	}

	times, exists := b.timeTo[c]
	if !exists {
		return -1, errors.New("no data at coord")
	}
	t, exists := times[snake_id]
	if !exists {
		return -1, errors.New("no data at snake at coord")
	}
	return t, nil
}