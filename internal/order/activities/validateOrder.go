package activities

import (
	"context"
	"fmt"
	"math/rand"

	"go.temporal.io/sdk/activity"
)

type ValidateOrderResult struct {
	OrderID     string
	IsAvailable bool
	Message     string
}

func ValidateOrder(ctx context.Context, orderID string) (*ValidateOrderResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ValidateOrder started", "orderID", orderID)

	// Mock: checking item availability
	available := checkItemAvailability(orderID)

	if !available {
		// Returning an error triggers Temporal's retry policy
		return nil, fmt.Errorf("item not available for order: %s", orderID)
	}

	logger.Info("ValidateOrder completed", "orderID", orderID, "available", true)

	return &ValidateOrderResult{
		OrderID:     orderID,
		IsAvailable: true,
		Message:     "item available",
	}, nil
}

// Returns false ~20% of the time to simulate unavailability
func checkItemAvailability(orderID string) bool {
	// In production: queries to inventory DB or service
	roll := rand.Intn(10)
	if roll < 2 {
		// 20% chance — item not available
		return false
	}
	return true
}
