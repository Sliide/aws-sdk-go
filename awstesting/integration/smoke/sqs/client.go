// +build integration

//Package sqs provides gucumber integration tests support.
package sqs

import (
	"github.com/sliide/aws-sdk-go/awstesting/integration/smoke"
	"github.com/sliide/aws-sdk-go/service/sqs"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@sqs", func() {
		gucumber.World["client"] = sqs.New(smoke.Session)
	})
}
