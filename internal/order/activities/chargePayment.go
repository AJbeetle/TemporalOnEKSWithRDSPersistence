package activities

import (
	"context"
	"fmt"
	"math/rand"
)

var processedPayments = make(map[string]bool)

func ChargePayment(ctx context.Context, orderID string) error {

	if processedPayments[orderID] {
		fmt.Println("Payment already processed:", orderID)
		return nil
	}

	fmt.Println("Charging payment for:", orderID)

	// simulate failure (important for retry demo)
	if rand.Intn(2) == 0 {
		return fmt.Errorf("payment failed")
	}

	fmt.Println("Payment success:", orderID)
	processedPayments[orderID] = true

	return nil
}
