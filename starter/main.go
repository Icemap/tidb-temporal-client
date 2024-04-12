package main

import (
	"context"
	"log"
	"time"

	"go.temporal.io/sdk/client"

	tidbClient "github.com/Icemap/tidb-temporal-client"
)

func main() {
	// Create a Temporal Client
	// A Temporal Client is a heavyweight object that should be created just once per process.
	clientOptions := client.Options{
		HostPort:  "192.168.0.3:7233",
		Namespace: "default",
	}
	c, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "TiDBVersionFetcher",
		TaskQueue: "tidb-task-queue",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions,
		tidbClient.TiDBWorkflowDefinition, time.Now().String())
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// Synchronously wait for the workflow completion.
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable get workflow result", err)
	}
	log.Println("Workflow result:", result)
}
