package logic

import "github.com/joram/jsnek/api"

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidDirection(t *testing.T) {
	vd := ValidDirection{}
	sr := api.SnakeRequest{
		You:api.Snake{
			Body:[]api.Coord{{2,2}},
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
		},
	}
	d := vd.decision(&sr)
	assert.Equal(t, d, api.UP)
}
