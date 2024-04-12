package tidbClient

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/workflow"
	"time"
)

func TiDBWorkflowDefinition(ctx workflow.Context, start string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("TiDB workflow started", "start_time", start)

	var result string
	err := workflow.ExecuteActivity(ctx, TiDBActivityDefinition, start).Get(ctx, &result)
	if err != nil {
		logger.Error("TiDB Activity failed.", "Error", err)
		return "", err
	}

	logger.Info("TiDB workflow completed.", "result", result)

	return result, nil
}

func TiDBActivityDefinition(ctx context.Context, time string) (string, error) {
	version := GetTiDBVersion()
	return fmt.Sprintf(
		"Requested at %s, retrieved version from TiDB as %s",
		time, version), nil
}
