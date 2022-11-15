package queue

import (
	"context"
	"log"
	"strconv"

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
	projectID, err := strconv.ParseUint(*output.Messages[0].Body, 10, 32)
	if err != nil {
		log.Fatal("error parsing project ID", err)
	}

	return NewTask(messageId, uint(projectID)), nil
}

func (q *AwsQueue) Publish(id uint) error {
	projectID := strconv.FormatUint(uint64(id), 10)

	_, err := q.client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &q.queueUrl,
		MessageBody: &projectID,
	})

	return err
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
