package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/joram/jsnek/api"
	"github.com/joram/jsnek/util"
)

type FailedGame struct {
	gameID string
	turn int
	validDirections []int
	gameRequest *api.SnakeRequest
}

func failedGames() []FailedGame{
	var failedGames = []FailedGame{
		{
			gameID: "25f58e9b-d100-4ce6-851c-e3e6769ae48f",
			turn:116,
			validDirections: []int{api.LEFT, api.DOWN},
		},
	}
	return failedGames
}

func (fg *FailedGame) getGameRequest() *api.SnakeRequest {
	if fg.gameRequest != nil {
		return fg.gameRequest
	}

	var gameStates []api.SnakeRequest
	content := util.GetFromS3("jsnek", fmt.Sprintf("%s.json", fg.gameID))
	err := json.Unmarshal(content, &gameStates)
	if err != nil {
		panic(err)
	}

	for _, gameState := range gameStates {
		if gameState.Turn == fg.turn {
			fg.gameRequest = &gameState
			return &gameState
		}
	}
	return nil
}

func main() {
	for _, game := range failedGames() {
		game.getGameRequest()
		spew.Dump(game)
	}
}
