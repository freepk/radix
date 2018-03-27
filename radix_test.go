package radix

import (
	"math/rand"
	"sort"
	"testing"
)

func testArraySeq(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	return data
}

func testArrayRnd(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Int()
	}
	return data
}

func TestRadixSort(t *testing.T) {
	iter := 1024
	size := 1024
	buffer := make([]int, size)
	for i := 0; i < iter; i++ {
		source := testArrayRnd(size)
		Ints(source, buffer, size)
		if !sort.IntsAreSorted(source) {
			t.Fail()
		}
	}
}

func BenchmarkRadixSortRnd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	original := testArrayRnd(size)
	buffer := make([]int, size)
	source := make([]int, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		Ints(source, buffer, size)
	}
}

func BenchmarkStandartSortRnd100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	original := testArrayRnd(size)
	source := make([]int, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		sort.Ints(source)
	}
}

func BenchmarkRadixSortSeq100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	original := testArraySeq(size)
	buffer := make([]int, size)
	source := make([]int, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		Ints(source, buffer, size)
	}
}

func BenchmarkStandartSortSeq100000(b *testing.B) {
	b.StopTimer()
	size := 100000
	original := testArraySeq(size)
	source := make([]int, size)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		sort.Ints(source)
	}
}
