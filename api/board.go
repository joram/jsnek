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


func (b* Board) PopulateDistances(you Snake){

	meLeft := &DistanceData{}
	meLeft.Calculate([]Coord{you.Head().Left()}, b)
	meRight := &DistanceData{}
	meRight.Calculate([]Coord{you.Head().Right()}, b)
	meUp := &DistanceData{}
	meUp.Calculate([]Coord{you.Head().Up()}, b)
	meDown := &DistanceData{}
	meDown.Calculate([]Coord{you.Head().Down()}, b)

	b.Data = map[string]*DistanceData{
		"me_left": meLeft,
		"me_right": meRight,
		"me_up": meUp,
		"me_down": meDown,
	}
	for _, snake := range b.Snakes {
		b.Data[snake.ID] = &DistanceData{}
		b.Data[snake.ID].Calculate(snake.Head().Adjacent(), b)
	}

	b.AbleToVisitCount = map[string]int{
		"left": meLeft.Count,
		"right": meRight.Count,
		"up": meUp.Count,
		"down": meDown.Count,
	}
}