// +build integration

//Package dynamodbstreams provides gucumber integration tests support.
package dynamodbstreams

import (
	"github.com/sliide/aws-sdk-go/awstesting/integration/smoke"
	"github.com/sliide/aws-sdk-go/service/dynamodbstreams"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@dynamodbstreams", func() {
		gucumber.World["client"] = dynamodbstreams.New(smoke.Session)
	})
}
