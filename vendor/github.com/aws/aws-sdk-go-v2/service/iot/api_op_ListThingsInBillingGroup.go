// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the things you have added to the given billing group. Requires permission
// to access the ListThingsInBillingGroup (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListThingsInBillingGroup(ctx context.Context, params *ListThingsInBillingGroupInput, optFns ...func(*Options)) (*ListThingsInBillingGroupOutput, error) {
	if params == nil {
		params = &ListThingsInBillingGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingsInBillingGroup", params, optFns, c.addOperationListThingsInBillingGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingsInBillingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingsInBillingGroupInput struct {

	// The name of the billing group.
	//
	// This member is required.
	BillingGroupName *string

	// The maximum number of results to return per request.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListThingsInBillingGroupOutput struct {

	// The token to use to get the next set of results. Will not be returned if
	// operation has returned all results.
	NextToken *string

	// A list of things in the billing group.
	Things []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingsInBillingGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingsInBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingsInBillingGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListThingsInBillingGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingsInBillingGroup(options.Region), middleware.Before); err != nil {
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
	return nil
}

// ListThingsInBillingGroupAPIClient is a client that implements the
// ListThingsInBillingGroup operation.
type ListThingsInBillingGroupAPIClient interface {
	ListThingsInBillingGroup(context.Context, *ListThingsInBillingGroupInput, ...func(*Options)) (*ListThingsInBillingGroupOutput, error)
}

var _ ListThingsInBillingGroupAPIClient = (*Client)(nil)

// ListThingsInBillingGroupPaginatorOptions is the paginator options for
// ListThingsInBillingGroup
type ListThingsInBillingGroupPaginatorOptions struct {
	// The maximum number of results to return per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingsInBillingGroupPaginator is a paginator for ListThingsInBillingGroup
type ListThingsInBillingGroupPaginator struct {
	options   ListThingsInBillingGroupPaginatorOptions
	client    ListThingsInBillingGroupAPIClient
	params    *ListThingsInBillingGroupInput
	nextToken *string
	firstPage bool
}

// NewListThingsInBillingGroupPaginator returns a new
// ListThingsInBillingGroupPaginator
func NewListThingsInBillingGroupPaginator(client ListThingsInBillingGroupAPIClient, params *ListThingsInBillingGroupInput, optFns ...func(*ListThingsInBillingGroupPaginatorOptions)) *ListThingsInBillingGroupPaginator {
	if params == nil {
		params = &ListThingsInBillingGroupInput{}
	}

	options := ListThingsInBillingGroupPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingsInBillingGroupPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingsInBillingGroupPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingsInBillingGroup page.
func (p *ListThingsInBillingGroupPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingsInBillingGroupOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListThingsInBillingGroup(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListThingsInBillingGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingsInBillingGroup",
	}
}