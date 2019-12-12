package logic

import "github.com/joram/jsnek/api"

type Responsibility interface {
	Decision(*api.SnakeRequest) int
	Taunt() string
}
