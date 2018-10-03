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
		MergeSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortAuxArrayBased(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseMergeSortAuxArrayBased(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortAuxArrayBased(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseHeapSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HeapSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}
