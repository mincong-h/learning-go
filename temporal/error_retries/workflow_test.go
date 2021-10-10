package error_retries

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
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

func (ts *WorkflowTestSuite) TestWorkflow() {
	// Given
	ts.env.OnActivity(MyActivity, mock.Anything, mock.Anything).Return("Hello, UnitTest!", nil)

	// When
	ts.env.ExecuteWorkflow(MyWorkflow, "UnitTest")

	// When
	ts.True(ts.env.IsWorkflowCompleted())
	ts.NoError(ts.env.GetWorkflowError())
}
