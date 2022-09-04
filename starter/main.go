// Starte is essentially a small program that is used to kickstart/kickoff our workflow.
// Theoretically we we could do this via RCLI or directly with GRPC.
// It's better to have a starter.
package main

import (
	"context"
	"log"

	"github.com/arjunmalhotra1/hellotemporal/helloworkflow"
	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create client.", err)
	}

	defer c.Close()

	// We have to start by defining some workflow options.
	// This just says to the temporal service that what
	// This just says to the Temporal service that which workflow you are starting and
	// which task it would start on.
	workflowOptions := client.StartWorkflowOptions{
		ID:        "hello-workflow-video",
		TaskQueue: "hello-world", // We have already defined the task queue in worker/main.go
	}

	// Now we need to take the workflow execution & start it.
	// "Temporal" is the argument that we will pass in to the workflow. It maps to the
	// "name" parameter which is a string parameter in the workflow. Hence we are putting that in the end.
	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworkflow.Workflow, "Temporal-first-video")
	// This is not an error from the workflow level code failing.
	// This is an error if the temporal servicet is unable to even register/run the
	// workflow in the first place.
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	// On the above line we have executed the workflow but we don't know what the result is.
	var result string
	err = we.Get(context.Background(), &result) // The result of execution of the workflow will be
	// stored in the "result" variable.

	// This err means that the workflow has already been successfully ran by the
	// temporal service. It's running but maybe the temporal service was down when we tried
	// to get the result or something.
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}
	log.Println("Workflow result: ", result)

}

// We have a workflow.
// We have an activity that it calls & returns something back to the user.
// We have a worker which can actually host that workflow code and can make it
// available when we send the workflow start request.
// We also have workflow starter which is able to actually send a signal to the temporal service
// that the workflow execution needs to start & then get the response back to the user with the
// workflow get call.
