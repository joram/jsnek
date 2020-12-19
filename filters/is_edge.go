package filters

import "github.com/joram/jsnek/api"

type IsEdge struct{}

func (ie IsEdge) Description() string { return "is edge" }
func (ie IsEdge) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	coord, err := sr.You.GetHead().Offset(direction)
	if coord == nil {
		return true, err
	}

	if coord.X == 0 {
		return false, nil
	}
	if coord.Y == 0 {
		return false, nil
	}
	if coord.X+1 == sr.Board.Width {
		return false, nil
	}
	if coord.Y+1 == sr.Board.Height {
		return false, nil
	}
	return true, nil
}
