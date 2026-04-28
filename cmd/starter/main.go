package main

// import (
// 	"context"
// 	"log"

// 	"learningTemp/internal/config"
// 	"learningTemp/internal/order"

// 	"go.temporal.io/sdk/client"
// )

// func main() {
// 	cfg := config.Load()
// 	c, err := client.Dial(client.Options{
// 		HostPort:  cfg.TemporalHost,
// 		Namespace: "default",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer c.Close()

// 	options := client.StartWorkflowOptions{
// 		ID:        "order-workflow-trying-success",
// 		TaskQueue: order.TaskQueue,
// 	}

// 	we, err := c.ExecuteWorkflow(context.Background(), options, order.OrderWorkflow, "ORDER-123456")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Started workflow:", we.GetID())
// }
