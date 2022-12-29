package numbers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/numbers"
)

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
