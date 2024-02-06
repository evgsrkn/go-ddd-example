package domain

import "github.com/pkg/errors"

var (
	StatusTodo       = Status{"To Do"}
	StatusInProgress = Status{"In Progress"}
	StatusDone       = Status{"Done"}
)

var TaskStatuses = []Status{
	StatusTodo,
	StatusInProgress,
	StatusDone,
}

type Status struct {
	a string
}

func StatusFromString(status string) (Status, error) {
	for _, s := range TaskStatuses {
		if status == s.String() {
			return s, nil
		}
	}

	return Status{}, errors.New("status '%s' does not exist")
}

func (s Status) String() string {
	return s.a
}

func (s Status) IsZero() bool {
	return s == Status{}
}

func (t Task) IsInTodo() bool {
	return t.status == StatusTodo
}

func (t Task) IsInProgress() bool {
	return t.status == StatusInProgress
}

func (t Task) IsDone() bool {
	return t.status == StatusDone
}

func (t *Task) Start() error {
	if !t.IsInTodo() {
		return errors.Errorf("task '%s' is already started", t.id)
	}

	t.status = StatusInProgress
	return nil
}

func (t *Task) Complete() error {
	if t.IsDone() {
		return errors.Errorf("task '%s' is already completed", t.id)
	}

	t.status = StatusDone
	return nil
}

func (t *Task) Cancel() error {
	if !t.IsInProgress() {
		return errors.Errorf("cannot cancel task '%s' it is not in progress", t.id)
	}

	t.status = StatusTodo
	return nil
}

func (t *Task) Reject() error {
	if !t.IsDone() {
		return errors.Errorf("cannot reject unfinished task '%s'", t.id)
	}

	t.status = StatusInProgress
	return nil
}
