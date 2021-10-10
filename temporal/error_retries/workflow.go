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

func MyWorkflowWithRetryPolicy(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    1 * time.Minute,
			MaximumAttempts:    5,
			// NOTE: compared to MyWorkflow, here we introduce a retry policy which includes
			// non-retryable error types. Temporal server will stop retry if error type matches
			// this list.
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

func MyWorkflowWithChildWorkflowRetryPolicy(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		WorkflowRunTimeout: 10 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    1 * time.Minute,
			MaximumAttempts:    5,
			// NOTE: compared to MyWorkflow, here we introduce a retry policy which includes
			// non-retryable error types. Temporal server will stop retry if error type matches
			// this list.
			NonRetryableErrorTypes: []string{"MyError"},
		},
	})
	logger := workflow.GetLogger(ctx)
	var result string
	err := workflow.ExecuteChildWorkflow(ctx, MyFailingWorkflow, name).Get(ctx, &result)
	if err != nil {
		logger.Error("Child workflow failed.", "Error", err)
		return "", err
	}
	logger.Info("Workflow completed.", "result", result)
	return result, nil
}

func MyFailingWorkflow(ctx workflow.Context, name string) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	})
	logger := workflow.GetLogger(ctx)
	err := fmt.Errorf("oops")
	logger.Error(fmt.Sprintf("Workflow failed: %v", err))
	return "", err
}

func MyActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}
