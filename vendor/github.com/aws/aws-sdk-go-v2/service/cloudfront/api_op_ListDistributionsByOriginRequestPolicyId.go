// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of distribution IDs for distributions that have a cache behavior
// that's associated with the specified origin request policy.
//
// You can optionally specify the maximum number of items to receive in the
// response. If the total number of items in the list exceeds the maximum that you
// specify, or the default maximum, the response is paginated. To get the next page
// of items, send a subsequent request that specifies the NextMarker value from
// the current response as the Marker value in the subsequent request.
func (c *Client) ListDistributionsByOriginRequestPolicyId(ctx context.Context, params *ListDistributionsByOriginRequestPolicyIdInput, optFns ...func(*Options)) (*ListDistributionsByOriginRequestPolicyIdOutput, error) {
	if params == nil {
		params = &ListDistributionsByOriginRequestPolicyIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByOriginRequestPolicyId", params, optFns, c.addOperationListDistributionsByOriginRequestPolicyIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByOriginRequestPolicyIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributionsByOriginRequestPolicyIdInput struct {

	// The ID of the origin request policy whose associated distribution IDs you want
	// to list.
	//
	// This member is required.
	OriginRequestPolicyId *string

	// Use this field when paginating results to indicate where to begin in your list
	// of distribution IDs. The response includes distribution IDs in the list that
	// occur after the marker. To get the next page of the list, set this field's value
	// to the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of distribution IDs that you want in the response.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListDistributionsByOriginRequestPolicyIdOutput struct {

	// A list of distribution IDs.
	DistributionIdList *types.DistributionIdList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsByOriginRequestPolicyIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByOriginRequestPolicyId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByOriginRequestPolicyId{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDistributionsByOriginRequestPolicyId"); err != nil {
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
	if err = addSpanRetryLoop(stack, options); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addCredentialSource(stack, options); err != nil {
		return err
	}
	if err = addOpListDistributionsByOriginRequestPolicyIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByOriginRequestPolicyId(options.Region), middleware.Before); err != nil {
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
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListDistributionsByOriginRequestPolicyId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDistributionsByOriginRequestPolicyId",
	}
}
