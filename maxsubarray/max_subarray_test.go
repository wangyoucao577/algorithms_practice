package maxsubarray

import (
	"reflect"
	"testing"
)

func TestMaxSubarray(t *testing.T) {

	/* This test case comes from
	"Introduction to Algorithms - Third Edition" 4.1 Maximum Subarray
	*/
	testCase := struct {
		input []int
		want  []int
	}{
		[]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
		[]int{18, 20, -7, 12},
	}

	out := findMaxSubarray(testCase.input)
	if !reflect.DeepEqual(testCase.want, out) {
		t.Errorf("findMaxSubarray on %v \ngot %v \nbut want %v", testCase.input, out, testCase.want)
	}

}
