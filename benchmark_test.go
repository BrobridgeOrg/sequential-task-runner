package sequential_task_runner

import (
	"testing"
	"time"
)

func Benchmark_Baseline_SingleWorker_BufferSize128(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(1),
		WithMaxPendingCount(128),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Baseline_SingleWorker_BufferSize2048(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(1),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Worker4_BufferSize2048(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_SingleWorker_BufferSize10240(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(1),
		WithMaxPendingCount(10240),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Worker4_BufferSize10240(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(10240),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_SingleWorker_BufferSize2048_SlowHandler(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(1),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			time.Sleep(1 * time.Millisecond)
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Worker4_BufferSize2048_SlowHandler(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			time.Sleep(1 * time.Millisecond)
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Worker8_BufferSize2048_SlowHandler(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			time.Sleep(1 * time.Millisecond)
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}

func Benchmark_Worker4_BufferSize4096_SlowHandler(b *testing.B) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(2048),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			time.Sleep(1 * time.Millisecond)
			return task
		}),
	)
	defer runner.Close()

	output := make(chan string, 10240)
	runner.Subscribe(func(result interface{}) {
		output <- result.(string)
	})

	b.ResetTimer()

	go func() {
		for n := 0; n < b.N; n++ {
			runner.AddTask("BenchmarkContent")
		}
	}()

	for n := 0; n < b.N; n++ {
		<-output
	}
}
