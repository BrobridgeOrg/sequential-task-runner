package sequential_task_runner

import (
	"errors"
	"sync"
)

type State int64

const (
	StateIdle State = iota
	StateReady
	StateRunning
	StateDone
)

var (
	ErrInvalidHandler = errors.New("str: invalid handler")
	ErrTimeout        = errors.New("str: timeout")
	ErrClosed         = errors.New("str: closed")
)

type Runner struct {
	options      *Options
	pendingCount int
	start        int
	end          int
	pending      chan int
	output       chan interface{}
	controlTable []State
	pendingTasks []interface{}
	fn           func(interface{})

	inputCond  *sync.Cond
	outputCond *sync.Cond
	mutex      sync.Mutex
	isClosed   bool
}

func NewRunner(opts ...Option) *Runner {

	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	r := &Runner{
		options:      options,
		pendingCount: 0,
		start:        -1,
		end:          -1,
		pending:      make(chan int, options.MaxPendingCount),
		output:       make(chan interface{}, options.MaxPendingCount),
		controlTable: make([]State, options.MaxPendingCount),
		pendingTasks: make([]interface{}, options.MaxPendingCount),
		isClosed:     false,
	}

	r.inputCond = sync.NewCond(&r.mutex)
	r.outputCond = sync.NewCond(&r.mutex)

	return r
}

func (r *Runner) produce(task interface{}) {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Wait for a slot to be available
	for r.pendingCount+1 > r.options.MaxPendingCount && !r.isClosed {
		r.inputCond.Wait()
	}

	if r.isClosed {
		return
	}

	// Add task to pending list
	r.end++
	if r.end == r.options.MaxPendingCount {
		r.end = 0
	}

	r.pendingTasks[r.end] = task
	r.controlTable[r.end] = StateReady

	r.pendingCount++

	// push to queue
	r.pending <- r.end
}

func (r *Runner) worker(id int) {

	for taskID := range r.pending {

		// Getting task
		//		r.mutex.Lock()
		r.controlTable[taskID] = StateRunning
		task := r.pendingTasks[taskID]
		//		r.mutex.Unlock()

		// execute task
		result := r.options.WorkerHandler(id, task)

		// Store result
		//		r.mutex.Lock()
		r.controlTable[taskID] = StateDone
		r.pendingTasks[taskID] = result
		//		r.mutex.Unlock()

		//		r.cond.Broadcast()
		r.mutex.Lock()
		r.outputCond.Signal()
		r.mutex.Unlock()
	}
}

func (r *Runner) startWorkers() {

	// Start workers
	for i := 0; i < r.options.WorkerCount; i++ {
		go r.worker(i)
	}
}

func (r *Runner) waitForResults() {

	for !r.isClosed {

		// Next position
		cur := r.start + 1
		if cur == r.options.MaxPendingCount {
			cur = 0
		}

		r.mutex.Lock()

		// Waiting for task to be done
		for r.controlTable[cur] != StateDone && !r.isClosed {
			r.outputCond.Wait()
		}

		if r.isClosed {
			r.mutex.Unlock()
			break
		}

		r.mutex.Unlock()

		r.start = cur

		result := r.pendingTasks[r.start]

		// Remove task from pending list
		r.pendingTasks[r.start] = nil
		r.controlTable[r.start] = StateIdle

		// Publish result
		r.output <- result
	}
}

func (r *Runner) subscribe() {
	for result := range r.output {
		r.fn(result)

		r.mutex.Lock()
		r.pendingCount--
		r.inputCond.Signal()
		r.mutex.Unlock()
	}
}

func (r *Runner) GetPendingCount() int {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.pendingCount
}

func (r *Runner) Subscribe(fn func(interface{})) error {

	if fn == nil {
		return ErrInvalidHandler
	}

	r.fn = fn

	r.startWorkers()

	go r.waitForResults()
	go r.subscribe()

	return nil
}

func (r *Runner) AddTask(task interface{}) error {

	for r.isClosed {
		return ErrClosed
	}

	r.produce(task)

	return nil
}

func (r *Runner) Close() error {

	r.mutex.Lock()
	defer r.mutex.Unlock()

	if r.isClosed {
		return nil
	}

	r.isClosed = true
	close(r.pending)
	close(r.output)

	return nil
}
