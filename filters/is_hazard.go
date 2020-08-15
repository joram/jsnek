package filters

import "github.com/joram/jsnek/api"

type IsHazardFilter struct{}

func (isf IsHazardFilter) Description() string { return "is hazard" }
func (isf IsHazardFilter) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	coord, err := sr.You.GetHead().Offset(direction)
	if coord == nil {
		return true, err
	}

	for _, hazard := range sr.Board.Hazards {
		if hazard.Equal(*coord) {
			return false, nil
		}
	}
	return true, nil
}
