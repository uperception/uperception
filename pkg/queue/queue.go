package queue

type Task struct {
	Id        string
	ProjectID uint
}

type Queue interface {
	GetTask() (*Task, error)
	Publish(id uint) error
	FinishTask(id string) error
}

func NewTask(messageID string, projectID uint) *Task {
	return &Task{
		Id:        messageID,
		ProjectID: projectID,
	}
}
