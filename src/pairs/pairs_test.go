package pairs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/pairs"
)

func TestPairListLength(t *testing.T) {
	testList := pairs.PairList{}
	assert.Empty(t, testList, "Initilized PairList not empty")
	assert.Equal(t, 0, testList.Len(), "Initilized PairList not empty")
	testList = append(testList, pairs.Pair{0, 0})
	assert.Equal(t, 1, testList.Len(), "PairList length did not update")
}

func TestPairListLess(t *testing.T) {
	testListLess := pairs.PairList{pairs.Pair{0, 0}, pairs.Pair{1, 1}, pairs.Pair{2, 2}}
	isLess := testListLess.Less(0, 1)
	isNotLess := testListLess.Less(2, 1)
	assert.True(t, isLess, "Element at index 0 was bigger than element at index 1")
	assert.False(t, isNotLess, "Element at index 1 was not bigger than element at index 2")
}

func TestPairListSwap(t *testing.T) {
	testListLess := pairs.PairList{pairs.Pair{0, 0}, pairs.Pair{1, 1}, pairs.Pair{2, 2}}
	testListLess.Swap(0, 1)
	assert.Equal(t, pairs.Pair{1, 1}, testListLess[0], "Swapping elements failed")
	assert.Equal(t, pairs.Pair{0, 0}, testListLess[1], "Swapping elements failed")
}

func TestSortMapToPairsReversed(t *testing.T) {
	mapToSort := map[int]int{}
	mapToSort[0] = 1
	mapToSort[1] = 10
	mapToSort[2] = -1
	actualList := pairs.SortMapToPairsReversed(mapToSort)
	expectedList := pairs.PairList{{1, 10}, {0, 1}, {2, -1}}

	assert.Equal(t, expectedList, actualList, "Sorting the map to pairs failed")
}

func TestSortMapToPairs(t *testing.T) {
	mapToSort := map[int]int{}
	mapToSort[0] = 1
	mapToSort[1] = 10
	mapToSort[2] = -1
	actualList := pairs.SortMapToPairs(mapToSort)
	expectedList := pairs.PairList{{2, -1}, {0, 1}, {1, 10}}

	assert.Equal(t, expectedList, actualList, "Sorting the map to pairs failed")
}
