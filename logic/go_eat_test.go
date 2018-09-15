package logic

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/joram/jsnek/api"
)

func TestGoEat(t *testing.T) {
	ge := GoEatOrthogonal{25}
	sr := api.SnakeRequest{
		You:api.Snake{
			Body:[]api.Coord{{2,2}},
			Health: 26,
		},

		Board:api.Board{
			Width:5,
			Height:5,
			Snakes: []api.Snake{{Body:[]api.Coord{
				{1,2},
				{1,3},
				{2,3},
				{3,3},
				{3,2},
			},}},
			Food: []api.Coord{{2,0}},
		},
	}
	d := ge.decision(&sr)
	assert.Equal(t, api.UNKNOWN, d)

	sr.You.Health = 5
	d = ge.decision(&sr)
	assert.Equal(t, api.UP, d)

	sr.Board.Food[0] = api.Coord{2, 100}
	d = ge.decision(&sr)
	assert.Equal(t, api.UNKNOWN, d)

	sr.Board.Snakes = []api.Snake{}
	d = ge.decision(&sr)
	assert.Equal(t, api.DOWN, d)

}
