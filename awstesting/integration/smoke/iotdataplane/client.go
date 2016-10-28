// +build integration

//Package iotdataplane provides gucumber integration tests support.
package iotdataplane

import (
	"github.com/sliide/aws-sdk-go/aws"
	"github.com/sliide/aws-sdk-go/awstesting/integration/smoke"
	"github.com/sliide/aws-sdk-go/service/iot"
	"github.com/sliide/aws-sdk-go/service/iotdataplane"
	"github.com/gucumber/gucumber"
)

func init() {
	gucumber.Before("@iotdataplane", func() {
		svc := iot.New(smoke.Session)
		result, err := svc.DescribeEndpoint(&iot.DescribeEndpointInput{})
		if err != nil {
			gucumber.World["error"] = err
			return
		}

		gucumber.World["client"] = iotdataplane.New(smoke.Session, aws.NewConfig().
			WithEndpoint(*result.EndpointAddress))
	})
}
