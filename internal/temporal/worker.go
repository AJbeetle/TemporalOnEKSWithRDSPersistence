package temporal

import (
	"log"

	"learningTemp/internal/order"
	"learningTemp/internal/order/activities"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func StartWorker(c client.Client) {

	w := worker.New(c, order.TaskQueue, worker.Options{})

	w.RegisterWorkflow(order.OrderWorkflow)

	w.RegisterActivity(activities.ValidateOrder)
	w.RegisterActivity(activities.ReserveInventory)
	w.RegisterActivity(activities.ChargePayment)
	w.RegisterActivity(activities.SendConfirmationEmail)

	log.Println("Worker started...")

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatal(err)
	}
}
