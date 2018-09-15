package logic

import "github.com/joram/jsnek/api"

type GoEatOrthogonal struct {
	HungryHealth int
}

func (ge GoEatOrthogonal) Taunt() string {
	return "EATING"
}

func (ge GoEatOrthogonal) Decision(sr *api.SnakeRequest) int {
	longestSnakeLength := -1
	for _, snake := range sr.Board.Snakes {
		if snake.ID != sr.You.ID {
			l := len(snake.Body)
			if l > longestSnakeLength {
				longestSnakeLength = l
			}
		}
	}

	if sr.You.Health > ge.HungryHealth && len(sr.You.Body) >= longestSnakeLength{
		return api.UNKNOWN
	}

	closestFood, err := sr.Board.ClosestFood(sr.You.Head())
	if err != nil {
		return api.UNKNOWN
	}

	d := sr.You.Head().NearestDirectionTo(*closestFood)
	nextCoord, err := sr.You.Head().Offset(d)
	if err != nil {
		return api.UNKNOWN
	}

	if sr.Board.IsEmpty(*nextCoord) {
		return d
	}

	return api.UNKNOWN
}
