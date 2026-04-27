package activities

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.temporal.io/sdk/activity"
)

func ValidateOrder(ctx context.Context, orderID string) error {
	fmt.Println("Validating order:", orderID)

	// mock failure sometimes
	if rand.Intn(3) == 0 {
		return fmt.Errorf("validation failed randomly")
	}

	return nil
}

func ReserveInventory(ctx context.Context, orderID string) error {
	fmt.Println("Reserving inventory for:", orderID)

	for i := 0; i < 5; i++ {
		activity.RecordHeartbeat(ctx, i)
		fmt.Println("Heartbeat sent:", i)
		time.Sleep(2 * time.Second)
	}

	return nil
}

func ChargePayment(ctx context.Context, orderID string) error {
	fmt.Println("Charging payment for:", orderID)

	// FAIL 50% intentionally
	if rand.Intn(2) == 0 {
		return fmt.Errorf("payment failed")
	}

	fmt.Println("Payment success")
	return nil
}

func SendConfirmationEmail(ctx context.Context, orderID string) error {
	fmt.Println("Email sent for:", orderID)
	return nil
}
