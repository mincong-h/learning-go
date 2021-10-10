package error_retries

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

func (ts *WorkflowTestSuite) TestWorkflow_CompletedSuccessfully() {
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

func (ts *WorkflowTestSuite) TestWorkflow_ExplicitRetryableError() {
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
	ts.Equal(executionCount, 2, "1st execution failed and 2nd execution succeed")
}

func (ts *WorkflowTestSuite) TestWorkflow_ImplicitRetryableError() {
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
	ts.Equal(executionCount, 2, "1st execution failed and 2nd execution succeed")
}

func (ts *WorkflowTestSuite) TestWorkflow_NonRetryableError() {
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
	ts.Equal(executionCount, 1, "1st execution failed but not retried")
}
