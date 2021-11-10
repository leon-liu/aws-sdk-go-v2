// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes one or more of your network interfaces.
func (c *Client) DescribeNetworkInterfaces(ctx context.Context, params *DescribeNetworkInterfacesInput, optFns ...func(*Options)) (*DescribeNetworkInterfacesOutput, error) {
	if params == nil {
		params = &DescribeNetworkInterfacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNetworkInterfaces", params, optFns, c.addOperationDescribeNetworkInterfacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNetworkInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeNetworkInterfaces.
type DescribeNetworkInterfacesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * addresses.private-ip-address - The private IPv4
	// addresses associated with the network interface.
	//
	// * addresses.primary - Whether
	// the private IPv4 address is the primary IP address associated with the network
	// interface.
	//
	// * addresses.association.public-ip - The association ID returned when
	// the network interface was associated with the Elastic IP address (IPv4).
	//
	// *
	// addresses.association.owner-id - The owner ID of the addresses associated with
	// the network interface.
	//
	// * association.association-id - The association ID
	// returned when the network interface was associated with an IPv4 address.
	//
	// *
	// association.allocation-id - The allocation ID returned when you allocated the
	// Elastic IP address (IPv4) for your network interface.
	//
	// * association.ip-owner-id
	// - The owner of the Elastic IP address (IPv4) associated with the network
	// interface.
	//
	// * association.public-ip - The address of the Elastic IP address
	// (IPv4) bound to the network interface.
	//
	// * association.public-dns-name - The
	// public DNS name for the network interface (IPv4).
	//
	// * attachment.attachment-id -
	// The ID of the interface attachment.
	//
	// * attachment.attach-time - The time that
	// the network interface was attached to an instance.
	//
	// *
	// attachment.delete-on-termination - Indicates whether the attachment is deleted
	// when an instance is terminated.
	//
	// * attachment.device-index - The device index to
	// which the network interface is attached.
	//
	// * attachment.instance-id - The ID of
	// the instance to which the network interface is attached.
	//
	// *
	// attachment.instance-owner-id - The owner ID of the instance to which the network
	// interface is attached.
	//
	// * attachment.status - The status of the attachment
	// (attaching | attached | detaching | detached).
	//
	// * availability-zone - The
	// Availability Zone of the network interface.
	//
	// * description - The description of
	// the network interface.
	//
	// * group-id - The ID of a security group associated with
	// the network interface.
	//
	// * group-name - The name of a security group associated
	// with the network interface.
	//
	// * ipv6-addresses.ipv6-address - An IPv6 address
	// associated with the network interface.
	//
	// * mac-address - The MAC address of the
	// network interface.
	//
	// * network-interface-id - The ID of the network interface.
	//
	// *
	// owner-id - The Amazon Web Services account ID of the network interface owner.
	//
	// *
	// private-ip-address - The private IPv4 address or addresses of the network
	// interface.
	//
	// * private-dns-name - The private DNS name of the network interface
	// (IPv4).
	//
	// * requester-id - The alias or Amazon Web Services account ID of the
	// principal or service that created the network interface.
	//
	// * requester-managed -
	// Indicates whether the network interface is being managed by an Amazon Web
	// Service (for example, Amazon Web Services Management Console, Auto Scaling, and
	// so on).
	//
	// * source-dest-check - Indicates whether the network interface performs
	// source/destination checking. A value of true means checking is enabled, and
	// false means checking is disabled. The value must be false for the network
	// interface to perform network address translation (NAT) in your VPC.
	//
	// * status -
	// The status of the network interface. If the network interface is not attached to
	// an instance, the status is available; if a network interface is attached to an
	// instance the status is in-use.
	//
	// * subnet-id - The ID of the subnet for the
	// network interface.
	//
	// * tag: - The key/value combination of a tag assigned to the
	// resource. Use the tag key in the filter name and the tag value as the filter
	// value. For example, to find all resources that have a tag with the key Owner and
	// the value TeamA, specify tag:Owner for the filter name and TeamA for the filter
	// value.
	//
	// * tag-key - The key of a tag assigned to the resource. Use this filter
	// to find all resources assigned a tag with a specific key, regardless of the tag
	// value.
	//
	// * vpc-id - The ID of the VPC for the network interface.
	Filters []types.Filter

	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// You cannot specify this parameter and the network interface IDs parameter in the
	// same request.
	MaxResults *int32

	// One or more network interface IDs. Default: Describes all your network
	// interfaces.
	NetworkInterfaceIds []string

	// The token to retrieve the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

// Contains the output of DescribeNetworkInterfaces.
type DescribeNetworkInterfacesOutput struct {

	// Information about one or more network interfaces.
	NetworkInterfaces []types.NetworkInterface

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNetworkInterfacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeNetworkInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeNetworkInterfaces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNetworkInterfaces(options.Region), middleware.Before); err != nil {
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

// DescribeNetworkInterfacesAPIClient is a client that implements the
// DescribeNetworkInterfaces operation.
type DescribeNetworkInterfacesAPIClient interface {
	DescribeNetworkInterfaces(context.Context, *DescribeNetworkInterfacesInput, ...func(*Options)) (*DescribeNetworkInterfacesOutput, error)
}

var _ DescribeNetworkInterfacesAPIClient = (*Client)(nil)

// DescribeNetworkInterfacesPaginatorOptions is the paginator options for
// DescribeNetworkInterfaces
type DescribeNetworkInterfacesPaginatorOptions struct {
	// The maximum number of items to return for this request. The request returns a
	// token that you can specify in a subsequent call to get the next set of results.
	// You cannot specify this parameter and the network interface IDs parameter in the
	// same request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNetworkInterfacesPaginator is a paginator for DescribeNetworkInterfaces
type DescribeNetworkInterfacesPaginator struct {
	options   DescribeNetworkInterfacesPaginatorOptions
	client    DescribeNetworkInterfacesAPIClient
	params    *DescribeNetworkInterfacesInput
	nextToken *string
	firstPage bool
}

// NewDescribeNetworkInterfacesPaginator returns a new
// DescribeNetworkInterfacesPaginator
func NewDescribeNetworkInterfacesPaginator(client DescribeNetworkInterfacesAPIClient, params *DescribeNetworkInterfacesInput, optFns ...func(*DescribeNetworkInterfacesPaginatorOptions)) *DescribeNetworkInterfacesPaginator {
	if params == nil {
		params = &DescribeNetworkInterfacesInput{}
	}

	options := DescribeNetworkInterfacesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNetworkInterfacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNetworkInterfacesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeNetworkInterfaces page.
func (p *DescribeNetworkInterfacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNetworkInterfacesOutput, error) {
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

	result, err := p.client.DescribeNetworkInterfaces(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// NetworkInterfaceAvailableWaiterOptions are waiter options for
// NetworkInterfaceAvailableWaiter
type NetworkInterfaceAvailableWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// NetworkInterfaceAvailableWaiter will use default minimum delay of 20 seconds.
	// Note that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, NetworkInterfaceAvailableWaiter will use default max delay of 120
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeNetworkInterfacesInput, *DescribeNetworkInterfacesOutput, error) (bool, error)
}

// NetworkInterfaceAvailableWaiter defines the waiters for
// NetworkInterfaceAvailable
type NetworkInterfaceAvailableWaiter struct {
	client DescribeNetworkInterfacesAPIClient

	options NetworkInterfaceAvailableWaiterOptions
}

// NewNetworkInterfaceAvailableWaiter constructs a NetworkInterfaceAvailableWaiter.
func NewNetworkInterfaceAvailableWaiter(client DescribeNetworkInterfacesAPIClient, optFns ...func(*NetworkInterfaceAvailableWaiterOptions)) *NetworkInterfaceAvailableWaiter {
	options := NetworkInterfaceAvailableWaiterOptions{}
	options.MinDelay = 20 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = networkInterfaceAvailableStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &NetworkInterfaceAvailableWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for NetworkInterfaceAvailable waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *NetworkInterfaceAvailableWaiter) Wait(ctx context.Context, params *DescribeNetworkInterfacesInput, maxWaitDur time.Duration, optFns ...func(*NetworkInterfaceAvailableWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for NetworkInterfaceAvailable waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *NetworkInterfaceAvailableWaiter) WaitForOutput(ctx context.Context, params *DescribeNetworkInterfacesInput, maxWaitDur time.Duration, optFns ...func(*NetworkInterfaceAvailableWaiterOptions)) (*DescribeNetworkInterfacesOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeNetworkInterfaces(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for NetworkInterfaceAvailable waiter")
}

func networkInterfaceAvailableStateRetryable(ctx context.Context, input *DescribeNetworkInterfacesInput, output *DescribeNetworkInterfacesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("NetworkInterfaces[].Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "available"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.NetworkInterfaceStatus)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.NetworkInterfaceStatus value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err != nil {
		var apiErr smithy.APIError
		ok := errors.As(err, &apiErr)
		if !ok {
			return false, fmt.Errorf("expected err to be of type smithy.APIError, got %w", err)
		}

		if "InvalidNetworkInterfaceID.NotFound" == apiErr.ErrorCode() {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeNetworkInterfaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeNetworkInterfaces",
	}
}
