package logic

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/joram/jsnek/api"
)

type GoMoreRoom struct {
	Ratio float32
}

func (ge GoMoreRoom) Taunt() string {
	return "More Room"
}

func (ge GoMoreRoom) Decision(sr *api.SnakeRequest) int {
	sr.Board.PopulateDistances(sr.You)

	biggest := -1
	biggestDir := ""
	dirCounts := 0
	sumRoomAvail := 0
	for dir, count := range sr.Board.AbleToVisitCount {
		dirCounts += 1
		sumRoomAvail += count
		if count > biggest {
			biggest = count
			biggestDir = dir
		}
	}
	avgRoomAvail := sumRoomAvail/dirCounts

	if float32(biggest)/float32(avgRoomAvail) > ge.Ratio {
		d := map[string]int{
			"up": api.UP,
			"down": api.DOWN,
			"left": api.LEFT,
			"right":api.RIGHT,
			"WFT!": api.UNKNOWN,
		}[biggestDir]

		spew.Dump(sr.Board.AbleToVisitCount)
		fmt.Printf("going %s to avoid small space", biggestDir)
		return d
	}

	return api.UNKNOWN
}
