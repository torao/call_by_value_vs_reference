# The slowness of downcast, call by value and reference in Golang

In Golang:

1. It's faster to use a pointer than to pass in an entity copy when using a structure as a parameter of a function.
2. The downcast of the pointer has an overhead.
3. Calling methods through an interface is quite expensive.

All of these may be similar to C/C++ and Java, but the impact on 3 seems to be a bit larger, and if you adopt Go for performance purposes, it may be difficult to reconcile the cost with the design of the object orientation and responsibility dividing.

```
sort$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkQuickSort/Entity-8         1000000000         0.182 ns/op
BenchmarkQuickSort/Image-8          1000000000         0.644 ns/op
BenchmarkQuickSort/ImageRef-8       1000000000         0.631 ns/op
```

## Results

[downcast_test.go](golang/downcast_test.go)

```
golang$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkEntityField-8                   2     571069256 ns/op
BenchmarkRefField-8                      2     597430914 ns/op
BenchmarkEntityDownCastField-8           1   46943004589 ns/op
BenchmarkRefDownCastField-8              1    1113006051 ns/op
BenchmarkEntityMethod-8                  1   12332693902 ns/op
BenchmarkRefMethod-8                     2     574504750 ns/op
BenchmarkEntityDownCastMethod-8          1   50146782632 ns/op
BenchmarkRefDownCastMethod-8             1    1186318261 ns/op
BenchmarkEntityNoCastMethod-8            1  195731203442 ns/op
BenchmarkRefNoCastMethod-8               1   27091621075 ns/op
```

[bubble_sort.go](golang/bubble_sort/bubblesort_test.go)

```
bubble_sort$ $ go test -bench .
goos: darwin
goarch: amd64
BenchmarkBubbleSort/CallByValue-8         1000000000         0.796 ns/op
BenchmarkBubbleSort/CallByReference-8     1000000000         0.417 ns/op
```

[clang/main.c](clang/main.c)

```
clang$ gcc main.c -o main
clang$ ./main
by value    : 2.196516[sec]
by reference: 0.461386[sec]
```