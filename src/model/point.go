package model

import "github.com/timhi/swiss-army-knife/src/distance"

type Point struct {
	X int
	Y int
}

func (orig Point) ManhattanDistance(dest Point) int {
	return distance.GetManhattanDistance(orig.X, orig.Y, dest.X, dest.Y)
}
