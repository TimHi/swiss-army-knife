package numbers_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/numbers"
)

func TestMaxInt(t *testing.T) {
	var expected int
	if ^uint(0) == 32 {
		expected = 1<<31 - 1
	} else {
		expected = 1<<63 - 1
	}
	result := numbers.MaxInt()
	assert.Equal(t, expected, result)
}

func TestAbsNegative(t *testing.T) {
	result := numbers.Abs(-1)
	assert.Equal(t, 1, result)
}

func TestAbsPositive(t *testing.T) {
	result := numbers.Abs(100)
	assert.Equal(t, 100, result)
}

func TestAbsNull(t *testing.T) {
	result := numbers.Abs(0)
	assert.Equal(t, 0, result)
}

var TestModData = []struct {
	x      int64
	y      int64
	result int64
}{
	{0, 2, 0},
	{2, 2, 0},
	{4, 2, 0},
	{5, 2, 1},
	{9, 2, 1},
}

func TestMod(t *testing.T) {
	for _, testData := range TestModData {
		r, e := numbers.Mod(testData.x, testData.y)
		m := fmt.Sprintln("Mod failed")
		assert.Equal(t, testData.result, r, m)
		assert.Nil(t, e, m)
	}
}

func TestModDivZero(t *testing.T) {
	r, e := numbers.Mod(0, 0)
	m := fmt.Sprintln("Mod failed")
	assert.Equal(t, int64(0), r, m)
	assert.NotNil(t, e, m)
}
