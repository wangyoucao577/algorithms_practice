package mysorts

import "testing"

const (
	benchmarkMaxArrayLen int = 1000
)

func BenchmarkBestCaseInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortInPlace(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(generateBestCase(benchmarkMaxArrayLen))
	}
}
func BenchmarkBestCaseQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseCountingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseRadixSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RadixSort(generateBestCase(benchmarkMaxArrayLen))
	}
}
func BenchmarkWorstCaseInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortInPlace(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseCountingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseRadixSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RadixSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}
