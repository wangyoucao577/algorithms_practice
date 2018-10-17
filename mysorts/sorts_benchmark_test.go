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

func BenchmarkBestCaseBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BucketSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseBinarySearchTreeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchTreeSort(generateBestCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkBestCaseRedBlackTreeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedBlackTreeSort(generateBestCase(benchmarkMaxArrayLen))
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

func BenchmarkWorstCaseBucketSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BucketSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseBinarySearchTreeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BinarySearchTreeSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}

func BenchmarkWorstCaseRedBlackTreeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedBlackTreeSort(generateWorstCase(benchmarkMaxArrayLen))
	}
}
