package activities

import (
	"context"
	"fmt"
	"math/rand"
)

func ValidateOrder(ctx context.Context, orderID string) error {
	fmt.Println("Validating order:", orderID)

	// simulate failure (like old code)
	if rand.Intn(3) == 0 {
		return fmt.Errorf("validation failed randomly")
	}

	return nil
}
