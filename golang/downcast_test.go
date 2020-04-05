package main

import (
	"math/rand"
	"sort"
	"testing"
)

type Image interface {
	Priority() int
	LessThanByDownCastField(other Image) bool
	LessThanByDownCastMethod(other Image) bool
}

func (data Data) Priority() int {
	return data.Value
}

func (data Data) LessThanByEntityField(other Data) bool {
	return data.Value < other.Value
}

func (data Data) LessThanByReferenceField(other *Data) bool {
	return data.Value < other.Value
}

func (data Data) LessThanByDownCast1Field(other Image) bool {
	return data.Value < other.(Data).Value
}

func (data Data) LessThanByDownCast2Field(other Image) bool {
	return data.Value < other.(*Data).Value
}

func (data Data) LessThanByEntityMethod(other Data) bool {
	return data.Priority() < other.Priority()
}

func (data Data) LessThanByReferenceMethod(other *Data) bool {
	return data.Priority() < other.Priority()
}

func (data Data) LessThanByDownCast1Method(other Image) bool {
	return data.Priority() < other.(Data).Priority()
}

func (data Data) LessThanByDownCast2Method(other Image) bool {
	return data.Priority() < other.(*Data).Priority()
}

const seed = 82749522
const arraySize = 10000000

func NewDataArray() []Data {
	arr := make([]Data, arraySize)
	rand.Seed(seed)
	for i := 0; i<arraySize; i++ {
		arr[i] = Data{}
		arr[i].Value = int(rand.Int31())
	}
	return arr
}

func NewImageArray() []Image {
	arr := make([]Image, arraySize)
	rand.Seed(seed)
	for i := 0; i<arraySize; i++ {
		data := Data{}
		data.Value = int(rand.Int31())
		arr[i] = data
	}
	return arr
}

func BenchmarkEntityField(b *testing.B) {
	arr := NewDataArray()
	b.ResetTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByEntityField(arr[j])})
}

func BenchmarkReferenceField(b *testing.B) {
	arr := NewDataArray()
	b.StartTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByReferenceField(&arr[j])})
}

func BenchmarkDownCastField(b *testing.B){
	arr := NewImageArray()
	b.ResetTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByDownCastField(arr[j])})
}

func BenchmarkEntityMethod(b *testing.B) {
	arr := NewDataArray()
	b.ResetTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByEntityMethod(arr[j])})
}

func BenchmarkReferenceMethod(b *testing.B) {
	arr := NewDataArray()
	b.StartTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByReferenceMethod(&arr[j])})
}

func BenchmarkDownCastMethod(b *testing.B){
	arr := NewImageArray()
	b.ResetTimer()
	sort.Slice(arr, func(i, j int) bool { return arr[i].LessThanByDownCastMethod(arr[j])})
}
