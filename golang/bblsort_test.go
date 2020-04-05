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

func BubbleSortByReference(a []*Data) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if CompareByReference(a[j], a[j-1]) < 0 {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
}

func BubbleSortByValue(a []Data) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if CompareByValue(a[j], a[j-1]) < 0 {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	const ArraySize = 10000
	master := make([]*Data, ArraySize)
	rand.Seed(82749522)
	for i := 0; i < len(master); i++ {
		master[i] = &Data{}
		master[i].Value = int(rand.Int31())
	}

	b.Run("CallByValue", func(b *testing.B) {
		b.StopTimer()
		data := make([]Data, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = *master[i]
		}
		b.StartTimer()
		BubbleSortByValue(data)
	})

	b.Run("CallByReference", func(b *testing.B) {
		b.StopTimer()
		data := make([]*Data, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = master[i]
		}
		b.StartTimer()
		BubbleSortByReference(data)
	})
}
