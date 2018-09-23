package mysorts

import (
	"math/rand"
	"testing"
)

func generateRandomCase(maxArrayLen int) []int {
	arrayLen := rand.Intn(maxArrayLen)

	result := make([]int, arrayLen, arrayLen)
	for i := 0; i < arrayLen; i++ {
		result[i] = rand.Intn(maxArrayLen)

		//TODO: generate negative numbers
	}
	return result
}

func isSorted(in []int, ascending bool) bool {

	for i := 0; i < len(in)-1; i++ {
		if ascending {
			if in[i] > in[i+1] {
				return false
			}
		} else {
			if in[i] < in[i+1] {
				return false
			}
		}
	}
	return true
}

func TestSortFunctionalWithRandomCases(t *testing.T) {
	randomCaseCount := 100
	maxArrayLen := 1000

	for i := 0; i < randomCaseCount; i++ {
		//generate case
		testCase := generateRandomCase(maxArrayLen)

		//insertion sort
		insertionSort(testCase)
		if !isSorted(testCase, true) {
			t.Errorf("insertionSort failed on \n%v\n", testCase)
			break
		}
	}
}
