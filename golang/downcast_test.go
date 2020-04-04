package main

import (
	"math/rand"
	"sort"
	"testing"
)

type Abstraction interface {
	Priority() int
}

func (data Data) Priority() int {
	return data.Value
}

func BenchmarkDownCast(b *testing.B) {
	const ArraySize = 10000000
	source := make([]*Data, ArraySize)
	rand.Seed(82749522)
	for i := 0; i < len(source); i++ {
		source[i] = &Data{}
		source[i].Value = int(rand.Int31())
	}

	b.Run("Direct", func(b *testing.B) {
		b.StopTimer()
		data := make([]Data, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = *source[i]
		}
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i].Value < data[j].Value
		})
	})

	b.Run("Pointer", func(b *testing.B) {
		b.StopTimer()
		data := make([]*Data, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = source[i]
		}
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i].Value < data[j].Value
		})
	})

	b.Run("Interface", func(b *testing.B) {
		b.StopTimer()
		data := make([]Abstraction, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = source[i]
		}
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i].Priority() < data[j].Priority()
		})
	})

	b.Run("DownCast", func(b *testing.B) {
		b.StopTimer()
		data := make([]Abstraction, ArraySize)
		for i := 0; i < len(data); i++ {
			data[i] = source[i]
		}
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i].(*Data).Value < data[j].(*Data).Value
		})
	})
}
