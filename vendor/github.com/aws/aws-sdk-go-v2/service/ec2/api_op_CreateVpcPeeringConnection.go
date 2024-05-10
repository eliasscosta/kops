// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests a VPC peering connection between two VPCs: a requester VPC that you
// own and an accepter VPC with which to create the connection. The accepter VPC
// can belong to another Amazon Web Services account and can be in a different
// Region to the requester VPC. The requester VPC and accepter VPC cannot have
// overlapping CIDR blocks.
//
// Limitations and rules apply to a VPC peering connection. For more information,
// see the [limitations]section in the VPC Peering Guide.
//
// The owner of the accepter VPC must accept the peering request to activate the
// peering connection. The VPC peering connection request expires after 7 days,
// after which it cannot be accepted or rejected.
//
// If you create a VPC peering connection request between VPCs with overlapping
// CIDR blocks, the VPC peering connection has a status of failed .
//
// [limitations]: https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html#vpc-peering-limitations
func (c *Client) CreateVpcPeeringConnection(ctx context.Context, params *CreateVpcPeeringConnectionInput, optFns ...func(*Options)) (*CreateVpcPeeringConnectionOutput, error) {
	if params == nil {
		params = &CreateVpcPeeringConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpcPeeringConnection", params, optFns, c.addOperationCreateVpcPeeringConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpcPeeringConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVpcPeeringConnectionInput struct {

	// The ID of the requester VPC. You must specify this parameter in the request.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The Amazon Web Services account ID of the owner of the accepter VPC.
	//
	// Default: Your Amazon Web Services account ID
	PeerOwnerId *string

	// The Region code for the accepter VPC, if the accepter VPC is located in a
	// Region other than the Region in which you make the request.
	//
	// Default: The Region in which you make the request.
	PeerRegion *string

	// The ID of the VPC with which you are creating the VPC peering connection. You
	// must specify this parameter in the request.
	PeerVpcId *string

	// The tags to assign to the peering connection.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateVpcPeeringConnectionOutput struct {

	// Information about the VPC peering connection.
	VpcPeeringConnection *types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVpcPeeringConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpcPeeringConnection{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVpcPeeringConnection"); err != nil {
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
	if err = addOpCreateVpcPeeringConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpcPeeringConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpcPeeringConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVpcPeeringConnection",
	}
}
