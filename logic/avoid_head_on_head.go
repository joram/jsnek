package logic

import (
	"github.com/joram/jsnek/api"
	"math/rand"
)

type AvoidHeadOnHead struct{}

func (ge AvoidHeadOnHead) Taunt() string {
	return "Avoid Potential Collision"
}

func (ge AvoidHeadOnHead) Decision(sr *api.SnakeRequest) int {

	dangerCoords := []api.Coord{}
	for _, snake := range sr.Board.Snakes {
		if snake.ID != sr.You.ID && len(snake.Body) >= len(sr.You.Body) {
			dangerCoords = append(dangerCoords, snake.GetHead().Adjacent()...)
		}
	}

	couldBeEaten := false
	safeDirections := []string{}
	avoidSquares := []api.Coord{}
	avoidSquares = append(avoidSquares, dangerCoords...)
	for dir, coord := range sr.You.GetHead().AdjacentMap() {
		badCoord := false
		for _, dangerCoord := range dangerCoords {
			if coord.Equal(dangerCoord) {
				badCoord = true
				break
			}
		}
		if badCoord {
			couldBeEaten = true
		} else {
			if sr.Board.IsEmpty(coord) {
				safeDirections = append(safeDirections, dir)
			}
		}
	}

	// Could GetHead on GetHead
	if couldBeEaten && len(safeDirections) > 0 {
		i := rand.Intn(len(safeDirections))
		safeDir := safeDirections[i]
		return api.StringToDir(safeDir)
	}

	return api.UNKNOWN
}
