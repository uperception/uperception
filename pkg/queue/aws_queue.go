package queue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type AwsQueue struct {
	queueArn string
	queueUrl string
	client   *sqs.Client
}

func NewAwsQueue(client *sqs.Client, arn string, url string) *AwsQueue {
	return &AwsQueue{
		client:   client,
		queueArn: arn,
		queueUrl: url,
	}
}

func (q *AwsQueue) GetTask() (*Task, error) {
	output, err := q.client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
		QueueUrl:            &q.queueUrl,
		MaxNumberOfMessages: 1,
	})

	if err != nil {
		return nil, err
	}

	if len(output.Messages) == 0 {
		return nil, nil
	}

	messageId := *output.Messages[0].MessageId
	messageBody := *output.Messages[0].Body

	return NewTask(messageId, messageBody), nil
}

func (q *AwsQueue) Publish(body *string) (*Task, error) {
	output, err := q.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &q.queueUrl,
		MessageBody: body,
	})

	if err != nil {
		return nil, err
	}

	return NewTask(*output.MessageId, *body), nil
}

func (q *AwsQueue) FinishTask(id string) error {
	_, err := q.client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
		QueueUrl: &q.queueUrl,
	})

	if err != nil {
		return err
	}

	return nil
}
