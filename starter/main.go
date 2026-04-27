package main

import (
	"context"
	"log"

	"learningTemp/workflows"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "order-workflow-trying-success",
		TaskQueue: "order-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, workflows.OrderWorkflow, "ORDER-123456")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Started workflow:", we.GetID())
}
