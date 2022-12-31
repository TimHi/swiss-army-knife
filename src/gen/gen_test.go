package gen_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/timhi/swiss-army-knife/src/gen"
)

var TestIsSliceData = []struct {
	input  interface{}
	result bool
}{
	{[]int{0}, true},
	{[]int{}, true},
	{[]string{""}, true},
	{[]string{}, true},
	{"", false},
	{0, false},
}

func TestIsSlice(t *testing.T) {
	for _, testData := range TestIsSliceData {
		r := gen.IsSlice(testData.input)
		m := fmt.Sprintln("IsSlice failed")
		assert.Equal(t, testData.result, r, m)
	}
}

var TestIsBiggerData = []struct {
	a      interface{}
	b      interface{}
	result bool
}{
	{[]int{1, 2, 3, 4}, []int{1, 2, 3}, true},
	{[]int{1, 2, 3}, []int{1, 2, 3}, false},
	{[]int{1, 2}, []int{1, 2, 3}, false},
	{[]string{}, []int{}, false},
	{[]string{"1"}, []int{}, true},
	{[]string{"1", "2"}, []int{1}, true},
}

func TestIsBigger(t *testing.T) {
	for _, testData := range TestIsBiggerData {
		a := gen.InterfaceSlice(testData.a)
		b := gen.InterfaceSlice(testData.b)
		r := gen.IsBigger(a, b)
		m := fmt.Sprintln("IsSlice failed")
		assert.Equal(t, testData.result, r, m)
	}
}
