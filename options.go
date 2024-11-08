package sequential_task_runner

type WorkerHandler func(workerID int, task interface{}) interface{}

type Option func(*Options)

type Options struct {
	WorkerCount     int
	MaxPendingCount int
	WorkerHandler   WorkerHandler
}

func WithWorkerCount(workerCount int) Option {
	return func(o *Options) {
		o.WorkerCount = workerCount
	}
}

func WithMaxPendingCount(maxPendingCount int) Option {
	return func(o *Options) {
		o.MaxPendingCount = maxPendingCount
	}
}

func WithWorkerHandler(handler WorkerHandler) Option {
	return func(o *Options) {
		o.WorkerHandler = handler
	}
}
