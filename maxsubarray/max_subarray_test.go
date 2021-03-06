package maxsubarray

import (
	"reflect"
	"testing"
)

func TestMaxSubarray(t *testing.T) {

	testCases := []struct {
		input []int
		want  []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{-1}, []int{-1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{0, -1}, []int{0}},
		{[]int{-1, -2}, []int{-1}},

		/* This test case comes from
		"Introduction to Algorithms - Third Edition" 4.1 Maximum Subarray */
		{
			[]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7},
			[]int{18, 20, -7, 12},
		},
	}

	for _, v := range testCases {
		out := FindMaxSubarrayByDivideAndConquer(v.input)
		if !reflect.DeepEqual(v.want, out) {
			t.Errorf("FindMaxSubarrayByDivideAndConquer on %v \ngot %v \nbut want %v", v.input, out, v.want)
		}
	}

	for _, v := range testCases {
		out := FindMaxSubarrayByViolent(v.input)
		if !reflect.DeepEqual(v.want, out) {
			t.Errorf("FindMaxSubarrayByViolent on %v \ngot %v \nbut want %v", v.input, out, v.want)
		}
	}

	for _, v := range testCases {
		out := FindMaxSubarrayByKadane(v.input)
		if !reflect.DeepEqual(v.want, out) {
			t.Errorf("FindMaxSubarrayByKadane on %v \ngot %v \nbut want %v", v.input, out, v.want)
		}
	}

}
