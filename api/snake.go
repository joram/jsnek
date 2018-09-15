package api

func (s *Snake) Head() Coord {
	return s.Body[0]
}

func (s *Snake) Tail() Coord {
	return s.Body[len(s.Body)-1]
}
