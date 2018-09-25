package mysorts

import (
	"math/rand"
	"testing"
)

type intSlice []int

func (s intSlice) Len() int           { return len(s) }
func (s intSlice) Less(i, j int) bool { return s[i] < s[j] }
func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intSlice) deepCopy() intSlice {
	newCopy := intSlice{}
	newCopy = append(newCopy, s...)
	return newCopy
}

func generateRandomCase(maxArrayLen int) intSlice {
	arrayLen := rand.Intn(maxArrayLen)

	result := make([]int, arrayLen, arrayLen)
	for i := 0; i < arrayLen; i++ {
		result[i] = rand.Intn(maxArrayLen)

		//TODO: generate negative numbers
	}
	return result
}

func generateBestCase(maxArrayLen int) intSlice {
	bestCase := make([]int, maxArrayLen, maxArrayLen)
	for i := 0; i < maxArrayLen; i++ {
		bestCase[i] = i // ascending sorted
	}
	return bestCase
}

func generateWorstCase(maxArrayLen int) intSlice {
	worstCase := make([]int, maxArrayLen, maxArrayLen)
	for i := 0; i < maxArrayLen; i++ {
		worstCase[i] = maxArrayLen - i // descending sorted by expect ascending
	}
	return worstCase
}

func isSorted(in intSlice, ascending bool) bool {

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
		caseForInsertionSort := testCase.deepCopy()
		insertionSort(caseForInsertionSort)
		if !isSorted(caseForInsertionSort, true) {
			t.Errorf("insertionSort failed on \n%v\n", caseForInsertionSort)
			break
		}

		//merge sort in-place implementation
		caseForMergeSort := testCase.deepCopy()
		mergeSort(caseForMergeSort)
		if !isSorted(caseForMergeSort, true) {
			t.Errorf("in-place mergeSort failed on \n%v\n", caseForMergeSort)
			break
		}

		//merge sort aux array based implementation
		// NOTE: can not done by only sort.Interface
		caseForMergeSortAuxArrayBased := testCase.deepCopy()
		mergeSortAuxArrayBased(caseForMergeSortAuxArrayBased)
		if !isSorted(caseForMergeSortAuxArrayBased, true) {
			t.Errorf("aux array based mergeSort failed on \n%v\n", caseForMergeSortAuxArrayBased)
			break
		}

	}
}

const (
	benchmarkMaxArrayLen int = 1000
)

func BenchmarkBestCaseInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSortAuxArrayBased(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertionSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mergeSortAuxArrayBased(generateWorstCase(benchmarkMaxArrayLen))
	}
}
