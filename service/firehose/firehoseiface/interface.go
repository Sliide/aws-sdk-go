// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package firehoseiface provides an interface to enable mocking the Amazon Kinesis Firehose service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package firehoseiface

import (
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/service/firehose"
)

// FirehoseAPI provides an interface to enable mocking the
// firehose.Firehose service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Firehose.
//    func myFunc(svc firehoseiface.FirehoseAPI) bool {
//        // Make svc.CreateDeliveryStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := firehose.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFirehoseClient struct {
//        firehoseiface.FirehoseAPI
//    }
//    func (m *mockFirehoseClient) CreateDeliveryStream(input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFirehoseClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type FirehoseAPI interface {
	CreateDeliveryStreamRequest(*firehose.CreateDeliveryStreamInput) (*request.Request, *firehose.CreateDeliveryStreamOutput)

	CreateDeliveryStream(*firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error)

	DeleteDeliveryStreamRequest(*firehose.DeleteDeliveryStreamInput) (*request.Request, *firehose.DeleteDeliveryStreamOutput)

	DeleteDeliveryStream(*firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error)

	DescribeDeliveryStreamRequest(*firehose.DescribeDeliveryStreamInput) (*request.Request, *firehose.DescribeDeliveryStreamOutput)

	DescribeDeliveryStream(*firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error)

	ListDeliveryStreamsRequest(*firehose.ListDeliveryStreamsInput) (*request.Request, *firehose.ListDeliveryStreamsOutput)

	ListDeliveryStreams(*firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error)

	PutRecordRequest(*firehose.PutRecordInput) (*request.Request, *firehose.PutRecordOutput)

	PutRecord(*firehose.PutRecordInput) (*firehose.PutRecordOutput, error)

	PutRecordBatchRequest(*firehose.PutRecordBatchInput) (*request.Request, *firehose.PutRecordBatchOutput)

	PutRecordBatch(*firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error)

	UpdateDestinationRequest(*firehose.UpdateDestinationInput) (*request.Request, *firehose.UpdateDestinationOutput)

	UpdateDestination(*firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error)
}

var _ FirehoseAPI = (*firehose.Firehose)(nil)
