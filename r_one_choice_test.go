package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOnlyOneChoice(t *testing.T) {
	ooc := OnlyOneChoice{}
	sr := SnakeRequest{
		You: Snake{
			Body: []Coord{{2, 2}},
		},

		Board: Board{
			Width:  5,
			Height: 5,
			Snakes: []Snake{{Body: []Coord{
				{1, 2},
				{1, 3},
				{2, 3},
				{3, 3},
				{3, 2},
			}}},
		},
	}
	d := ooc.decision(&sr)
	assert.Equal(t, d, UP)
}
