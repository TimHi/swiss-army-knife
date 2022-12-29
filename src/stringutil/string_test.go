package stringutil_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/stringutil"
)

func TestIsUpperTrue(t *testing.T) {
	isUpper := stringutil.IsUpper("AAAA\n")
	assert.True(t, isUpper)
	isUpper = stringutil.IsUpper("A")
	assert.True(t, isUpper)
	isUpper = stringutil.IsUpper("   B")
	assert.True(t, isUpper)
}

func TestIsntUpper(t *testing.T) {
	isUpper := stringutil.IsUpper("AaAA\n")
	assert.False(t, isUpper)
	isUpper = stringutil.IsUpper("a")
	assert.False(t, isUpper)
	isUpper = stringutil.IsUpper("   Ba")
	assert.False(t, isUpper)
	isUpper = stringutil.IsUpper("")
	assert.False(t, isUpper)
}

func TestNotEmpty(t *testing.T) {
	notEmpty := "\n"
	isEmpty := stringutil.NotEmpty(notEmpty)
	assert.True(t, isEmpty)

	notEmpty = "A \n"
	isEmpty = stringutil.NotEmpty(notEmpty)
	assert.True(t, isEmpty)

	empty := ""
	isEmpty = stringutil.NotEmpty(empty)
	assert.False(t, isEmpty)
}
