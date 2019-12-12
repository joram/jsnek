package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var mockReq2 = `{
	"game": {
		"id": "abc"
	},
	"turn": 10,
	"board": {
		"height": 10,
		"width": 10,
		"food": [{
				"x": 1,
				"y": 1
		}],
		"snakes": [{
			"id": "123",
			"name": "snek",
			"health": 1,
			"body": [{
					"x": 1,
					"y": 1
			}]
		}]
	},
	"you": {
		"id": "123",
		"name": "snek",
		"health": 1,
		"body": [{
				"x": 1,
				"y": 1
		}]
	}
}`

func TestBoardControl(t *testing.T) {
	expected := SnakeRequest{
		Game: Game{
			ID: "abc",
		},
		Turn: 10,
		Board: Board{
			Width:  10,
			Height: 10,
			Food: []Coord{
				Coord{1, 1},
			},
			Snakes: []Snake{
				{
					ID:     "123",
					Name:   "snek",
					Health: 1,
					Body: []Coord{
						Coord{1, 1},
					},
				},
			},
		},
		You: Snake{
			ID:     "123",
			Name:   "snek",
			Health: 1,
			Body: []Coord{
				Coord{1, 1},
			},
		},
	}

	result := SnakeRequest{}
	req := requestWithBody(mockReq2)
	err := DecodeSnakeRequest(req, &result)

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	result.Board.PopulateDistances(result.You)
	//move := logic.GoMoreRoom{}.Decision(&result)
	//assert.Equal(t, 1, move)

}
