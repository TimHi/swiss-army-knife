package distance

import (
	"github.com/timhi/swiss-army-knife/src/numbers"
)

// Calculate the "Taxicab Geometry" from given coordinate X1,Y1 to coordinate X2,Y2.
// See https://en.wikipedia.org/wiki/Taxicab_geometry for more details.
func GetManhattanDistance(x1, y1, x2, y2 int) int {
	return numbers.Abs(x2-x1) + numbers.Abs(y2-y1)
}
