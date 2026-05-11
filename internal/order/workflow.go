package order

import (
	"time"

	"fmt"
	"learningTemp/internal/order/activities"
	"learningTemp/internal/order/models"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func OrderWorkflow(ctx workflow.Context, orderID string) (string, error) {

	// Validate
	validateCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    3,
			InitialInterval:    2 * time.Second,
			BackoffCoefficient: 1.0,
		},
	})

	var validateResult *models.ValidateOrderResult
	if err := workflow.ExecuteActivity(validateCtx, activities.ValidateOrder, orderID).
		Get(validateCtx, &validateResult); err != nil {
		return "validation failed", err
	}

	// Inventory
	inventoryCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 60, // long running — 3 steps × 2s + buffer
		HeartbeatTimeout:    time.Second * 5,  // if no heartbeat for 5s, Temporal marks it failed
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    5,
			InitialInterval:    5 * time.Second,
			BackoffCoefficient: 1.0,
		},
	})

	var inventoryResult *models.ReserveInventoryResult
	if err := workflow.ExecuteActivity(inventoryCtx, activities.ReserveInventory, validateResult).
		Get(inventoryCtx, &inventoryResult); err != nil {
		return "inventory failed", err
	}

	// Payment
	paymentCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts:    3,
			InitialInterval:    10 * time.Second,
			BackoffCoefficient: 1.0,
		},
	})

	var chargeResult *models.ChargePaymentResult
	if err := workflow.ExecuteActivity(paymentCtx, activities.ChargePayment, validateResult, inventoryResult).
		Get(paymentCtx, &chargeResult); err != nil {
		return "payment failed", err
	}

	// Email
	emailCtx := workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	})

	var emailResult *models.SendEmailResult
	if err := workflow.ExecuteActivity(emailCtx, activities.SendConfirmationEmail, chargeResult).
		Get(emailCtx, &emailResult); err != nil {
		return "email failed", err
	}

	return fmt.Sprintf("order %s completed — charge %s", orderID, chargeResult.ChargeID), nil
}
