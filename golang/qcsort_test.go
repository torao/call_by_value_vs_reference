package main

import (
	"math/rand"
	"sort"
	"testing"
)

func BenchmarkQuickSort(b *testing.B) {
	const ArraySize = 100000
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
		sort.Slice(data, func(i, j int) bool {
			return CompareByValue(data[i], data[j]) < 0
		})
	})

	b.Run("CallByReference", func(b *testing.B) {
		b.StopTimer()
		data := make([]*Data, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = master[i]
		}
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return CompareByReference(data[i], data[j]) < 0
		})
	})
}
