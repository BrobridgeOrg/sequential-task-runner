package sequential_task_runner

import (
	"errors"
	"sync"
)

type State int64

var (
	ErrInvalidHandler = errors.New("str: invalid handler")
	ErrClosed         = errors.New("str: closed")
)

type Runner struct {
	options      *Options
	pendingCount int
	end          int
	pending      chan int
	output       chan interface{}
	tasks        []*Task
	fn           func(interface{})

	inputCond   *sync.Cond
	outputCond  *sync.Cond
	mutex       sync.Mutex
	outputMutex sync.Mutex
	isClosed    bool
}

func NewRunner(opts ...Option) *Runner {

	options := &Options{}

	for _, opt := range opts {
		opt(options)
	}

	r := &Runner{
		options:      options,
		pendingCount: 0,
		end:          -1,
		tasks:        make([]*Task, options.MaxPendingCount),
		pending:      make(chan int, options.MaxPendingCount),
		output:       make(chan interface{}, options.MaxPendingCount),
		isClosed:     false,
	}

	for i := 0; i < options.MaxPendingCount; i++ {
		r.tasks[i] = &Task{}
	}

	r.inputCond = sync.NewCond(&r.mutex)
	r.outputCond = sync.NewCond(&r.outputMutex)

	return r
}

func (r *Runner) updateLastPosition() int {

	r.outputCond.L.Lock()
	defer r.outputCond.L.Unlock()

	end := r.end + 1
	if end == r.options.MaxPendingCount {
		end = 0
	}

	//fmt.Println("updateLastPosition", end)

	r.end = end

	// Notify for the end position change
	r.outputCond.Signal()

	return end
}

func (r *Runner) produce(data interface{}) {

	r.inputCond.L.Lock()
	defer r.inputCond.L.Unlock()

	// Wait for a slot to be available
	for r.pendingCount == r.options.MaxPendingCount && !r.isClosed {
		r.inputCond.Wait()
	}

	if r.isClosed {
		return
	}

	r.pendingCount++

	// Add task to pending list
	end := r.updateLastPosition()

	r.tasks[end].Update(StateReady, data)

	//	fmt.Println("produce", r.end, data)

	// push to queue
	r.pending <- end
}

func (r *Runner) worker(id int) {

	for taskID := range r.pending {

		// Getting task
		task := r.tasks[taskID]
		task.SetState(StateRunning)

		data := task.Data

		// execute task
		//fmt.Println("worker", id, "task", taskID)
		result := r.options.WorkerHandler(id, data)

		// Store result
		task.Update(StateDone, result)

		r.outputCond.Signal()
	}
}

func (r *Runner) startWorkers() {

	// Start workers
	for i := 0; i < r.options.WorkerCount; i++ {
		go r.worker(i)
	}
}

func (r *Runner) waitForResults() {

	go r.subscribe()

	r.outputCond.L.Lock()
	defer r.outputCond.L.Unlock()

	cur := -1

	for !r.isClosed {

		// Next position
		nextPos := cur + 1
		if nextPos == r.options.MaxPendingCount {
			nextPos = 0
		}

		// No more results
		for r.tasks[nextPos].State != StateDone && !r.isClosed {
			r.outputCond.Wait()
		}

		if r.isClosed {
			break
		}

		cur = nextPos

		task := r.tasks[cur]

		result := task.Data
		//fmt.Println("waitForResults reset", "cur", cur, "end", r.end, "state", task.State, StateDone)

		// Clear
		task.Reset()

		// Publish result
		r.output <- result
	}
}

func (r *Runner) subscribe() {
	for result := range r.output {
		r.fn(result)

		r.inputCond.L.Lock()
		r.pendingCount--
		r.inputCond.L.Unlock()

		r.inputCond.Signal()
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
