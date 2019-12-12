package api

func (s *Snake) Head() Coord {
	return s.Body[0]
}

func (s *Snake) Tail() Coord {
	return s.Body[len(s.Body)-1]
}

func (s *Snake) JustAte() bool {
	tailNeck := s.Body[len(s.Body)-2]
	return s.Tail().Equal(tailNeck)
}
