# sequential-task-runner

High performance task runner with support for sequential and parallel execution.

## Benchmark

Here is the benchmark result:

```
go test -bench . -benchmem -cpu=4
goos: darwin
goarch: amd64
pkg: github.com/BrobridgeOrg/sequential-task-runner
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
Benchmark_Baseline_SingleWorker_BufferSize128-4       	 3330404	       359.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_SingleWorker_BufferSize2048-4      	 3028326	       411.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize2048-4                    	 2455048	       485.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_SingleWorker_BufferSize10240-4              	 3149592	       390.4 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize10240-4                   	 2526714	       474.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_SingleWorker_BufferSize2048_SlowHandler-4   	    1054	   1139250 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize2048_SlowHandler-4        	    4285	    283941 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker8_BufferSize2048_SlowHandler-4        	    8194	    142029 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker16_BufferSize2048_SlowHandler-4       	   16828	     71124 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize4096_SlowHandler-4        	    4290	    284935 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/BrobridgeOrg/sequential-task-runner	28.183s
```

## License

Licensed under the Apache License

## Authors

Copyright(c) 2024 Fred Chien <fred@brobridge.com>
