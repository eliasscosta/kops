// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a patch baseline.
func (c *Client) GetPatchBaseline(ctx context.Context, params *GetPatchBaselineInput, optFns ...func(*Options)) (*GetPatchBaselineOutput, error) {
	if params == nil {
		params = &GetPatchBaselineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPatchBaseline", params, optFns, c.addOperationGetPatchBaselineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPatchBaselineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPatchBaselineInput struct {

	// The ID of the patch baseline to retrieve.
	//
	// To retrieve information about an Amazon Web Services managed patch baseline,
	// specify the full Amazon Resource Name (ARN) of the baseline. For example, for
	// the baseline AWS-AmazonLinuxDefaultPatchBaseline , specify
	// arn:aws:ssm:us-east-2:733109147000:patchbaseline/pb-0e392de35e7c563b7 instead of
	// pb-0e392de35e7c563b7 .
	//
	// This member is required.
	BaselineId *string

	noSmithyDocumentSerde
}

type GetPatchBaselineOutput struct {

	// A set of rules used to include patches in the baseline.
	ApprovalRules *types.PatchRuleGroup

	// A list of explicitly approved patches for the baseline.
	ApprovedPatches []string

	// Returns the specified compliance severity level for approved patches in the
	// patch baseline.
	ApprovedPatchesComplianceLevel types.PatchComplianceLevel

	// Indicates whether the list of approved patches includes non-security updates
	// that should be applied to the managed nodes. The default value is false .
	// Applies to Linux managed nodes only.
	ApprovedPatchesEnableNonSecurity *bool

	// The ID of the retrieved patch baseline.
	BaselineId *string

	// The date the patch baseline was created.
	CreatedDate *time.Time

	// A description of the patch baseline.
	Description *string

	// A set of global filters used to exclude patches from the baseline.
	GlobalFilters *types.PatchFilterGroup

	// The date the patch baseline was last modified.
	ModifiedDate *time.Time

	// The name of the patch baseline.
	Name *string

	// Returns the operating system specified for the patch baseline.
	OperatingSystem types.OperatingSystem

	// Patch groups included in the patch baseline.
	PatchGroups []string

	// A list of explicitly rejected patches for the baseline.
	RejectedPatches []string

	// The action specified to take on patches included in the RejectedPatches list. A
	// patch can be allowed only if it is a dependency of another package, or blocked
	// entirely along with packages that include it as a dependency.
	RejectedPatchesAction types.PatchAction

	// Information about the patches to use to update the managed nodes, including
	// target operating systems and source repositories. Applies to Linux managed nodes
	// only.
	Sources []types.PatchSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPatchBaselineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPatchBaseline{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetPatchBaseline"); err != nil {
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
	if err = addOpGetPatchBaselineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPatchBaseline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPatchBaseline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetPatchBaseline",
	}
}
