package sort

import (
	"math/rand"
	"sort"
	"testing"
)

type Entity struct {
	Value   int
	Payload [100]byte
}

type Image interface {
	LessThanByImage(other Image) bool
	LessThanByImageRef(other *Image) bool
}

func (n Entity) LessThanByEntity(other Entity) bool {
	return n.Value < other.Value
}

func (n Entity) LessThanByImage(other Image) bool {
	return n.Value < other.(Entity).Value
}

func (n Entity) LessThanByImageRef(other *Image) bool {
	return n.Value < (*other).(Entity).Value
}

func BenchmarkQuickSort(b *testing.B) {
	count := 800000
	nodesE := make([]Entity, count)
	nodes1 := make([]Image, count)
	nodes2 := make([]Image, count)
	for i := 0; i < count; i++ {
		nodesE[i] = Entity{Value: int(rand.Int31())}
		nodes1[i] = nodesE[i]
		nodes2[i] = nodesE[i]
	}

	b.Run("Entity", func(b *testing.B) {
		sort.Slice(nodesE, func(i, j int) bool {
			return nodesE[i].LessThanByEntity(nodesE[j])
		})
	})

	b.Run("Image", func(b *testing.B) {
		sort.Slice(nodes1, func(i, j int) bool {
			return nodes1[i].LessThanByImage(nodes1[j])
		})
	})

	b.Run("ImageRef", func(b *testing.B) {
		sort.Slice(nodes2, func(i, j int) bool {
			return nodes2[i].LessThanByImageRef(&nodes2[j])
		})
	})
}
