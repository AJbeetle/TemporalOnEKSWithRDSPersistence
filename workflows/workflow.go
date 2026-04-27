package workflows

import (
	"time"

	"learningTemp/activities"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func OrderWorkflow(ctx workflow.Context, orderID string) (string, error) {

	// Activity 1: ValidateOrder
	validateOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
			InitialInterval: time.Second * 2,
		},
	}
	ctx1 := workflow.WithActivityOptions(ctx, validateOptions)

	err := workflow.ExecuteActivity(ctx1, activities.ValidateOrder, orderID).Get(ctx, nil)
	if err != nil {
		return "Validation Failed", err
	}

	// Activity 2: ReserveInventory
	reserveOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 20,
		HeartbeatTimeout:    time.Second * 5,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 5,
			InitialInterval: time.Second * 5,
		},
	}
	ctx2 := workflow.WithActivityOptions(ctx, reserveOptions)

	err = workflow.ExecuteActivity(ctx2, activities.ReserveInventory, orderID).Get(ctx, nil)
	if err != nil {
		return "Inventory Failed", err
	}

	// Activity 3: ChargePayment
	paymentOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 3,
			InitialInterval: time.Second * 10,
		},
	}
	ctx3 := workflow.WithActivityOptions(ctx, paymentOptions)

	err = workflow.ExecuteActivity(ctx3, activities.ChargePayment, orderID).Get(ctx, nil)
	if err != nil {
		return "Payment Failed", err
	}

	emailOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}

	ctx4 := workflow.WithActivityOptions(ctx, emailOptions)

	// Activity 4: SendConfirmationEmail
	err = workflow.ExecuteActivity(ctx4, activities.SendConfirmationEmail, orderID).Get(ctx, nil)

	if err != nil {
		return "Email Failed", err
	}

	return "Order Completed: " + orderID, nil
}
