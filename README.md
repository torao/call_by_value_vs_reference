# call by value v.s. call by reference

Which is faster, a pointer or an entity?

[egonspace's report](https://github.com/egonspace/simulations/blob/master/sort/sort_test.go) shows that the
function calls in golang are faster for passing an entity with a copy than for passing a pointer. In fact,
when I ran the program, I could see that call by value is faster than call by reference.

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
golang$ go run main.go
by value    : 0.874962[sec]
by reference: 0.221388[sec]
```

```
clang$ gcc main.c -o main
clang$ ./main
by value    : 1.218750[sec]
by reference: 0.578125[sec]
```