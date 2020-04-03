package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

type Node struct {
	P    int64
	Data float64
	Big [100]byte // to increase the cost of passing it as parameter
}

func (n Node) LessThanN(other Node) bool {
	return n.Data < other.Data
}

func (n Node) Priority() int64 {
	return n.P
}

func (n Node) LessThan(other SortItem) bool {
	return n.Data < other.(Node).Data
}

func (n Node) LessThanP(other *SortItemP) bool {
	return n.Data < (*other).(Node).Data
}

type SortItem interface {
	Priority() int64
	LessThan(other SortItem) bool
}

type SortItemP interface {
	Priority() int64
	LessThanP(other *SortItemP) bool
}

func NewItems(count int) ([]Node, []SortItem, []SortItemP) {
	resultN := make([]Node, count)
	result := make([]SortItem, count)
	resultP := make([]SortItemP, count)
	for i := 0; i < count; i++ {
		resultN[i] = Node{
			P: 100,
			Data: rand.Float64(),
		}
		result[i] = resultN[i]
		resultP[i] = resultN[i]
	}
	return resultN, result, resultP
}

func TestSort(t *testing.T) {
	nodesN, nodes, nodesP := NewItems(10000000)

	var start, end time.Time
	start = time.Now()
	sort.Slice(nodesN, func(i, j int) bool {
		if nodesN[i].Priority() != nodesN[j].Priority() {
			return nodesN[i].Priority() < nodesN[j].Priority()
		}
		return nodesN[i].LessThanN(nodesN[j])
	})
	end = time.Now()
	fmt.Printf("sort time(structure) = %f\n", end.Sub(start).Seconds())

	start = time.Now()
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Priority() != nodes[j].Priority() {
			return nodes[i].Priority() < nodes[j].Priority()
		}
		return nodes[i].LessThan(nodes[j])
	})

	end = time.Now()
	fmt.Printf("sort time(interface) = %f\n", end.Sub(start).Seconds())

	start = time.Now()
	sort.Slice(nodesP, func(i, j int) bool {
		if nodesP[i].Priority() != nodesP[j].Priority() {
			return nodesP[i].Priority() < nodesP[j].Priority()
		}
		return nodesP[i].LessThanP(&nodesP[j])
	})
	end = time.Now()
	fmt.Printf("sort time(*interface) = %f\n", end.Sub(start).Seconds())
}