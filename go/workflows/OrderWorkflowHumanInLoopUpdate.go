package workflows

// import (
// 	"errors"
// 	"time"

// 	"temporal-order-management/activities"

// 	"github.com/google/uuid"

// 	"temporal-order-management/resources"

// 	"go.temporal.io/sdk/temporal"
// 	"go.temporal.io/sdk/workflow"
// )

// func OrderWorkflowHumanInLoopUpdate(ctx workflow.Context, input resources.OrderInput) (*resources.OrderOutput, error) {
// 	logger := workflow.GetLogger(ctx)
// 	logger.Info("Processing order started", "orderId", input.OrderId)

// 	// activity options
// 	activityOptions := workflow.ActivityOptions{
// 		StartToCloseTimeout: 5 * time.Second,
// 		RetryPolicy: &temporal.RetryPolicy{
// 			InitialInterval:    1 * time.Second,
// 			BackoffCoefficient: 2.0,
// 			MaximumInterval:    30 * time.Second,
// 		},
// 	}
// 	ctx = workflow.WithActivityOptions(ctx, activityOptions)

// 	// local activity options
// 	localActivityOptions := workflow.LocalActivityOptions{
// 		StartToCloseTimeout: 5 * time.Second,
// 	}
// 	laCtx := workflow.WithLocalActivityOptions(ctx, localActivityOptions)

// 	// Expose progress as query
// 	progress, err := resources.SetQueryHandlerForProgress(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Get items
// 	items := resources.Items{}
// 	err = workflow.ExecuteLocalActivity(laCtx, activities.GetItems).Get(ctx, &items)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Check Fraud
// 	var result1 string
// 	err = workflow.ExecuteActivity(ctx, activities.CheckFraud, input).Get(ctx, &result1)
// 	if err != nil {
// 		return nil, err
// 	}

// 	*progress = 25
// 	workflow.Sleep(ctx, 3*time.Second)

// 	// Prepare Shipment
// 	var result2 string
// 	err = workflow.ExecuteActivity(ctx, activities.PrepareShipment, input).Get(ctx, &result2)
// 	if err != nil {
// 		return nil, err
// 	}

// 	*progress = 50
// 	workflow.Sleep(ctx, 3*time.Second)

// 	// Charge Customer
// 	var result3 string
// 	err = workflow.ExecuteActivity(ctx, activities.ChargeCustomer, input).Get(ctx, &result3)
// 	if err != nil {
// 		return nil, err
// 	}

// 	*progress = 75

// 	// Start timer and wait for timer to fire or update
// 	address, isCancelled := resources.UpdateApprovalTimer(ctx)
// 	if isCancelled {
// 		// Rollback Charge Customer
// 		var result3 string
// 		_ = workflow.ExecuteActivity(ctx, activities.ChargeCustomerRollback, input).Get(ctx, &result3)

// 		return nil, errors.New("time limit for approval has been exceeded")
// 	}

// 	input.Address = address
// 	workflow.Sleep(ctx, 3*time.Second)

// 	// Ship Orders
// 	var shipItems []workflow.Future
// 	for _, item := range items {
// 		logger.Info("Shipping item " + item.Description)
// 		shipItem := workflow.ExecuteActivity(ctx, activities.ShipOrder, input, item)
// 		shipItems = append(shipItems, shipItem)
// 	}

// 	// Wait for all items to ship
// 	for _, shipItem := range shipItems {
// 		err = shipItem.Get(ctx, nil)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	*progress = 100

// 	// Generate Tracking Id
// 	trackingId := uuid.New().String()

// 	output := &resources.OrderOutput{
// 		TrackingId: trackingId,
// 		Address:    input.Address,
// 	}

// 	return output, nil
// }
