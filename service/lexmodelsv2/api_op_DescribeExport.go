// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a specific export.
func (c *Client) DescribeExport(ctx context.Context, params *DescribeExportInput, optFns ...func(*Options)) (*DescribeExportOutput, error) {
	if params == nil {
		params = &DescribeExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExport", params, optFns, c.addOperationDescribeExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportInput struct {

	// The unique identifier of the export to describe.
	//
	// This member is required.
	ExportId *string

	noSmithyDocumentSerde
}

type DescribeExportOutput struct {

	// The date and time that the export was created.
	CreationDateTime *time.Time

	// A pre-signed S3 URL that points to the bot or bot locale archive. The URL is
	// only available for 5 minutes after calling the DescribeExport operation.
	DownloadUrl *string

	// The unique identifier of the described export.
	ExportId *string

	// The status of the export. When the status is Complete the export archive file is
	// available for download.
	ExportStatus types.ExportStatus

	// If the exportStatus is failed, contains one or more reasons why the export could
	// not be completed.
	FailureReasons []string

	// The file format used in the files that describe the bot or bot locale.
	FileFormat types.ImportExportFileFormat

	// The last date and time that the export was updated.
	LastUpdatedDateTime *time.Time

	// The bot, bot ID, and optional locale ID of the exported bot or bot locale.
	ResourceSpecification *types.ExportResourceSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeExport{}, middleware.After)
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
	if err = addOpDescribeExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExport(options.Region), middleware.Before); err != nil {
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

// DescribeExportAPIClient is a client that implements the DescribeExport
// operation.
type DescribeExportAPIClient interface {
	DescribeExport(context.Context, *DescribeExportInput, ...func(*Options)) (*DescribeExportOutput, error)
}

var _ DescribeExportAPIClient = (*Client)(nil)

// BotExportCompletedWaiterOptions are waiter options for BotExportCompletedWaiter
type BotExportCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BotExportCompletedWaiter will use default minimum delay of 10 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, BotExportCompletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
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
	Retryable func(context.Context, *DescribeExportInput, *DescribeExportOutput, error) (bool, error)
}

// BotExportCompletedWaiter defines the waiters for BotExportCompleted
type BotExportCompletedWaiter struct {
	client DescribeExportAPIClient

	options BotExportCompletedWaiterOptions
}

// NewBotExportCompletedWaiter constructs a BotExportCompletedWaiter.
func NewBotExportCompletedWaiter(client DescribeExportAPIClient, optFns ...func(*BotExportCompletedWaiterOptions)) *BotExportCompletedWaiter {
	options := BotExportCompletedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = botExportCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BotExportCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BotExportCompleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *BotExportCompletedWaiter) Wait(ctx context.Context, params *DescribeExportInput, maxWaitDur time.Duration, optFns ...func(*BotExportCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BotExportCompleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *BotExportCompletedWaiter) WaitForOutput(ctx context.Context, params *DescribeExportInput, maxWaitDur time.Duration, optFns ...func(*BotExportCompletedWaiterOptions)) (*DescribeExportOutput, error) {
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

		out, err := w.client.DescribeExport(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for BotExportCompleted waiter")
}

func botExportCompletedStateRetryable(ctx context.Context, input *DescribeExportInput, output *DescribeExportOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("exportStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.ExportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ExportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("exportStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Deleting"
		value, ok := pathValue.(types.ExportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ExportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("exportStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.ExportStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ExportStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeExport",
	}
}
