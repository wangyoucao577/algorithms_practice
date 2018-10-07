package mysorts

import "testing"

func TestSelectByPredefinedCases(t *testing.T) {
	cases := []struct {
		array intSlice
		i     int
		want  int
	}{
		{intSlice{}, 0, -1},
		{intSlice{}, -1, -1},
		{intSlice{0}, 1, 0},
		{intSlice{1}, 1, 0},
		{intSlice{0, 1}, 2, 1},
	}

	for _, v := range cases {
		got := RandomizedSelectNth(v.array, v.i)
		if got != v.want {
			t.Errorf("for %d-th of input array %v, expect %v but got %v", v.i, v.array, v.want, got)
		}
	}
}

func TestSelectByRandomCases(t *testing.T) {
	randomCaseCount := 100
	maxArrayLen := 1000

	for j := 0; j < randomCaseCount; j++ {
		//generate case
		testCase := generateRandomCase(maxArrayLen)
		if testCase.Len() == 0 {
			continue
		}

		//sort for check result
		caseForQuickSort := testCase.deepCopy()
		QuickSort(caseForQuickSort)
		if !isSorted(caseForQuickSort, true) {
			t.Errorf("in-place QuickSort failed on \n%v\n", caseForQuickSort)
			break
		}

		median := testCase.Len() / 2
		caseForSelection := testCase.deepCopy()
		got := RandomizedSelectNth(caseForSelection, median)
		gotV := caseForSelection[got]
		medianV := caseForQuickSort[median-1]
		if gotV != medianV {
			t.Errorf("Select median %v in %v, want %v but got %v", median, testCase, medianV, gotV)
			break
		}

		for k, selectedV := range caseForSelection {
			if k < got {
				// expect all elements before got are smaller than value of got
				if selectedV > medianV {
					t.Errorf("expect %v <= %v but not, array after select: %v", selectedV, medianV, caseForQuickSort)
					break
				}
			} else if k > got {
				// expect all elements before got are bigger than value of got
				if selectedV < medianV {
					t.Errorf("expect %v >= %v but not, array after select: %v", selectedV, medianV, caseForQuickSort)
					break
				}
			}
		}
	}
}

func BenchmarkWorstCaseSelectionMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCase := generateWorstCase(benchmarkMaxArrayLen)
		median := testCase.Len() / 2
		RandomizedSelectNth(testCase, median)
	}
}

func BenchmarkBestCaseSelectionMedian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testCase := generateBestCase(benchmarkMaxArrayLen)
		median := testCase.Len() / 2
		RandomizedSelectNth(testCase, median)
	}
}
