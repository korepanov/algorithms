package sort

import "testing"

func Test_mergeSort(t *testing.T) {
	testSort(t, MergeSort)
}

func Benchmark_mergeSort(b *testing.B) {
	benchmarkSort(b, MergeSort)
}

func Test_mergeSortParallel(t *testing.T) {
	testSort(t, MergeSortParallel)
}

func Benchmark_mergeSortParallel(b *testing.B) {
	benchmarkSort(b, MergeSortParallel)
}
