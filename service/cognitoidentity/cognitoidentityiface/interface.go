// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitoidentityiface provides an interface to enable mocking the Amazon Cognito Identity service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cognitoidentityiface

import (
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/service/cognitoidentity"
)

// CognitoIdentityAPI provides an interface to enable mocking the
// cognitoidentity.CognitoIdentity service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Cognito Identity.
//    func myFunc(svc cognitoidentityiface.CognitoIdentityAPI) bool {
//        // Make svc.CreateIdentityPool request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cognitoidentity.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCognitoIdentityClient struct {
//        cognitoidentityiface.CognitoIdentityAPI
//    }
//    func (m *mockCognitoIdentityClient) CreateIdentityPool(input *cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCognitoIdentityClient{}
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
type CognitoIdentityAPI interface {
	CreateIdentityPoolRequest(*cognitoidentity.CreateIdentityPoolInput) (*request.Request, *cognitoidentity.IdentityPool)

	CreateIdentityPool(*cognitoidentity.CreateIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	DeleteIdentitiesRequest(*cognitoidentity.DeleteIdentitiesInput) (*request.Request, *cognitoidentity.DeleteIdentitiesOutput)

	DeleteIdentities(*cognitoidentity.DeleteIdentitiesInput) (*cognitoidentity.DeleteIdentitiesOutput, error)

	DeleteIdentityPoolRequest(*cognitoidentity.DeleteIdentityPoolInput) (*request.Request, *cognitoidentity.DeleteIdentityPoolOutput)

	DeleteIdentityPool(*cognitoidentity.DeleteIdentityPoolInput) (*cognitoidentity.DeleteIdentityPoolOutput, error)

	DescribeIdentityRequest(*cognitoidentity.DescribeIdentityInput) (*request.Request, *cognitoidentity.IdentityDescription)

	DescribeIdentity(*cognitoidentity.DescribeIdentityInput) (*cognitoidentity.IdentityDescription, error)

	DescribeIdentityPoolRequest(*cognitoidentity.DescribeIdentityPoolInput) (*request.Request, *cognitoidentity.IdentityPool)

	DescribeIdentityPool(*cognitoidentity.DescribeIdentityPoolInput) (*cognitoidentity.IdentityPool, error)

	GetCredentialsForIdentityRequest(*cognitoidentity.GetCredentialsForIdentityInput) (*request.Request, *cognitoidentity.GetCredentialsForIdentityOutput)

	GetCredentialsForIdentity(*cognitoidentity.GetCredentialsForIdentityInput) (*cognitoidentity.GetCredentialsForIdentityOutput, error)

	GetIdRequest(*cognitoidentity.GetIdInput) (*request.Request, *cognitoidentity.GetIdOutput)

	GetId(*cognitoidentity.GetIdInput) (*cognitoidentity.GetIdOutput, error)

	GetIdentityPoolRolesRequest(*cognitoidentity.GetIdentityPoolRolesInput) (*request.Request, *cognitoidentity.GetIdentityPoolRolesOutput)

	GetIdentityPoolRoles(*cognitoidentity.GetIdentityPoolRolesInput) (*cognitoidentity.GetIdentityPoolRolesOutput, error)

	GetOpenIdTokenRequest(*cognitoidentity.GetOpenIdTokenInput) (*request.Request, *cognitoidentity.GetOpenIdTokenOutput)

	GetOpenIdToken(*cognitoidentity.GetOpenIdTokenInput) (*cognitoidentity.GetOpenIdTokenOutput, error)

	GetOpenIdTokenForDeveloperIdentityRequest(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*request.Request, *cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput)

	GetOpenIdTokenForDeveloperIdentity(*cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput) (*cognitoidentity.GetOpenIdTokenForDeveloperIdentityOutput, error)

	ListIdentitiesRequest(*cognitoidentity.ListIdentitiesInput) (*request.Request, *cognitoidentity.ListIdentitiesOutput)

	ListIdentities(*cognitoidentity.ListIdentitiesInput) (*cognitoidentity.ListIdentitiesOutput, error)

	ListIdentityPoolsRequest(*cognitoidentity.ListIdentityPoolsInput) (*request.Request, *cognitoidentity.ListIdentityPoolsOutput)

	ListIdentityPools(*cognitoidentity.ListIdentityPoolsInput) (*cognitoidentity.ListIdentityPoolsOutput, error)

	LookupDeveloperIdentityRequest(*cognitoidentity.LookupDeveloperIdentityInput) (*request.Request, *cognitoidentity.LookupDeveloperIdentityOutput)

	LookupDeveloperIdentity(*cognitoidentity.LookupDeveloperIdentityInput) (*cognitoidentity.LookupDeveloperIdentityOutput, error)

	MergeDeveloperIdentitiesRequest(*cognitoidentity.MergeDeveloperIdentitiesInput) (*request.Request, *cognitoidentity.MergeDeveloperIdentitiesOutput)

	MergeDeveloperIdentities(*cognitoidentity.MergeDeveloperIdentitiesInput) (*cognitoidentity.MergeDeveloperIdentitiesOutput, error)

	SetIdentityPoolRolesRequest(*cognitoidentity.SetIdentityPoolRolesInput) (*request.Request, *cognitoidentity.SetIdentityPoolRolesOutput)

	SetIdentityPoolRoles(*cognitoidentity.SetIdentityPoolRolesInput) (*cognitoidentity.SetIdentityPoolRolesOutput, error)

	UnlinkDeveloperIdentityRequest(*cognitoidentity.UnlinkDeveloperIdentityInput) (*request.Request, *cognitoidentity.UnlinkDeveloperIdentityOutput)

	UnlinkDeveloperIdentity(*cognitoidentity.UnlinkDeveloperIdentityInput) (*cognitoidentity.UnlinkDeveloperIdentityOutput, error)

	UnlinkIdentityRequest(*cognitoidentity.UnlinkIdentityInput) (*request.Request, *cognitoidentity.UnlinkIdentityOutput)

	UnlinkIdentity(*cognitoidentity.UnlinkIdentityInput) (*cognitoidentity.UnlinkIdentityOutput, error)

	UpdateIdentityPoolRequest(*cognitoidentity.IdentityPool) (*request.Request, *cognitoidentity.IdentityPool)

	UpdateIdentityPool(*cognitoidentity.IdentityPool) (*cognitoidentity.IdentityPool, error)
}

var _ CognitoIdentityAPI = (*cognitoidentity.CognitoIdentity)(nil)
