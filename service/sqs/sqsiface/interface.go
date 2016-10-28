// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sqsiface provides an interface to enable mocking the Amazon Simple Queue Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sqsiface

import (
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/service/sqs"
)

// SQSAPI provides an interface to enable mocking the
// sqs.SQS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Simple Queue Service.
//    func myFunc(svc sqsiface.SQSAPI) bool {
//        // Make svc.AddPermission request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sqs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSQSClient struct {
//        sqsiface.SQSAPI
//    }
//    func (m *mockSQSClient) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSQSClient{}
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
type SQSAPI interface {
	AddPermissionRequest(*sqs.AddPermissionInput) (*request.Request, *sqs.AddPermissionOutput)

	AddPermission(*sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)

	ChangeMessageVisibilityRequest(*sqs.ChangeMessageVisibilityInput) (*request.Request, *sqs.ChangeMessageVisibilityOutput)

	ChangeMessageVisibility(*sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)

	ChangeMessageVisibilityBatchRequest(*sqs.ChangeMessageVisibilityBatchInput) (*request.Request, *sqs.ChangeMessageVisibilityBatchOutput)

	ChangeMessageVisibilityBatch(*sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)

	CreateQueueRequest(*sqs.CreateQueueInput) (*request.Request, *sqs.CreateQueueOutput)

	CreateQueue(*sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)

	DeleteMessageRequest(*sqs.DeleteMessageInput) (*request.Request, *sqs.DeleteMessageOutput)

	DeleteMessage(*sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)

	DeleteMessageBatchRequest(*sqs.DeleteMessageBatchInput) (*request.Request, *sqs.DeleteMessageBatchOutput)

	DeleteMessageBatch(*sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)

	DeleteQueueRequest(*sqs.DeleteQueueInput) (*request.Request, *sqs.DeleteQueueOutput)

	DeleteQueue(*sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)

	GetQueueAttributesRequest(*sqs.GetQueueAttributesInput) (*request.Request, *sqs.GetQueueAttributesOutput)

	GetQueueAttributes(*sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)

	GetQueueUrlRequest(*sqs.GetQueueUrlInput) (*request.Request, *sqs.GetQueueUrlOutput)

	GetQueueUrl(*sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)

	ListDeadLetterSourceQueuesRequest(*sqs.ListDeadLetterSourceQueuesInput) (*request.Request, *sqs.ListDeadLetterSourceQueuesOutput)

	ListDeadLetterSourceQueues(*sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)

	ListQueuesRequest(*sqs.ListQueuesInput) (*request.Request, *sqs.ListQueuesOutput)

	ListQueues(*sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)

	PurgeQueueRequest(*sqs.PurgeQueueInput) (*request.Request, *sqs.PurgeQueueOutput)

	PurgeQueue(*sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)

	ReceiveMessageRequest(*sqs.ReceiveMessageInput) (*request.Request, *sqs.ReceiveMessageOutput)

	ReceiveMessage(*sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)

	RemovePermissionRequest(*sqs.RemovePermissionInput) (*request.Request, *sqs.RemovePermissionOutput)

	RemovePermission(*sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)

	SendMessageRequest(*sqs.SendMessageInput) (*request.Request, *sqs.SendMessageOutput)

	SendMessage(*sqs.SendMessageInput) (*sqs.SendMessageOutput, error)

	SendMessageBatchRequest(*sqs.SendMessageBatchInput) (*request.Request, *sqs.SendMessageBatchOutput)

	SendMessageBatch(*sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)

	SetQueueAttributesRequest(*sqs.SetQueueAttributesInput) (*request.Request, *sqs.SetQueueAttributesOutput)

	SetQueueAttributes(*sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
}

var _ SQSAPI = (*sqs.SQS)(nil)
