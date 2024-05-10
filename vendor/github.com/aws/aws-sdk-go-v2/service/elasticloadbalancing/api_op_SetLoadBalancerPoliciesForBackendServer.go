// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Replaces the set of policies associated with the specified port on which the
// EC2 instance is listening with a new set of policies. At this time, only the
// back-end server authentication policy type can be applied to the instance ports;
// this policy type is composed of multiple public key policies.
//
// Each time you use SetLoadBalancerPoliciesForBackendServer to enable the
// policies, use the PolicyNames parameter to list the policies that you want to
// enable.
//
// You can use DescribeLoadBalancers or DescribeLoadBalancerPolicies to verify that the policy is associated with the EC2 instance.
//
// For more information about enabling back-end instance authentication, see [Configure Back-end Instance Authentication] in
// the Classic Load Balancers Guide. For more information about Proxy Protocol, see
// [Configure Proxy Protocol Support]in the Classic Load Balancers Guide.
//
// [Configure Back-end Instance Authentication]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-create-https-ssl-load-balancer.html#configure_backendauth_clt
// [Configure Proxy Protocol Support]: https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/enable-proxy-protocol.html
func (c *Client) SetLoadBalancerPoliciesForBackendServer(ctx context.Context, params *SetLoadBalancerPoliciesForBackendServerInput, optFns ...func(*Options)) (*SetLoadBalancerPoliciesForBackendServerOutput, error) {
	if params == nil {
		params = &SetLoadBalancerPoliciesForBackendServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetLoadBalancerPoliciesForBackendServer", params, optFns, c.addOperationSetLoadBalancerPoliciesForBackendServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetLoadBalancerPoliciesForBackendServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetLoadBalancerPoliciesForBackendServer.
type SetLoadBalancerPoliciesForBackendServerInput struct {

	// The port number associated with the EC2 instance.
	//
	// This member is required.
	InstancePort *int32

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The names of the policies. If the list is empty, then all current polices are
	// removed from the EC2 instance.
	//
	// This member is required.
	PolicyNames []string

	noSmithyDocumentSerde
}

// Contains the output of SetLoadBalancerPoliciesForBackendServer.
type SetLoadBalancerPoliciesForBackendServerOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetLoadBalancerPoliciesForBackendServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetLoadBalancerPoliciesForBackendServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetLoadBalancerPoliciesForBackendServer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetLoadBalancerPoliciesForBackendServer"); err != nil {
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
	if err = addOpSetLoadBalancerPoliciesForBackendServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetLoadBalancerPoliciesForBackendServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetLoadBalancerPoliciesForBackendServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetLoadBalancerPoliciesForBackendServer",
	}
}
