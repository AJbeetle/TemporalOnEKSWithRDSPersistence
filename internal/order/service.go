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

func (s *Service) CreateOrder(ctx context.Context, orderID string) error {
	_, err := s.client.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		ID:        "order-" + orderID, // ✅ idempotency at workflow level
		TaskQueue: TaskQueue,
	}, OrderWorkflow, orderID)

	return err
}
