package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGoEat(t *testing.T) {
	ge := GoEatOrthogonal{25}
	sr := SnakeRequest{
		You:Snake{
			Body:[]Coord{{2,2}},
			Health: 26,
		},

		Board:Board{
			Width:5,
			Height:5,
			Snakes: []Snake{{Body:[]Coord{
				{1,2},
				{1,3},
				{2,3},
				{3,3},
				{3,2},
			},}},
			Food: []Coord{{2,0}},
		},
	}
	d := ge.decision(&sr)
	assert.Equal(t, UNKNOWN, d)

	sr.You.Health = 5
	d = ge.decision(&sr)
	assert.Equal(t, UP, d)

	sr.Board.Food[0] = Coord{2, 100}
	d = ge.decision(&sr)
	assert.Equal(t, UNKNOWN, d)

	sr.Board.Snakes = []Snake{}
	d = ge.decision(&sr)
	assert.Equal(t, DOWN, d)

}
