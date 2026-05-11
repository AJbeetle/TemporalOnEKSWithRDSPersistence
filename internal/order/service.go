package order

import (
	"context"

	"go.temporal.io/sdk/client"
)

type Service struct {
	client client.Client
}

func NewService(c client.Client) *Service {
	return &Service{client: c}
}

// service.go
func (s *Service) CreateOrder(ctx context.Context, orderID string) (client.WorkflowRun, error) {
	run, err := s.client.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "order-" + orderID,
		TaskQueue: TaskQueue,
	}, OrderWorkflow, orderID)

	if err != nil {
		return nil, err
	}

	return run, nil
}
