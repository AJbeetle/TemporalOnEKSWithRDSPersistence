package activities

import (
	"context"
	"fmt"
)

var sentEmails = make(map[string]bool)

func SendConfirmationEmail(ctx context.Context, orderID string) error {

	if sentEmails[orderID] {
		fmt.Println("Email already sent:", orderID)
		return nil
	}

	fmt.Println("Email sent for:", orderID)
	sentEmails[orderID] = true

	return nil
}
