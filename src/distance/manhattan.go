package distance

import (
	"github.com/timhi/swiss-army-knife/src/numbers"
)

func GetManhattanDistance(x1, y1, x2, y2 int) int {
	return numbers.Abs(x2-x1) + numbers.Abs(y2-y1)
}
