package queue

import "github.com/aws/aws-sdk-go-v2/service/sqs"

type AwsQueue struct {
	queueArn string
	client   *sqs.Client
}

func NewAwsQueue(client *sqs.Client, arn string) *AwsQueue {
	return &AwsQueue{
		client:   client,
		queueArn: arn,
	}
}

func (q *AwsQueue) GetTask() (*Task, error) {
	return nil, nil
}

func (q *AwsQueue) FinishTask(id string) error {
	return nil
}
