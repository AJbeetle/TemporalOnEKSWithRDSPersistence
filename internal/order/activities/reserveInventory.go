package activities

import (
	"context"
	"fmt"
	"time"

	"learningTemp/internal/order/models"

	"go.temporal.io/sdk/activity"
)

var reservedInventory = make(map[string]bool)

func ReserveInventory(ctx context.Context, validateResult *models.ValidateOrderResult) (*models.ReserveInventoryResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ReserveInventory started",
		"orderID", validateResult.OrderID,
		"itemID", validateResult.ItemID,
	)

	// Idempotency check - if already reserved, don't deduct stock again
	// In production: SELECT id FROM reservations WHERE order_id = ?
	if reservedInventory[validateResult.OrderID] {
		logger.Info("already reserved, skipping", "orderID", validateResult.OrderID)
		return &models.ReserveInventoryResult{
			OrderID:       validateResult.OrderID,
			ReservationID: "existing-" + validateResult.OrderID,
			ItemID:        validateResult.ItemID,
			StockDeducted: true,
		}, nil
	}

	// Simulate long running DB write with heartbeat every 2s
	// In production: UPDATE inventory SET stock = stock - 1 WHERE item_id = ?
	totalSteps := 3
	for i := 1; i <= totalSteps; i++ {

		// Checking if workflow cancelled/timed out before doing more work
		if ctx.Err() != nil {
			return nil, fmt.Errorf("ReserveInventory cancelled at step %d: %w", i, ctx.Err())
		}

		// Mock DB write: deducting stock in steps
		// In production: UPDATE inventory SET stock = stock - 1 WHERE item_id = ?
		logger.Info("Deducting stock",
			"step", i,
			"totalSteps", totalSteps,
			"itemID", validateResult.ItemID,
		)

		// Heartbeat every 2s as per spec
		activity.RecordHeartbeat(ctx, fmt.Sprintf("step %d of %d", i, totalSteps))
		time.Sleep(2 * time.Second)
	}

	// Mock: mark as reserved in our in-memory store
	// In production: INSERT INTO reservations (order_id, reserved_at) VALUES (?, ?)
	reservedInventory[validateResult.OrderID] = true
	reservationID := fmt.Sprintf("RES-%s-%d", validateResult.OrderID, time.Now().Unix())

	logger.Info("ReserveInventory completed",
		"orderID", validateResult.OrderID,
		"reservationID", reservationID,
	)

	return &models.ReserveInventoryResult{
		OrderID:       validateResult.OrderID,
		ReservationID: reservationID,
		ItemID:        validateResult.ItemID,
		StockDeducted: true,
	}, nil
}
