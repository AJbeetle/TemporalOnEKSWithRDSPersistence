package order

import (
	"time"

	"learningTemp/internal/order/activities"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func OrderWorkflow(ctx workflow.Context, orderID string) (string, error) {

	// Validate
	ctx1 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	})

	if err := workflow.ExecuteActivity(ctx1, activities.ValidateOrder, orderID).Get(ctx, nil); err != nil {
		return "validation failed", err
	}

	// Inventory
	ctx2 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 20,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
		},
	})

	if err := workflow.ExecuteActivity(ctx2, activities.ReserveInventory, orderID).Get(ctx, nil); err != nil {
		return "inventory failed", err
	}

	// Payment
	ctx3 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
		},
	})

	if err := workflow.ExecuteActivity(ctx3, activities.ChargePayment, orderID).Get(ctx, nil); err != nil {
		return "payment failed", err
	}

	// Email
	ctx4 := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	})

	if err := workflow.ExecuteActivity(ctx4, activities.SendConfirmationEmail, orderID).Get(ctx, nil); err != nil {
		return "email failed", err
	}

	return "order completed", nil
}
