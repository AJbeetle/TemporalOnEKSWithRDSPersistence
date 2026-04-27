package main

import (
	"log"

	"learningTemp/activities"
	"learningTemp/workflows"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	w := worker.New(c, "order-task-queue", worker.Options{})

	w.RegisterWorkflow(workflows.OrderWorkflow)
	w.RegisterActivity(activities.ValidateOrder)
	w.RegisterActivity(activities.ReserveInventory)
	w.RegisterActivity(activities.ChargePayment)
	w.RegisterActivity(activities.SendConfirmationEmail)

	log.Println("Worker started...")
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatal(err)
	}
}
