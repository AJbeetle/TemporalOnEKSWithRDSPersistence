// internal/order/activities/validate.go
package activities

import (
	"context"
	"fmt"
	"math/rand"

	"learningTemp/internal/order/models"

	"go.temporal.io/sdk/activity"
)

func ValidateOrder(ctx context.Context, orderID string) (*models.ValidateOrderResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder started", "orderID", orderID)

	// Mock: check item availability
	// In production: SELECT item_id, price, stock FROM orders WHERE id = ?
	available := rand.Intn(10) >= 2 // 80% available
	if !available {
		return nil, fmt.Errorf("item not available for order: %s", orderID)
	}

	// Mock: return order details
	// In production: these come from your DB
	result := &models.ValidateOrderResult{
		OrderID: orderID,
		ItemID:  fmt.Sprintf("ITEM-%s", orderID), // mock item ID
		Price:   int64(rand.Intn(90000) + 10000), // mock price: 100.00 to 999.99 in paise
	}

	logger.Info("ValidateOrder completed",
		"orderID", orderID,
		"itemID", result.ItemID,
		"price", result.Price,
	)

	return result, nil
}
