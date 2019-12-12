package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	UNKNOWN
)

var DIRECTIONS = [...]int{
	UP,
	DOWN,
	LEFT,
	RIGHT,
}

func DirToString(d int) string {
	return map[int]string{
		UP:      "up",
		DOWN:    "down",
		LEFT:    "left",
		RIGHT:   "right",
		UNKNOWN: "WFT!",
	}[d]
}

func StringToDir(s string) int {
	return map[string]int{
		"up":    UP,
		"down":  DOWN,
		"left":  LEFT,
		"right": RIGHT,
		"WFT!":  UNKNOWN,
	}[s]
}

type Coord struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (c *Coord) String() string {
	return fmt.Sprintf("%d_%d", c.X, c.Y)
}

type Snake struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Health int     `json:"health"`
	Body   []Coord `json:"body"`
}

type Board struct {
	Height           int                      `json:"height"`
	Width            int                      `json:"width"`
	Food             []Coord                  `json:"food"`
	Snakes           []Snake                  `json:"snakes"`
	Data             map[string]*DistanceData `json:"data"`
	AbleToVisitCount map[string]int           `json:"able_to_visit"`
}

type Game struct {
	ID string `json:"id"`
}

type SnakeRequest struct {
	Game  Game  `json:"game"`
	Turn  int   `json:"turn"`
	Board Board `json:"board"`
	You   Snake `json:"you"`
}

type StartResponse struct {
	Color string `json:"color,omitempty"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Taunt string `json:"taunt"`
}

func DecodeSnakeRequest(req *http.Request, decoded *SnakeRequest) error {
	err := json.NewDecoder(req.Body).Decode(&decoded)
	return err
}
