package logic

import "github.com/joram/jsnek/api"

type GoSurroundOtherSnake struct {
}

func (gsos GoSurroundOtherSnake) Taunt() string {
	return "EATING"
}

func (gsos GoSurroundOtherSnake) Decision(sr *api.SnakeRequest) int {
	var targetSnake = *api.Snake{}
	for _, otherSnake := range sr.OtherSnakes(){
		if otherSnake.Health < sr.You.Health {
			
		}
	}
}
