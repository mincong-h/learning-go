package error_retries

import (
	"context"
	"fmt"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type MyError struct {
	Message string
}

func (err *MyError) Error() string {
	return err.Message
}

func MyWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)
	var result string
	err := workflow.ExecuteActivity(ctx, MyActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Workflow completed.", "result", result)
	return result, nil
}

func MyWorkflow2(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)
	err := fmt.Errorf("oops")
	logger.Error(fmt.Sprintf("Workflow failed: %v", err))
	return "", err
}

func MyWorkflow3(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:        1 * time.Second,
			BackoffCoefficient:     2,
			MaximumInterval:        1 * time.Minute,
			MaximumAttempts:        5,
			NonRetryableErrorTypes: []string{"MyError"},
		},
	})
	logger := workflow.GetLogger(ctx)
	var result string
	err := workflow.ExecuteActivity(ctx, MyActivity, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Workflow completed.", "result", result)
	return result, nil
}

func MyActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}
