package logic

import (
	"github.com/joram/jsnek/api"
	"math/rand"
)

type AvoidHeadOnHead struct {}

func (ge AvoidHeadOnHead) Taunt() string {
	return "Avoid Potential Collision"
}

func (ge AvoidHeadOnHead) Decision(sr *api.SnakeRequest) int {
	couldBeEaten := false
	safeDirections := []string{}
	avoidSquares := []api.Coord{}
	for _, snake := range sr.Board.Snakes {
		if snake.ID != sr.You.ID && len(snake.Body) >= len(sr.You.Body){
			dangerCoords := snake.Head().Adjacent()
			avoidSquares = append(avoidSquares, dangerCoords...)
			for dir, coord := range sr.You.Head().AdjacentMap() {
				for _, dangerCoord := range dangerCoords {
					if coord.Equal(dangerCoord) {
						couldBeEaten = true
					} else {
						safeDirections = append(safeDirections, dir)
					}
				}
			}
		}
	}

	// Could Head on Head
	if couldBeEaten {
		i := rand.Intn(len(safeDirections))
		safeDir := safeDirections[i]
		return safeDir
	}

	return api.UNKNOWN
}
