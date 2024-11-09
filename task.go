package sequential_task_runner

const (
	StateIdle State = iota
	StateReady
	StateRunning
	StateDone
)

type Task struct {
	State State
	Data  interface{}
}

func (t *Task) Reset() {
	t.Data = nil
	t.State = StateIdle
}

func (t *Task) SetState(state State) {
	t.State = state
}

func (t *Task) Update(state State, data interface{}) {
	t.Data = data
	t.State = state
}

func (t *Task) Drain() interface{} {
	defer t.Reset()
	return t.Data
}
