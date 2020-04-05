package main

import (
	"math"
	"testing"
)

type Entity struct {
	Value   int
	Payload [100]uint8
}

type Image interface {
	Priority() int
}

func (entity Entity) Priority() int {
	return entity.Value
}

func BenchmarkEntityField(b *testing.B) {
	lessThan := func(a, b Entity) bool {
		return a.Value < b.Value
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkRefField(b *testing.B) {
	lessThan := func(a, b *Entity) bool {
		return a.Value < b.Value
	}
	x1, x2 := &Entity{}, &Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkEntityDownCastField(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.(Entity).Value < b.(Entity).Value
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkRefDownCastField(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.(*Entity).Value < b.(*Entity).Value
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(&x1, &x2)
		}
	}
}

func BenchmarkEntityMethod(b *testing.B) {
	lessThan := func(a, b Entity) bool {
		return a.Priority() < b.Priority()
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkRefMethod(b *testing.B) {
	lessThan := func(a, b *Entity) bool {
		return a.Priority() < b.Priority()
	}
	x1, x2 := &Entity{}, &Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkEntityDownCastMethod(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.(Entity).Priority() < b.(Entity).Priority()
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkRefDownCastMethod(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.(*Entity).Priority() < b.(*Entity).Priority()
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(&x1, &x2)
		}
	}
}

func BenchmarkEntityNoCastMethod(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.Priority() < b.Priority()
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(x1, x2)
		}
	}
}

func BenchmarkRefNoCastMethod(b *testing.B) {
	lessThan := func(a, b Image) bool {
		return a.Priority() < b.Priority()
	}
	x1, x2 := Entity{}, Entity{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < math.MaxInt32; j++ {
			lessThan(&x1, &x2)
		}
	}
}
