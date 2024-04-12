package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"log"

	tidbClient "github.com/Icemap/tidb-temporal-client"
)

func main() {
	// Create a Temporal Client
	// A Temporal Client is a heavyweight object that should be created just once per process.
	clientOptions := client.Options{
		HostPort:  "192.168.0.3:7233",
		Namespace: "default",
	}

	temporalClient, err := client.Dial(clientOptions)
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer temporalClient.Close()

	tidbWorker := worker.New(temporalClient, "tidb-task-queue", worker.Options{})
	registerWFOptions := workflow.RegisterOptions{
		Name: "TiDBVersionFetcher",
	}
	tidbWorker.RegisterWorkflowWithOptions(tidbClient.TiDBWorkflowDefinition, registerWFOptions)
	tidbWorker.RegisterActivity(tidbClient.TiDBActivityDefinition)

	// Run the Worker
	err = tidbWorker.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start Worker", err)
	}
}
