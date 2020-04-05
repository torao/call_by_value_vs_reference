# call by value v.s. call by reference

Which is faster, a pointer or an entity?

egonspace [report](https://github.com/egonspace/simulations/blob/master/sort/sort_test.go)
that the function calls in golang are faster for passing an entity with a copy than for passing
a pointer. In fact, when I ran the program, I could see that call by value is faster than call
by reference.

```
golang$ go test sort_test.go -v
=== RUN   TestSort
sort time(structure) = 7.834212
sort time(interface) = 25.420997
sort time(*interface) = 24.915953
--- PASS: TestSort (63.83s)
PASS
ok      command-line-arguments  63.842s
```

## Results

```
golang$ go test -bench .
goos: darwin
goarch: amd64
BenchmarkBubbleSort/CallByValue-8         1000000000         0.00722 ns/op
BenchmarkBubbleSort/CallByReference-8     1000000000         0.00116 ns/op
BenchmarkQuickSort/CallByValue-8          1000000000         0.000247 ns/op
BenchmarkQuickSort/CallByReference-8      1000000000         0.000100 ns/op
```

```
clang$ gcc main.c -o main
clang$ ./main
by value    : 1.218750[sec]
by reference: 0.578125[sec]
```