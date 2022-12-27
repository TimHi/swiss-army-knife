package model_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/model"
)

var TestManhattanDistance = []struct {
	orig   model.Point
	dest   model.Point
	result int
}{
	{orig: model.Point{X: 0, Y: 0}, dest: model.Point{X: 1, Y: 0}, result: 1},
	{orig: model.Point{X: 0, Y: 0}, dest: model.Point{X: -1, Y: 0}, result: 1},
	{orig: model.Point{X: 0, Y: 0}, dest: model.Point{X: 0, Y: 1}, result: 1},
	{orig: model.Point{X: 0, Y: 0}, dest: model.Point{X: 0, Y: -1}, result: 1},
}

func TestPointGetManhattanDistance(t *testing.T) {
	for _, testData := range TestManhattanDistance {
		r := testData.orig.ManhattanDistance(testData.dest)
		m := fmt.Sprintf("GetManhattanDistance failed: X1: %d, Y2: %d, X2: %d, Y2: %d Expected Result was: %d", testData.orig.X, testData.orig.Y, testData.dest.X, testData.dest.Y, testData.result)
		assert.Equal(t, testData.result, r, m)
	}
}
