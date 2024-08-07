package messages

import (
	"go.temporal.io/sdk/workflow"
)

// "getProgress" query handler
func SetQueryHandlerForProgress(ctx workflow.Context) (*int, error) {
	logger := workflow.GetLogger(ctx)

	progress := 0

	err := workflow.SetQueryHandler(ctx, "getProgress", func() (int, error) {
		return progress, nil
	})
	if err != nil {
		logger.Error("SetQueryHandler failed for getProgress: " + err.Error())
		return nil, err
	}

	return &progress, nil
}
