package command

import (
	"context"

	"github.com/evgsrkn/go-ddd-example/task/internal/domain"
	"github.com/pkg/errors"
)

type CreateTask struct {
	Id          string
	UserId      string
	Name        string
	Description string
	Status      string
}

type CreateTaskHandler CommandHandler[CreateTask]

type createTaskHandler struct {
	repo domain.WriteRepository
}

func NewCreateTaskHandler(repo domain.WriteRepository) createTaskHandler {
	return createTaskHandler{repo: repo}
}

func (h createTaskHandler) Handle(ctx context.Context, cmd CreateTask) (err error) {
	status, err := domain.StatusFromString(cmd.Status)
	if err != nil {
		return errors.Wrap(err, "cannot create task")
	}

	task, err := domain.NewTask(
		cmd.Id,
		cmd.UserId,
		cmd.Name,
		cmd.Description,
		status,
	)
	if err != nil {
		return err
	}

	return h.repo.CreateTask(ctx, task)
}
