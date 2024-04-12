package tidbClient

import (
	"fmt"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
	"testing"
	"time"
)

func Test_Workflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	timeNow := time.Now().String()
	mockVersion := "8.0.11-TiDB-v8.1.0"
	mockReturn := fmt.Sprintf(
		"Requested at %s, retrieved version from TiDB as %s",
		timeNow, mockVersion)
	env.OnActivity(TiDBActivityDefinition, mock.Anything, timeNow).Return(mockReturn, nil)

	env.ExecuteWorkflow(TiDBWorkflowDefinition, timeNow)

	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
	var result string
	require.NoError(t, env.GetWorkflowResult(&result))

	fmt.Println(result)
}

func Test_Activity(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestActivityEnvironment()
	env.RegisterActivity(TiDBActivityDefinition)

	timeNow := time.Now().String()
	val, err := env.ExecuteActivity(TiDBActivityDefinition, timeNow)
	require.NoError(t, err)

	var res string
	require.NoError(t, val.Get(&res))
	fmt.Println(res)
}
