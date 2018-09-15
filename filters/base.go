package filters

import "github.com/joram/jsnek/api"

type DecisionFilter interface {
	Allowed(int, *api.SnakeRequest) (bool, error)
	Description() string
}
