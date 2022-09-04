// Worker is like a process that will host the "Workflow" and "Activity" code.
// Worker is equivalent to a server in a microservice architecture.
// So if our business application starts getting a lot of load and a lot more users.
// We have to scale our workers in addition to the temporal cluster.

// We have to run the "worker" like a go proces so we need to give worker a pakcage main.
// A worker's point and purpose is that it will host our workflow code.
// so first thing we do is instantiate a temporal client.
// Just like we would have a sql/redis/postgres client.
package main

import (
	"log"

	"github.com/arjunmalhotra1/hellotemporal/helloworkflow"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {

	c, err := client.NewLazyClient(client.Options{})
	if err != nil {
		log.Fatal("Unable to make client", err)
	}

	// Since this is not a connection pooling it is a resource, so
	// we want to make sure that it is closed after it's done being used.
	defer c.Close()

	// Now we create our worker.
	// We pass the worker the client & pass it a task queue.
	// task queue is a routing mechanism or a load balancing mechanism in temporal.
	w := worker.New(c, "hello-world", worker.Options{})

	// Next what we have to do is to register our workflow and activity implementations with the worker.
	// We register this workflow. We do this so that the worker that we created can process it.
	w.RegisterWorkflow(helloworkflow.Workflow)

	w.RegisterActivity(helloworkflow.Activity)

	// Run the worker.
	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start workflow.", err)
	}
}
