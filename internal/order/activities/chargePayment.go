package activities

import (
	"context"
	"fmt"
	"math/rand"

	"learningTemp/internal/order/models"

	"go.temporal.io/sdk/activity"
)

// Mock charge store — in production this is your DB
// SELECT charge_id FROM charges WHERE order_id = ?
var processedCharges = make(map[string]*models.ChargePaymentResult)

func ChargePayment(ctx context.Context, validateResult *models.ValidateOrderResult, inventoryResult *models.ReserveInventoryResult) (*models.ChargePaymentResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ChargePayment started",
		"orderID", validateResult.OrderID,
		"reservationID", inventoryResult.ReservationID,
		"amount", validateResult.Price,
	)

	// ── Idempotency check ────────────────────────────────────────
	// If this activity already ran (e.g. worker crashed after charging
	// but before returning result), don't charge again
	// In production: SELECT * FROM charges WHERE order_id = ?
	if existing, ok := processedCharges[validateResult.OrderID]; ok {
		logger.Info("payment already processed, returning existing charge",
			"orderID", validateResult.OrderID,
			"chargeID", existing.ChargeID,
		)
		return existing, nil
	}

	// INTENTIONALLY FAIL 50% — demonstrates automatic retry
	if rand.Intn(2) == 0 {
		return nil, fmt.Errorf("payment gateway error: transaction declined for order %s", validateResult.OrderID)
	}

	// Mock: charge the payment
	// In production: call Stripe with IdempotencyKey = orderID
	// stripe.Charge(ctx, &stripe.ChargeParams{ IdempotencyKey: &orderID, ... })
	chargeID := fmt.Sprintf("CHG-%s-%d", validateResult.OrderID, rand.Intn(999999))

	result := &models.ChargePaymentResult{
		OrderID:  validateResult.OrderID,
		ChargeID: chargeID,
		Amount:   validateResult.Price,
		Status:   "success",
	}

	// Persist before returning — in production: INSERT INTO charges (...)
	processedCharges[validateResult.OrderID] = result

	logger.Info("ChargePayment completed",
		"orderID", validateResult.OrderID,
		"chargeID", chargeID,
		"amount", validateResult.Price,
	)

	return result, nil
}
