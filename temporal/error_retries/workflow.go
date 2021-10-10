package error_retries

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

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

func MyActivity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "!", nil
}
