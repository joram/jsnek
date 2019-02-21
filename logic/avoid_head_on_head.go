package logic

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/joram/jsnek/api"
	"math/rand"
)

type AvoidHeadOnHead struct {}

func (ge AvoidHeadOnHead) Taunt() string {
	return "Avoid Potential Collision"
}

func (ge AvoidHeadOnHead) Decision(sr *api.SnakeRequest) int {

	dangerCoords := []api.Coord{}
	for _, snake := range sr.Board.Snakes {
		if snake.ID != sr.You.ID && len(snake.Body) >= len(sr.You.Body){
			dangerCoords = append(dangerCoords, snake.Head().Adjacent()...)
		}
	}

	couldBeEaten := false
	safeDirections := []string{}
	avoidSquares := []api.Coord{}
	avoidSquares = append(avoidSquares, dangerCoords...)
	for dir, coord := range sr.You.Head().AdjacentMap() {
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
			if sr.Board.IsEmpty(coord){
				safeDirections = append(safeDirections, dir)
			}
		}
	}

	println(couldBeEaten)
	spew.Dump(safeDirections)

	// Could Head on Head
	if couldBeEaten {
		i := rand.Intn(len(safeDirections))
		safeDir := safeDirections[i]
		return api.StringToDir(safeDir)
	}

	return api.UNKNOWN
}
