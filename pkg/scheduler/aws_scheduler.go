package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/leometzger/mmonitoring/pkg/models"
)

type AwsScheduler struct {
	client *eventbridge.Client
}

func NewCloudwatchScheduler(client *eventbridge.Client) *AwsScheduler {
	return &AwsScheduler{client: client}
}

func (s *AwsScheduler) Schedule(id string, schedule models.Schedule) error {
	ctx := context.Background()
	expression := s.fromScheduleToCronExpression(schedule)

	_, err := s.client.PutRule(ctx, &eventbridge.PutRuleInput{
		Name:               &id,
		ScheduleExpression: &expression,
	})

	return err
}

func (s *AwsScheduler) EnableSchedule(id string) error {
	ctx := context.Background()
	_, err := s.client.EnableRule(ctx, &eventbridge.EnableRuleInput{
		Name: &id,
	})

	return err
}

func (s *AwsScheduler) DisableSchedule(id string) error {
	ctx := context.Background()
	_, err := s.client.DisableRule(ctx, &eventbridge.DisableRuleInput{
		Name: &id,
	})

	return err
}

func (s *AwsScheduler) fromScheduleToCronExpression(schedule models.Schedule) string {
	// default fake expression
	return "cron(15 12 * * ? *)"
}
