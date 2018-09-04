package main

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
