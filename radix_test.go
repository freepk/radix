package radix

import (
	"math/rand"
	"sort"
	"testing"
)

const (
	benchArraySize = 100000
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

func TestRadixNotSwaped(t *testing.T) {
	source := []int{14878831, 15503653, 3207093}
	buffer := make([]int, 3)
	Ints(source, buffer, 3)
	if !sort.IntsAreSorted(source) {
		t.Fail()
	}
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

func BenchmarkRadixSortRnd(b *testing.B) {
	b.StopTimer()
	original := testArrayRnd(benchArraySize)
	buffer := make([]int, benchArraySize)
	source := make([]int, benchArraySize)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		Ints(source, buffer, benchArraySize)
	}
}

func BenchmarkStandartSortRnd(b *testing.B) {
	b.StopTimer()
	original := testArrayRnd(benchArraySize)
	source := make([]int, benchArraySize)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		sort.Ints(source)
	}
}

func BenchmarkRadixSortSeq(b *testing.B) {
	b.StopTimer()
	original := testArraySeq(benchArraySize)
	buffer := make([]int, benchArraySize)
	source := make([]int, benchArraySize)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		Ints(source, buffer, benchArraySize)
	}
}

func BenchmarkStandartSortSeq(b *testing.B) {
	b.StopTimer()
	original := testArraySeq(benchArraySize)
	source := make([]int, benchArraySize)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		copy(source, original)
		sort.Ints(source)
	}
}
