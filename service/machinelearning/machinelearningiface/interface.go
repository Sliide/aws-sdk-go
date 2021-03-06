// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package machinelearningiface provides an interface to enable mocking the Amazon Machine Learning service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package machinelearningiface

import (
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/service/machinelearning"
)

// MachineLearningAPI provides an interface to enable mocking the
// machinelearning.MachineLearning service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Machine Learning.
//    func myFunc(svc machinelearningiface.MachineLearningAPI) bool {
//        // Make svc.AddTags request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := machinelearning.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMachineLearningClient struct {
//        machinelearningiface.MachineLearningAPI
//    }
//    func (m *mockMachineLearningClient) AddTags(input *machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMachineLearningClient{}
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
type MachineLearningAPI interface {
	AddTagsRequest(*machinelearning.AddTagsInput) (*request.Request, *machinelearning.AddTagsOutput)

	AddTags(*machinelearning.AddTagsInput) (*machinelearning.AddTagsOutput, error)

	CreateBatchPredictionRequest(*machinelearning.CreateBatchPredictionInput) (*request.Request, *machinelearning.CreateBatchPredictionOutput)

	CreateBatchPrediction(*machinelearning.CreateBatchPredictionInput) (*machinelearning.CreateBatchPredictionOutput, error)

	CreateDataSourceFromRDSRequest(*machinelearning.CreateDataSourceFromRDSInput) (*request.Request, *machinelearning.CreateDataSourceFromRDSOutput)

	CreateDataSourceFromRDS(*machinelearning.CreateDataSourceFromRDSInput) (*machinelearning.CreateDataSourceFromRDSOutput, error)

	CreateDataSourceFromRedshiftRequest(*machinelearning.CreateDataSourceFromRedshiftInput) (*request.Request, *machinelearning.CreateDataSourceFromRedshiftOutput)

	CreateDataSourceFromRedshift(*machinelearning.CreateDataSourceFromRedshiftInput) (*machinelearning.CreateDataSourceFromRedshiftOutput, error)

	CreateDataSourceFromS3Request(*machinelearning.CreateDataSourceFromS3Input) (*request.Request, *machinelearning.CreateDataSourceFromS3Output)

	CreateDataSourceFromS3(*machinelearning.CreateDataSourceFromS3Input) (*machinelearning.CreateDataSourceFromS3Output, error)

	CreateEvaluationRequest(*machinelearning.CreateEvaluationInput) (*request.Request, *machinelearning.CreateEvaluationOutput)

	CreateEvaluation(*machinelearning.CreateEvaluationInput) (*machinelearning.CreateEvaluationOutput, error)

	CreateMLModelRequest(*machinelearning.CreateMLModelInput) (*request.Request, *machinelearning.CreateMLModelOutput)

	CreateMLModel(*machinelearning.CreateMLModelInput) (*machinelearning.CreateMLModelOutput, error)

	CreateRealtimeEndpointRequest(*machinelearning.CreateRealtimeEndpointInput) (*request.Request, *machinelearning.CreateRealtimeEndpointOutput)

	CreateRealtimeEndpoint(*machinelearning.CreateRealtimeEndpointInput) (*machinelearning.CreateRealtimeEndpointOutput, error)

	DeleteBatchPredictionRequest(*machinelearning.DeleteBatchPredictionInput) (*request.Request, *machinelearning.DeleteBatchPredictionOutput)

	DeleteBatchPrediction(*machinelearning.DeleteBatchPredictionInput) (*machinelearning.DeleteBatchPredictionOutput, error)

	DeleteDataSourceRequest(*machinelearning.DeleteDataSourceInput) (*request.Request, *machinelearning.DeleteDataSourceOutput)

	DeleteDataSource(*machinelearning.DeleteDataSourceInput) (*machinelearning.DeleteDataSourceOutput, error)

	DeleteEvaluationRequest(*machinelearning.DeleteEvaluationInput) (*request.Request, *machinelearning.DeleteEvaluationOutput)

	DeleteEvaluation(*machinelearning.DeleteEvaluationInput) (*machinelearning.DeleteEvaluationOutput, error)

	DeleteMLModelRequest(*machinelearning.DeleteMLModelInput) (*request.Request, *machinelearning.DeleteMLModelOutput)

	DeleteMLModel(*machinelearning.DeleteMLModelInput) (*machinelearning.DeleteMLModelOutput, error)

	DeleteRealtimeEndpointRequest(*machinelearning.DeleteRealtimeEndpointInput) (*request.Request, *machinelearning.DeleteRealtimeEndpointOutput)

	DeleteRealtimeEndpoint(*machinelearning.DeleteRealtimeEndpointInput) (*machinelearning.DeleteRealtimeEndpointOutput, error)

	DeleteTagsRequest(*machinelearning.DeleteTagsInput) (*request.Request, *machinelearning.DeleteTagsOutput)

	DeleteTags(*machinelearning.DeleteTagsInput) (*machinelearning.DeleteTagsOutput, error)

	DescribeBatchPredictionsRequest(*machinelearning.DescribeBatchPredictionsInput) (*request.Request, *machinelearning.DescribeBatchPredictionsOutput)

	DescribeBatchPredictions(*machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error)

	DescribeBatchPredictionsPages(*machinelearning.DescribeBatchPredictionsInput, func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool) error

	DescribeDataSourcesRequest(*machinelearning.DescribeDataSourcesInput) (*request.Request, *machinelearning.DescribeDataSourcesOutput)

	DescribeDataSources(*machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error)

	DescribeDataSourcesPages(*machinelearning.DescribeDataSourcesInput, func(*machinelearning.DescribeDataSourcesOutput, bool) bool) error

	DescribeEvaluationsRequest(*machinelearning.DescribeEvaluationsInput) (*request.Request, *machinelearning.DescribeEvaluationsOutput)

	DescribeEvaluations(*machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error)

	DescribeEvaluationsPages(*machinelearning.DescribeEvaluationsInput, func(*machinelearning.DescribeEvaluationsOutput, bool) bool) error

	DescribeMLModelsRequest(*machinelearning.DescribeMLModelsInput) (*request.Request, *machinelearning.DescribeMLModelsOutput)

	DescribeMLModels(*machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error)

	DescribeMLModelsPages(*machinelearning.DescribeMLModelsInput, func(*machinelearning.DescribeMLModelsOutput, bool) bool) error

	DescribeTagsRequest(*machinelearning.DescribeTagsInput) (*request.Request, *machinelearning.DescribeTagsOutput)

	DescribeTags(*machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error)

	GetBatchPredictionRequest(*machinelearning.GetBatchPredictionInput) (*request.Request, *machinelearning.GetBatchPredictionOutput)

	GetBatchPrediction(*machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error)

	GetDataSourceRequest(*machinelearning.GetDataSourceInput) (*request.Request, *machinelearning.GetDataSourceOutput)

	GetDataSource(*machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error)

	GetEvaluationRequest(*machinelearning.GetEvaluationInput) (*request.Request, *machinelearning.GetEvaluationOutput)

	GetEvaluation(*machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error)

	GetMLModelRequest(*machinelearning.GetMLModelInput) (*request.Request, *machinelearning.GetMLModelOutput)

	GetMLModel(*machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error)

	PredictRequest(*machinelearning.PredictInput) (*request.Request, *machinelearning.PredictOutput)

	Predict(*machinelearning.PredictInput) (*machinelearning.PredictOutput, error)

	UpdateBatchPredictionRequest(*machinelearning.UpdateBatchPredictionInput) (*request.Request, *machinelearning.UpdateBatchPredictionOutput)

	UpdateBatchPrediction(*machinelearning.UpdateBatchPredictionInput) (*machinelearning.UpdateBatchPredictionOutput, error)

	UpdateDataSourceRequest(*machinelearning.UpdateDataSourceInput) (*request.Request, *machinelearning.UpdateDataSourceOutput)

	UpdateDataSource(*machinelearning.UpdateDataSourceInput) (*machinelearning.UpdateDataSourceOutput, error)

	UpdateEvaluationRequest(*machinelearning.UpdateEvaluationInput) (*request.Request, *machinelearning.UpdateEvaluationOutput)

	UpdateEvaluation(*machinelearning.UpdateEvaluationInput) (*machinelearning.UpdateEvaluationOutput, error)

	UpdateMLModelRequest(*machinelearning.UpdateMLModelInput) (*request.Request, *machinelearning.UpdateMLModelOutput)

	UpdateMLModel(*machinelearning.UpdateMLModelInput) (*machinelearning.UpdateMLModelOutput, error)

	WaitUntilBatchPredictionAvailable(*machinelearning.DescribeBatchPredictionsInput) error

	WaitUntilDataSourceAvailable(*machinelearning.DescribeDataSourcesInput) error

	WaitUntilEvaluationAvailable(*machinelearning.DescribeEvaluationsInput) error

	WaitUntilMLModelAvailable(*machinelearning.DescribeMLModelsInput) error
}

var _ MachineLearningAPI = (*machinelearning.MachineLearning)(nil)
