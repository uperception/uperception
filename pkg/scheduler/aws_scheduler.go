package scheduler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/eventbridge"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	"github.com/leometzger/mmonitoring/pkg/models"
)

type AwsScheduler struct {
	client      *eventbridge.Client
	targetQueue string
}

func NewCloudwatchScheduler(client *eventbridge.Client, targetQueue string) *AwsScheduler {
	return &AwsScheduler{
		client:      client,
		targetQueue: targetQueue,
	}
}

func (s *AwsScheduler) Schedule(schedule models.Schedule, config ScheduleConfig) error {
	ctx := context.Background()
	expression := s.fromScheduleToCronExpression(schedule)

	_, err := s.client.PutRule(ctx, &eventbridge.PutRuleInput{
		Name:               &config.Id,
		ScheduleExpression: &expression,
	})

	if err != nil {
		return err
	}

	_, err = s.client.PutTargets(ctx, &eventbridge.PutTargetsInput{
		Rule: &config.Id,
		Targets: []types.Target{
			{
				Arn:   &s.targetQueue,
				Id:    &config.Id,
				Input: &config.Payload,
			},
		},
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

func (s *AwsScheduler) DeleteSchedule(id string) error {
	ctx := context.Background()
	_, err := s.client.DeleteRule(ctx, &eventbridge.DeleteRuleInput{
		Name: &id,
	})

	return err
}

func (s *AwsScheduler) fromScheduleToCronExpression(schedule models.Schedule) string {
	expression := schedule.ToCronExpression()

	return "cron(" + expression + ")"
}
