package api

import (
	"errors"
	"math"
)

func (c Coord) AdjacentMap() map[string]Coord {
	return map[string]Coord{
		"down":  {c.X + 0, c.Y - 1},
		"up":    {c.X + 0, c.Y + 1},
		"right": {c.X + 1, c.Y + 0},
		"left":  {c.X - 1, c.Y + 0},
	}
}

func (c Coord) Adjacent() []Coord {
	return []Coord{
		c.Up(),
		c.Down(),
		c.Left(),
		c.Right(),
	}
}

func (c Coord) IsAdjacent(other Coord) bool {
	for _, a := range c.Adjacent() {
		if a.Equal(other) {
			return true
		}
	}
	return false
}

func (c Coord) SurroundingCoords() []Coord {
	return []Coord{
		c.Up(),
		c.Down(),
		c.Left(),
		c.Right(),

		{c.X + 1, c.Y + 1},
		{c.X - 1, c.Y - 1},
		{c.X + 1, c.Y - 1},
		{c.X - 1, c.Y + 1},
	}

}

func (c Coord) Equal(other Coord) bool {
	return c.X == other.X && c.Y == other.Y
}

func (c Coord) Left() Coord  { return Coord{X: c.X - 1, Y: c.Y} }
func (c Coord) Right() Coord { return Coord{X: c.X + 1, Y: c.Y} }
func (c Coord) Up() Coord    { return Coord{X: c.X, Y: c.Y + 1} }
func (c Coord) Down() Coord  { return Coord{X: c.X, Y: c.Y - 1} }

func (c Coord) DirectionTo(other Coord) int {
	xd := other.X - c.X
	yd := other.Y - c.Y
	if xd == +0 && yd == +1 {
		return UP
	}
	if xd == +0 && yd == -1 {
		return DOWN
	}
	if xd == -1 && yd == +0 {
		return LEFT
	}
	if xd == +1 && yd == +0 {
		return RIGHT
	}
	return UNKNOWN
}

func (c Coord) Offset(d int) (*Coord, error) {
	if d == UP {
		return &Coord{c.X, c.Y + 1}, nil
	}
	if d == DOWN {
		return &Coord{c.X, c.Y - 1}, nil
	}
	if d == LEFT {
		return &Coord{c.X - 1, c.Y}, nil
	}
	if d == RIGHT {
		return &Coord{c.X + 1, c.Y}, nil
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
		return UP
	}
	return DOWN
}

func (c Coord) OrthogonalDistance(other Coord) float64 {
	xd := float64(other.X - c.X)
	yd := float64(other.Y - c.Y)
	return math.Sqrt(xd*xd + yd*yd)
}
