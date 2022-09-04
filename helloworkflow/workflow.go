package helloworkflow

import (
	"context"
	"time"

	"go.temporal.io/sdk/workflow"
)

// Workflows & Activites are special purpose functions.
// Workflows require one parameter always that is "workflow.Context".
// Workflows also have a required return "error"
func Workflow(ctx workflow.Context, name string) (string, error) {
	// This provides the meta data or configuration about how our activity will be executed
	// in the context of this workflow.
	ao := workflow.ActivityOptions{
		// How long the workflows have to wait to pick up a task.
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
	}
	// Now the context has the Activity Options. So if we use the context in any calls,
	// throughout the workflows from now on, it will actually use that context with those options.
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)

	var result string
	err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
	// The activity could fail to even be sent to the service itself. So we would want to know
	// to be able to resend it.
	// Temporal provides built-in retries and timeouts for activities.
	// So even if the activity were to fail in the default case would be to retry it
	// automatically. So that we don't have to deal with it.
	if err != nil {
		logger.Error("Activity Failed.", "Error", err)
	}
	return result, nil
}

// Note: Activity requires the Golang built in context and not the temporal SDK workflow context.
// Activity are also required to return an error.
func Activity(ctx context.Context, name string) (string, error) {
	return "Hello " + name, nil
}
