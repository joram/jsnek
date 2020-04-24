package filters

import "github.com/joram/jsnek/api"

type IsThreatenedFilter struct{}

func (isf IsThreatenedFilter) Description() string { return "is threatened" }
func (isf IsThreatenedFilter) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	coord, err := sr.You.Head().Offset(direction)
	if err != nil {
		return false, err
	}
	for _, snake := range sr.OtherSnakes() {
		head := snake.Head()
		for _, threatenedCoord := range head.SurroundingCoords() {
			if coord.Equal(threatenedCoord) {
				return false, nil
			}
		}
	}
	return true, nil
}
