package filters

import "github.com/joram/jsnek/api"

type IsUnknownFilter struct {}

func (isf IsUnknownFilter) Allowed(direction int, sr *api.SnakeRequest) (bool, error) {
	return direction != api.UNKNOWN, nil
}