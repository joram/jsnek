package filters

import "github.com/joram/jsnek/api"

type IsSolidFilter struct {}

func (isf IsSolidFilter) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	coord, err := sr.You.Head().Offset(direction)
	if err != nil {
		return false, err
	}
	if sr.Board.IsEmpty(*coord) {
		return true, nil
	}
	return false, nil
}
