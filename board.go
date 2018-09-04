package main

import "errors"

func (b *Board) IsEmpty(c Coord) bool {
	if c.X >=b.Width {
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
