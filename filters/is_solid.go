package filters

import "github.com/joram/jsnek/api"

type IsSolidFilter struct{}

func (isf IsSolidFilter) Description() string { return "is solid" }

func (isf IsSolidFilter) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	coord, err := sr.You.GetHead().Offset(direction)
	if err != nil {
		return false, err
	}
	return sr.Board.IsEmpty(*coord), nil
}
