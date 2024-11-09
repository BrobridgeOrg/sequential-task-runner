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
Benchmark_Baseline_SingleWorker_BufferSize128-4       	 3502114	       340.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_SingleWorker_BufferSize2048-4      	 3495802	       338.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize2048-4                    	 2910336	       419.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_SingleWorker_BufferSize10240-4              	 3402405	       346.1 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize10240-4                   	 2834820	       404.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_SingleWorker_BufferSize2048_SlowHandler-4   	    1050	   1146070 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize2048_SlowHandler-4        	    4275	    285961 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker8_BufferSize2048_SlowHandler-4        	    7954	    143237 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker16_BufferSize2048_SlowHandler-4       	   16732	     71724 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize4096_SlowHandler-4        	    4270	    286049 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/BrobridgeOrg/sequential-task-runner	33.142s
```

## License

Licensed under the Apache License

## Authors

Copyright(c) 2024 Fred Chien <fred@brobridge.com>
