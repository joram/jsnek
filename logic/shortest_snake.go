package logic

import (
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/util"
)

type ShortestSnake struct {
	LengthCompensation int
}

func (ss ShortestSnake) Taunt() string {
	return "EATING (Shortest Snake)"
}

func (ss ShortestSnake) Decision(sr *api.SnakeRequest) int {
	longestSnakeLength := -1
	for _, snake := range sr.OtherSnakes() {
		longestSnakeLength = util.Max(longestSnakeLength, len(snake.Body))
	}

	if len(sr.You.Body) >= longestSnakeLength+ss.LengthCompensation {
		return api.UNKNOWN
	}

	return GoEatOrthogonal{100}.Decision(sr)
}
