package error_retries

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/testsuite"
)

type WorkflowTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite

	env *testsuite.TestWorkflowEnvironment
}

func TestWorkflowTestSuite(t *testing.T) {
	suite.Run(t, new(WorkflowTestSuite))
}

func (ts *WorkflowTestSuite) SetupTest() {
	ts.env = ts.NewTestWorkflowEnvironment()
	ts.env.RegisterWorkflow(MyWorkflow)
}

func (ts *WorkflowTestSuite) AfterTest() {
	ts.env.AssertExpectations(ts.T())
}

func (ts *WorkflowTestSuite) TestActivityError_CompletedSuccessfullyWithoutError() {
	// Given
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return("Hello, UnitTest!", nil)

	// When
	ts.env.ExecuteWorkflow(MyWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())
	ts.NoError(ts.env.GetWorkflowError())

	var result string
	ts.NoError(ts.env.GetWorkflowResult(&result))
	ts.Equal("Hello, UnitTest!", result)
}

func (ts *WorkflowTestSuite) TestActivityError_ExplicitRetryableError() {
	// Given
	executionCount := 0
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return(func(ctx context.Context, msg string) (string, error) {
		executionCount++
		if executionCount == 1 {
			return "", temporal.NewApplicationError("oops", "test")
		} else {
			return "Hello, UnitTest!", nil
		}
	})

	// When
	ts.env.ExecuteWorkflow(MyWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())
	ts.NoError(ts.env.GetWorkflowError())

	var result string
	ts.NoError(ts.env.GetWorkflowResult(&result))
	ts.Equal("Hello, UnitTest!", result)
	ts.Equal(2, executionCount, "1st execution failed and 2nd execution succeed")
}

func (ts *WorkflowTestSuite) TestActivityError_ImplicitRetryableError() {
	// Given
	executionCount := 0
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return(func(ctx context.Context, msg string) (string, error) {
		executionCount++
		if executionCount == 1 {
			// Temporal transforms this error into temporal.NewApplicationError
			// which is retryable
			return "", fmt.Errorf("oops")
		} else {
			return "Hello, UnitTest!", nil
		}
	})

	// When
	ts.env.ExecuteWorkflow(MyWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())
	ts.NoError(ts.env.GetWorkflowError())

	var result string
	ts.NoError(ts.env.GetWorkflowResult(&result))
	ts.Equal("Hello, UnitTest!", result)
	ts.Equal(2, executionCount, "1st execution failed and 2nd execution succeed")
}

func (ts *WorkflowTestSuite) TestActivityError_NonRetryableError() {
	// Given
	executionCount := 0
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return(func(ctx context.Context, msg string) (string, error) {
		executionCount++
		if executionCount == 1 {
			return "", temporal.NewNonRetryableApplicationError("oops", "test", nil)
		} else {
			return "Hello, UnitTest!", nil
		}
	})

	// When
	ts.env.ExecuteWorkflow(MyWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())

	var err *temporal.ApplicationError
	ts.True(errors.As(ts.env.GetWorkflowError(), &err))
	ts.True(err.NonRetryable())
	ts.True(strings.Contains(err.Error(), "oops"))
	ts.Equal(1, executionCount, "1st execution failed but not retried")
}

func (ts *WorkflowTestSuite) TestActivityError_NonRetryForErrorsInRetryPolicy() {
	// Given
	executionCount := 0
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return(func(ctx context.Context, msg string) (string, error) {
		executionCount++
		if executionCount == 1 {
			return "", &MyError{
				Message: "oops",
			}
		} else {
			return "Hello, UnitTest!", nil
		}
	})

	// When
	ts.env.ExecuteWorkflow(MyWorkflowWithRetryPolicy, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())

	var err *temporal.ApplicationError
	ts.True(errors.As(ts.env.GetWorkflowError(), &err))
	// NOTE: this error is marked as retryable but it was not retried
	ts.False(err.NonRetryable())

	ts.True(strings.Contains(err.Error(), "oops"))
	ts.Equal(1, executionCount, "1st execution failed but not retried")
}

func (ts *WorkflowTestSuite) TestActivityError_KeepFailingUntilMaximumAttempts() {
	// Given
	executionCount := 0
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return(func(ctx context.Context, msg string) (string, error) {
		executionCount++
		return "", fmt.Errorf("oops (%d)", executionCount)
	})

	// When
	ts.env.ExecuteWorkflow(MyWorkflowWithRetryPolicy, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())

	var err *temporal.ApplicationError
	ts.True(errors.As(ts.env.GetWorkflowError(), &err))
	ts.False(err.NonRetryable())

	ts.True(strings.Contains(err.Error(), "oops"))
	ts.Equal(5, executionCount, "maximum attempts reached")
}

func (ts *WorkflowTestSuite) TestWorkflowError_NonRetry() {
	// Given

	// When
	ts.env.ExecuteWorkflow(MyFailingWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())

	var err *temporal.WorkflowExecutionError
	ts.True(errors.As(ts.env.GetWorkflowError(), &err))
}

func (ts *WorkflowTestSuite) TestWorkflowError_RetryWorkflow() {
	// Given
	ts.env.SetStartWorkflowOptions(client.StartWorkflowOptions{
		// NOTE: compared to the default test environment, here we introduce a retry policy so that
		// Temporal server will retry all the time when error occurred.
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    1 * time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    1 * time.Minute,
			MaximumAttempts:    5,
		},
	})

	// When
	ts.env.ExecuteWorkflow(MyFailingWorkflow, "UnitTest")

	// Then
	ts.True(ts.env.IsWorkflowCompleted())

	var err *temporal.WorkflowExecutionError
	ts.True(errors.As(ts.env.GetWorkflowError(), &err))
	// TODO: how to ensure that the workflow is retried??
}
