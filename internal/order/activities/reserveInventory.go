package activities

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
)

var reservedInventory = make(map[string]bool)

func ReserveInventory(ctx context.Context, orderID string) error {

	if reservedInventory[orderID] {
		fmt.Println("Inventory already reserved:", orderID)
		return nil
	}

	fmt.Println("Reserving inventory:", orderID)

	// simulate long running work + heartbeat
	for i := 0; i < 5; i++ {
		activity.RecordHeartbeat(ctx, i)
		fmt.Println("Heartbeat sent:", i)
		time.Sleep(2 * time.Second)
	}

	reservedInventory[orderID] = true
	return nil
}
