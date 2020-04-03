package main

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 100

type Data struct {
	value   int
	payload [SIZE]uint8
}

func compare_by_reference(a *Data, b *Data) int {
	if (*a).value < (*b).value {
		return -1
	}
	if (*a).value > (*b).value {
		return 1
	}
	return 0
}

func compare_by_value(a Data, b Data) int {
	if a.value < b.value {
		return -1
	}
	if a.value > b.value {
		return 1
	}
	return 0
}

func bblsort_by_reference(a []*Data) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if compare_by_reference(a[j], a[j-1]) < 0 {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
}

func bblsort_by_value(a []Data) {
	l := len(a)
	for i := 0; i < l; i++ {
		for j := 1; j < l-i; j++ {
			if compare_by_value(a[j], a[j-1]) < 0 {
				tmp := a[j]
				a[j] = a[j-1]
				a[j-1] = tmp
			}
		}
	}
}

const LEN = 10000

func main() {
	data1 := make([]Data, LEN)
	data2 := make([]*Data, LEN)
	rand.Seed(82749522)
	for i := 0; i < len(data1); i++ {
		data2[i] = &Data{}
		data2[i].value = int(rand.Int31())
		data1[i] = *data2[i]
	}

	t0 := time.Now()
	bblsort_by_value(data1)
	t1 := time.Now()
	fmt.Printf("by value    : %f[sec]\n", t1.Sub(t0).Seconds())

	t2 := time.Now()
	bblsort_by_reference(data2)
	t3 := time.Now()
	fmt.Printf("by reference: %f[sec]\n", t3.Sub(t2).Seconds())
}
