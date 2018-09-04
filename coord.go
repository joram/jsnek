package main

import (
	"math"
	"errors"
)

func (c Coord) Adjacent() []Coord {
	return []Coord{
		{c.X + 0, c.Y + 1},
		{c.X + 0, c.Y - 1},
		{c.X + 1, c.Y + 0},
		{c.X - 1, c.Y + 0},
	}
}

func (c Coord) Equal(other Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c Coord) DirectionTo(other Coord) int {
	xd := other.X - c.X
	yd := other.Y - c.Y
	if xd == +0 && yd == -1 { return UP }
	if xd == +0 && yd == +1 { return DOWN }
	if xd == -1 && yd == +0 { return LEFT }
	if xd == +1 && yd == +0 { return RIGHT }
	return UNKNOWN
}

func (c Coord) Offset(d int) (*Coord, error) {
	if d == UP {
		return &Coord{c.X, c.Y-1}, nil
	}
	if d == DOWN {
		return &Coord{c.X, c.Y+1}, nil
	}
	if d == LEFT {
		return &Coord{c.X-1, c.Y}, nil
	}
	if d == RIGHT {
		return &Coord{c.X+1, c.Y}, nil
	}
	return nil, errors.New("not a valid direction")
}

func (c Coord) NearestDirectionTo(other Coord) int {
	xd := other.X - c.X
	yd := other.Y - c.Y
	if xd == 0 && yd == 0 {
		return UNKNOWN
	}
	if math.Abs(float64(xd)) > math.Abs(float64(yd)) {
		if xd > 0 {
			return RIGHT
		}
		return LEFT
	}
	if yd > 0 {
		return DOWN
	}
	return UP
}

func (c Coord) OrthogonalDistance(other Coord) float64 {
	xd := float64(other.X - c.X)
	yd := float64(other.Y - c.Y)
	return math.Sqrt(xd*xd + yd*yd)
}
