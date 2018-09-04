package main

type GoEatOrthogonal struct {
	hungryHealth int
}

func (ge GoEatOrthogonal) taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) decision(sr *SnakeRequest) int {
	longestSnakeLength := -1
	for _, snake := range sr.Board.Snakes {
		if snake.ID != sr.You.ID {
			l := len(snake.Body)
			if l > longestSnakeLength {
				longestSnakeLength = l
			}
		}
	}

	if sr.You.Health > ge.hungryHealth && len(sr.You.Body) >= longestSnakeLength{
		return UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.Head())
	if err != nil {
		return UNKNOWN
	}

	d := sr.You.Head().NearestDirectionTo(*closestFood)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return UNKNOWN
}
