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
Benchmark_Baseline_SingleWorker_BufferSize128-4                	 3202452	       392.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_SingleWorker_BufferSize2048-4               	 2919516	       413.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize2048-4                             	 2442498	       490.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_SingleWorker_BufferSize10240-4                       	 2763142	       441.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_Worker4_BufferSize10240-4                            	 2199294	       536.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_SingleWorker_BufferSize2048_SlowHandler-4   	    1060	   1141427 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_Worker4_BufferSize2048_SlowHandler-4        	    4213	    285407 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_Worker8_BufferSize2048_SlowHandler-4        	    4310	    283628 ns/op	       0 B/op	       0 allocs/op
Benchmark_Baseline_Worker4_BufferSize4096_SlowHandler-4        	    4225	    285766 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/BrobridgeOrg/sequential-task-runner	17.067s
```

## License

Licensed under the Apache License

## Authors

Copyright(c) 2024 Fred Chien <fred@brobridge.com>
