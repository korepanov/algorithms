package sort

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

type ints []int

func (i ints) Iter() func() (int, bool) {
	index := 0
	return func() (val int, ok bool) {
		if index >= len(i) {
			return
		}
		val, ok = i[index], true
		index++
		return
	}
}

func basicTestSort(t *testing.T, testFunc func(arr []int)) {
	arr := []int{5, 4, 2, 1, 2}
	expArr := []int{1, 2, 2, 4, 5}
	testFunc(arr)
	if !reflect.DeepEqual(arr, expArr) {
		t.Log(fmt.Sprintf("expect:%v", expArr) + fmt.Sprintf("but get:%v", arr))
		t.Fail()
	}
}

func testSort(t *testing.T, testFunc func(arr []int)) {
	arrSize := rand.Intn(100)
	arr := make([]int, arrSize, arrSize)
	expArr := make([]int, arrSize, arrSize)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	copy(expArr, arr)
	insertionSort(expArr)
	testFunc(arr)
	if !reflect.DeepEqual(arr, expArr) {
		t.Log(fmt.Sprintf("expect:\n%v\n", expArr) + fmt.Sprintf("but get:\n%v\n", arr))
		t.Fail()
	}
}

func benchmarkSort(b *testing.B, testFunc func(arr []int)) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arrSize := 10000
		arr := make([]int, arrSize, arrSize)
		for j := range arr {
			arr[j] = rand.Intn(10000) + 1
		}
		b.StartTimer()
		testFunc(arr)
	}
}

func BenchmarkSortWithSize(b *testing.B, testFunc func(arr []int), arrSize int) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := make([]int, arrSize, arrSize)
		for j := range arr {
			arr[j] = rand.Intn(arrSize) + 1
		}
		b.StartTimer()
		testFunc(arr)
	}
}
