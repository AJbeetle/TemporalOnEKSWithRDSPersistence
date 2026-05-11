// internal/order/activities/send_email.go
package activities

import (
	"context"
	"fmt"

	"learningTemp/internal/order/models"

	"go.temporal.io/sdk/activity"
)

// Mock email store — in production: SELECT id FROM sent_emails WHERE order_id = ?
var sentEmails = make(map[string]*models.SendEmailResult)

func SendConfirmationEmail(ctx context.Context, chargeResult *models.ChargePaymentResult) (*models.SendEmailResult, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SendConfirmationEmail started",
		"orderID", chargeResult.OrderID,
		"chargeID", chargeResult.ChargeID,
	)

	// ── Idempotency check ────────────────────────────────────────
	// Email sends must never be duplicated — customer gets spammed
	// In production: SELECT id FROM sent_emails WHERE order_id = ?
	if existing, ok := sentEmails[chargeResult.OrderID]; ok {
		logger.Info("email already sent, skipping",
			"orderID", chargeResult.OrderID,
			"email", existing.Email,
		)
		return existing, nil
	}

	// Mock: send email
	// In production: SES/SendGrid call here
	email := fmt.Sprintf("customer-%s@example.com", chargeResult.OrderID)
	fmt.Printf("📧 email sent to %s — Order %s confirmed, Charge %s, Amount %d Rs\n",
		email,
		chargeResult.OrderID,
		chargeResult.ChargeID,
		chargeResult.Amount,
	)

	result := &models.SendEmailResult{
		OrderID:  chargeResult.OrderID,
		ChargeID: chargeResult.ChargeID,
		Email:    email,
		Sent:     true,
	}

	// Persist before returning — in production: INSERT INTO sent_emails (...)
	sentEmails[chargeResult.OrderID] = result

	logger.Info("SendConfirmationEmail completed",
		"orderID", chargeResult.OrderID,
		"email", email,
	)

	return result, nil
}
