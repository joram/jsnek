package api

import (
	"errors"
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

	var occupied []Coord
	for _, snake := range b.Snakes {
		for _, coord := range snake.Body {
			if coord.Equal(snake.Tail()) && !snake.JustAte() {
				continue
			}
			occupied = append(occupied, coord)
		}
	}

	for _, o := range occupied {
		if o.Equal(c) {
			return false
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

func (b *Board) PopulateDistances(you Snake) {
	b.Data = map[string]*DistanceData{}
	b.AbleToVisitCount = map[string]int{
		"left":  0,
		"right": 0,
		"up":    0,
		"down":  0,
	}

	left := you.Head().Left()
	if b.IsEmpty(left) {
		meLeft := &DistanceData{}
		meLeft.Calculate([]Coord{left}, b)
		b.Data["me_left"] = meLeft
		b.AbleToVisitCount["left"] = meLeft.Count
	}

	right := you.Head().Right()
	if b.IsEmpty(right) {
		meRight := &DistanceData{}
		meRight.Calculate([]Coord{right}, b)
		b.Data["me_right"] = meRight
		b.AbleToVisitCount["right"] = meRight.Count
	}

	up := you.Head().Up()
	if b.IsEmpty(up) {
		meUp := &DistanceData{}
		meUp.Calculate([]Coord{up}, b)
		b.Data["me_up"] = meUp
		b.AbleToVisitCount["up"] = meUp.Count
	}

	down := you.Head().Down()
	if b.IsEmpty(down) {
		meDown := &DistanceData{}
		meDown.Calculate([]Coord{down}, b)
		b.Data["me_down"] = meDown
		b.AbleToVisitCount["down"] = meDown.Count
	}

	for _, snake := range b.Snakes {
		b.Data[snake.ID] = &DistanceData{}
		b.Data[snake.ID].Calculate(snake.Head().Adjacent(), b)
	}
}
