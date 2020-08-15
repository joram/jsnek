package logic

import (
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/util"
)

type EatWhenShortestSnake struct {
	IgnoreHazardFood   bool
	LengthCompensation int
}

func (ss EatWhenShortestSnake) Taunt() string {
	return "EATING (Shortest Snake)"
}

func (ss EatWhenShortestSnake) Decision(sr *api.SnakeRequest) int {
	longestSnakeLength := -1
	for _, snake := range sr.OtherSnakes() {
		longestSnakeLength = util.Max(longestSnakeLength, len(snake.Body))
	}

	if len(sr.You.Body) >= longestSnakeLength+ss.LengthCompensation {
		return api.UNKNOWN
	}

	return GoEatOrthogonal{ss.IgnoreHazardFood, 100, }.Decision(sr)
}
