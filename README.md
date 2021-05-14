# Benchmark results

## Without sync.Pool 

Environment Linux, cpu i3

```
go test -bench=BenchmarkWithoutPool

goos: linux
goarch: amd64
pkg: go-benchmarks
cpu: Intel(R) Core(TM) i3-6006U CPU @ 2.00GHz
BenchmarkWithoutPool-4   	     315	   4368705 ns/op
PASS
ok  	go-benchmarks	137.583s
```

Environment Linux, cpu i5

```
go test -bench=BenchmarkWithoutPool

goos: linux
goarch: amd64
pkg: go-benchmarks
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkWithoutPool-8   	     603	   3078665 ns/op
PASS
ok  	go-benchmarks	151.674s
```


Environment MacOS

```

```

## With sync.Pool

Environment Linux, cpu i3

```
go test -bench=BenchmarkWithPool

goos: linux
goarch: amd64
pkg: rebrain/module05/pool
cpu: Intel(R) Core(TM) i3-6006U CPU @ 2.00GHz
BenchmarkWithPool-4   	     206	   5170697 ns/op
PASS
ok  	go-benchmarks	94.455s
```

Environment Linux, cpu i5

```
go test -bench=BenchmarkWithPool

goos: linux
goarch: amd64
pkg: go-benchmarks
cpu: Intel(R) Core(TM) i5-10210U CPU @ 1.60GHz
BenchmarkWithPool-8   	     374	   4880955 ns/op
PASS
ok  	go-benchmarks	100.269s
```

Environment MacOS

```

```
