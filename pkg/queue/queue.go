package queue

type Task struct {
	Id   string
	Body string
}

type Queue interface {
	GetTask() (*Task, error)
	Publish(body *string) (*Task, error)
	FinishTask(id string) error
}

func NewTask(messageID string, messageBody string) *Task {
	return &Task{
		Id:   messageID,
		Body: messageBody,
	}
}
