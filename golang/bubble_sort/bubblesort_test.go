package main

import (
	"math/rand"
	"testing"
)

const PayloadSize = 100

type Data struct {
	Value   int
	Payload [PayloadSize]uint8
}

func CompareByReference(a *Data, b *Data) int {
	if a.Value < b.Value {
		return -1
	}
	if a.Value > b.Value {
		return 1
	}
	return 0
}

func CompareByValue(a Data, b Data) int {
	if a.Value < b.Value {
		return -1
	}
	if a.Value > b.Value {
		return 1
	}
	return 0
}

func BubbleSort(a []*Data, lessThan func(i, j int) bool) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if lessThan(j, j-1) {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	const ArraySize = 10000
	source := make([]*Data, ArraySize)
	rand.Seed(82749522)
	for i := 0; i < len(source); i++ {
		source[i] = &Data{}
		source[i].Value = int(rand.Int31())
	}

	b.Run("CallByValue", func(b *testing.B) {
		arr := make([]*Data, ArraySize)
		copy(arr, source)
		b.ResetTimer()
		BubbleSort(arr, func(i, j int) bool {
			return CompareByValue(*arr[i], *arr[j]) < 0
		})
	})

	b.Run("CallByReference", func(b *testing.B) {
		arr := make([]*Data, ArraySize)
		copy(arr, source)
		b.ResetTimer()
		BubbleSort(arr, func(i, j int) bool {
			return CompareByReference(arr[i], arr[j]) < 0
		})
	})
}
