package distance_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/distance"
)

var TestManhattanDistance = []struct {
	x1     int
	y1     int
	x2     int
	y2     int
	result int
}{
	{x1: 0, y1: 0, x2: 1, y2: 0, result: 1},
	{x1: 0, y1: 0, x2: -1, y2: 0, result: 1},
	{x1: 0, y1: 0, x2: 0, y2: 1, result: 1},
	{x1: 0, y1: 0, x2: 0, y2: -1, result: 1},
	{x1: 0, y1: 0, x2: 0, y2: 5, result: 5},
	{x1: -1, y1: 0, x2: 0, y2: 5, result: 6},
}

func TestGetManhattanDistance(t *testing.T) {
	for _, testData := range TestManhattanDistance {
		r := distance.GetManhattanDistance(testData.x1, testData.y1, testData.x2, testData.y2)
		m := fmt.Sprintf("GetManhattanDistance failed: X1: %d, Y2: %d, X2: %d, Y2: %d Expected Result was: %d", testData.x1, testData.y1, testData.x2, testData.y2, testData.result)
		assert.Equal(t, testData.result, r, m)
	}
}
