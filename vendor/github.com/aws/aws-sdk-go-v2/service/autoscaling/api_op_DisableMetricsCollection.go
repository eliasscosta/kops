// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disables group metrics collection for the specified Auto Scaling group.
func (c *Client) DisableMetricsCollection(ctx context.Context, params *DisableMetricsCollectionInput, optFns ...func(*Options)) (*DisableMetricsCollectionOutput, error) {
	if params == nil {
		params = &DisableMetricsCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableMetricsCollection", params, optFns, c.addOperationDisableMetricsCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableMetricsCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableMetricsCollectionInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Identifies the metrics to disable.
	//
	// You can specify one or more of the following metrics:
	//
	//   - GroupMinSize
	//
	//   - GroupMaxSize
	//
	//   - GroupDesiredCapacity
	//
	//   - GroupInServiceInstances
	//
	//   - GroupPendingInstances
	//
	//   - GroupStandbyInstances
	//
	//   - GroupTerminatingInstances
	//
	//   - GroupTotalInstances
	//
	//   - GroupInServiceCapacity
	//
	//   - GroupPendingCapacity
	//
	//   - GroupStandbyCapacity
	//
	//   - GroupTerminatingCapacity
	//
	//   - GroupTotalCapacity
	//
	//   - WarmPoolDesiredCapacity
	//
	//   - WarmPoolWarmedCapacity
	//
	//   - WarmPoolPendingCapacity
	//
	//   - WarmPoolTerminatingCapacity
	//
	//   - WarmPoolTotalCapacity
	//
	//   - GroupAndWarmPoolDesiredCapacity
	//
	//   - GroupAndWarmPoolTotalCapacity
	//
	// If you omit this property, all metrics are disabled.
	//
	// For more information, see [Auto Scaling group metrics] in the Amazon EC2 Auto Scaling User Guide.
	//
	// [Auto Scaling group metrics]: https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-cloudwatch-monitoring.html#as-group-metrics
	Metrics []string

	noSmithyDocumentSerde
}

type DisableMetricsCollectionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisableMetricsCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDisableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDisableMetricsCollection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DisableMetricsCollection"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDisableMetricsCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableMetricsCollection(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDisableMetricsCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DisableMetricsCollection",
	}
}
