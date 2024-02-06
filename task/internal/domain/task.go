package domain

import "github.com/pkg/errors"

type Task struct {
	id          string
	userId      string
	name        string
	description string
	status      Status
}

func NewTask(id, userId, name, description string, status Status) (Task, error) {
	if id == "" {
		return Task{}, errors.New("id is empty")
	}
	if name == "" {
		return Task{}, errors.New("name is empty")
	}
	if description == "" {
		return Task{}, errors.New("description is empty")
	}
	if status.IsZero() {
		return Task{}, errors.New("invalid status")
	}

	return Task{
		id:          id,
		userId:      userId,
		name:        name,
		description: description,
		status:      status,
	}, nil
}

func (t *Task) Assign(userId string) error {
	if t.userId == userId {
		return errors.Errorf("cannot reassign same user '%s' to taks '%s'", userId, t.id)
	}

	t.userId = userId

	return nil
}

func (t *Task) Rename(name string) error {
	if t.name == name {
		return errors.Errorf("cannot update task '%s' with same name '%s'", t.id, name)
	}

	t.name = name

	return nil
}

func (t *Task) ChangeDescription(description string) error {
	if t.name == description {
		return errors.Errorf("cannot update task '%s' with same description '%s'", t.id, description)
	}

	t.description = description

	return nil
}
