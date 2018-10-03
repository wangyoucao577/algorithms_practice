package mysorts

import (
	"testing"
)

func TestSortFunctionalWithRandomCases(t *testing.T) {
	randomCaseCount := 100
	maxArrayLen := 1000

	for i := 0; i < randomCaseCount; i++ {
		//generate case
		testCase := generateRandomCase(maxArrayLen)

		//insertion sort
		caseForInsertionSort := testCase.deepCopy()
		InsertionSort(caseForInsertionSort)
		if !isSorted(caseForInsertionSort, true) {
			t.Errorf("insertionSort failed on \n%v\n", caseForInsertionSort)
			break
		}

		//merge sort in-place implementation
		caseForMergeSort := testCase.deepCopy()
		MergeSortInPlace(caseForMergeSort)
		if !isSorted(caseForMergeSort, true) {
			t.Errorf("in-place mergeSort failed on \n%v\n", caseForMergeSort)
			break
		}

		//merge sort aux array based implementation
		// NOTE: can not done by only sort.Interface
		caseForMergeSortAuxArrayBased := testCase.deepCopy()
		MergeSort(caseForMergeSortAuxArrayBased)
		if !isSorted(caseForMergeSortAuxArrayBased, true) {
			t.Errorf("aux array based mergeSort failed on \n%v\n", caseForMergeSortAuxArrayBased)
			break
		}

		//heap sort in-place implementation
		caseForHeapSort := testCase.deepCopy()
		HeapSort(caseForHeapSort)
		if !isSorted(caseForHeapSort, true) {
			t.Errorf("in-place heapSort failed on \n%v\n", caseForHeapSort)
			break
		}

	}
}
