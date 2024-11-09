package sequential_task_runner

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunner_AddTask(t *testing.T) {

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(100),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	// Add a task to the runner
	for i := 0; i < 100; i++ {
		runner.AddTask(i)
	}

	assert.Equal(t, runner.GetPendingCount(), 100)
}

func TestRunner_Workers(t *testing.T) {

	var wg sync.WaitGroup

	count := int64(0)

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(100),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			result := atomic.AddInt64(&count, 1)

			wg.Done()

			return result
		}),
	)
	defer runner.Close()

	// Start workers
	runner.startWorkers()

	// Add a task to the runner
	for i := 0; i < 100; i++ {
		wg.Add(1)
		runner.AddTask(i)
	}

	assert.Equal(t, runner.GetPendingCount(), 100)

	wg.Wait()

	for i := 0; i < 100; i++ {
		//state := runner.controlTable[i]
		//assert.Equal(t, StateDone, state)
		assert.Equal(t, StateDone, runner.tasks[i].State)
	}
}

func TestRunner_Subscribe(t *testing.T) {

	targetTaskCount := 5000000

	var wg sync.WaitGroup

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(1024),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task.(int) + 1
		}),
	)
	defer runner.Close()

	expected := 0
	err := runner.Subscribe(func(result interface{}) {
		expected++
		assert.Equal(t, expected, result)
		if expected%100000 == 0 {
			t.Logf("Task %d finished", result)
		}
		wg.Done()
	})
	if !assert.Nil(t, err) {
		return
	}

	// Add a task to the runner
	for i := 0; i < targetTaskCount; i++ {
		wg.Add(1)
		runner.AddTask(i)
	}

	wg.Wait()

	assert.Equal(t, runner.GetPendingCount(), 0)
	assert.Equal(t, targetTaskCount, expected)
}

func TestRunner_UnsteadyHandler(t *testing.T) {

	targetTaskCount := 200

	var wg sync.WaitGroup

	processed := int64(0)

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(100),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			atomic.AddInt64(&processed, 1)
			return task
		}),
	)
	defer runner.Close()

	expected := 0
	err := runner.Subscribe(func(result interface{}) {
		assert.Equal(t, expected, result)
		expected++
		t.Logf("Task %d finished", result.(int)+1)
		wg.Done()
	})
	if !assert.Nil(t, err) {
		return
	}

	// Add a task to the runner
	for i := 0; i < targetTaskCount; i++ {
		wg.Add(1)
		runner.AddTask(i)
	}

	wg.Wait()

	assert.Equal(t, runner.GetPendingCount(), 0)
	assert.Equal(t, int64(targetTaskCount), processed)
	assert.Equal(t, targetTaskCount, expected)
}

func TestRunner_SlowOutput(t *testing.T) {

	targetTaskCount := 200

	var wg sync.WaitGroup

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(10),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			return task
		}),
	)
	defer runner.Close()

	expected := 0
	err := runner.Subscribe(func(result interface{}) {
		// Sequential output
		assert.Equal(t, expected, result)
		expected++
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		wg.Done()
	})
	if !assert.Nil(t, err) {
		return
	}

	// Add a task to the runner
	for i := 0; i < targetTaskCount; i++ {
		wg.Add(1)
		runner.AddTask(i)
	}

	wg.Wait()

	assert.Equal(t, runner.GetPendingCount(), 0)
	assert.Equal(t, targetTaskCount, expected)
}

func TestRunner_UnsteadyHandler_And_SlowOutput(t *testing.T) {

	targetTaskCount := 4000

	var wg sync.WaitGroup

	rand.Seed(time.Now().UnixNano())

	// Create a new runner
	runner := NewRunner(
		WithWorkerCount(4),
		WithMaxPendingCount(10),
		WithWorkerHandler(func(workerID int, task interface{}) interface{} {
			time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
			return task
		}),
	)
	defer runner.Close()

	expected := 0
	err := runner.Subscribe(func(result interface{}) {
		// Sequential output
		assert.Equal(t, expected, result)
		expected++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		wg.Done()
	})
	if !assert.Nil(t, err) {
		return
	}

	// Add a task to the runner
	for i := 0; i < targetTaskCount; i++ {
		wg.Add(1)
		runner.AddTask(i)
	}

	wg.Wait()

	assert.Equal(t, runner.GetPendingCount(), 0)
	assert.Equal(t, targetTaskCount, expected)
}
