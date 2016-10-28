// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ecriface provides an interface to enable mocking the Amazon EC2 Container Registry service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ecriface

import (
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/service/ecr"
)

// ECRAPI provides an interface to enable mocking the
// ecr.ECR service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon EC2 Container Registry.
//    func myFunc(svc ecriface.ECRAPI) bool {
//        // Make svc.BatchCheckLayerAvailability request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ecr.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockECRClient struct {
//        ecriface.ECRAPI
//    }
//    func (m *mockECRClient) BatchCheckLayerAvailability(input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockECRClient{}
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
type ECRAPI interface {
	BatchCheckLayerAvailabilityRequest(*ecr.BatchCheckLayerAvailabilityInput) (*request.Request, *ecr.BatchCheckLayerAvailabilityOutput)

	BatchCheckLayerAvailability(*ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error)

	BatchDeleteImageRequest(*ecr.BatchDeleteImageInput) (*request.Request, *ecr.BatchDeleteImageOutput)

	BatchDeleteImage(*ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error)

	BatchGetImageRequest(*ecr.BatchGetImageInput) (*request.Request, *ecr.BatchGetImageOutput)

	BatchGetImage(*ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error)

	CompleteLayerUploadRequest(*ecr.CompleteLayerUploadInput) (*request.Request, *ecr.CompleteLayerUploadOutput)

	CompleteLayerUpload(*ecr.CompleteLayerUploadInput) (*ecr.CompleteLayerUploadOutput, error)

	CreateRepositoryRequest(*ecr.CreateRepositoryInput) (*request.Request, *ecr.CreateRepositoryOutput)

	CreateRepository(*ecr.CreateRepositoryInput) (*ecr.CreateRepositoryOutput, error)

	DeleteRepositoryRequest(*ecr.DeleteRepositoryInput) (*request.Request, *ecr.DeleteRepositoryOutput)

	DeleteRepository(*ecr.DeleteRepositoryInput) (*ecr.DeleteRepositoryOutput, error)

	DeleteRepositoryPolicyRequest(*ecr.DeleteRepositoryPolicyInput) (*request.Request, *ecr.DeleteRepositoryPolicyOutput)

	DeleteRepositoryPolicy(*ecr.DeleteRepositoryPolicyInput) (*ecr.DeleteRepositoryPolicyOutput, error)

	DescribeImagesRequest(*ecr.DescribeImagesInput) (*request.Request, *ecr.DescribeImagesOutput)

	DescribeImages(*ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error)

	DescribeRepositoriesRequest(*ecr.DescribeRepositoriesInput) (*request.Request, *ecr.DescribeRepositoriesOutput)

	DescribeRepositories(*ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error)

	GetAuthorizationTokenRequest(*ecr.GetAuthorizationTokenInput) (*request.Request, *ecr.GetAuthorizationTokenOutput)

	GetAuthorizationToken(*ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error)

	GetDownloadUrlForLayerRequest(*ecr.GetDownloadUrlForLayerInput) (*request.Request, *ecr.GetDownloadUrlForLayerOutput)

	GetDownloadUrlForLayer(*ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error)

	GetRepositoryPolicyRequest(*ecr.GetRepositoryPolicyInput) (*request.Request, *ecr.GetRepositoryPolicyOutput)

	GetRepositoryPolicy(*ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error)

	InitiateLayerUploadRequest(*ecr.InitiateLayerUploadInput) (*request.Request, *ecr.InitiateLayerUploadOutput)

	InitiateLayerUpload(*ecr.InitiateLayerUploadInput) (*ecr.InitiateLayerUploadOutput, error)

	ListImagesRequest(*ecr.ListImagesInput) (*request.Request, *ecr.ListImagesOutput)

	ListImages(*ecr.ListImagesInput) (*ecr.ListImagesOutput, error)

	PutImageRequest(*ecr.PutImageInput) (*request.Request, *ecr.PutImageOutput)

	PutImage(*ecr.PutImageInput) (*ecr.PutImageOutput, error)

	SetRepositoryPolicyRequest(*ecr.SetRepositoryPolicyInput) (*request.Request, *ecr.SetRepositoryPolicyOutput)

	SetRepositoryPolicy(*ecr.SetRepositoryPolicyInput) (*ecr.SetRepositoryPolicyOutput, error)

	UploadLayerPartRequest(*ecr.UploadLayerPartInput) (*request.Request, *ecr.UploadLayerPartOutput)

	UploadLayerPart(*ecr.UploadLayerPartInput) (*ecr.UploadLayerPartOutput, error)
}

var _ ECRAPI = (*ecr.ECR)(nil)
