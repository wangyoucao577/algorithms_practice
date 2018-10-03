package mysorts

import "testing"

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

func BenchmarkBestCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heapSort(generateBestCase(benchmarkMaxArrayLen))
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

func BenchmarkWorstCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heapSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}
