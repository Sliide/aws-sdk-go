// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package marketplacemetering_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sliide/aws-sdk-go/aws"
	"github.com/sliide/aws-sdk-go/aws/session"
	"github.com/sliide/aws-sdk-go/service/marketplacemetering"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMarketplaceMetering_MeterUsage() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := marketplacemetering.New(sess)

	params := &marketplacemetering.MeterUsageInput{
		DryRun:         aws.Bool(true),               // Required
		ProductCode:    aws.String("ProductCode"),    // Required
		Timestamp:      aws.Time(time.Now()),         // Required
		UsageDimension: aws.String("UsageDimension"), // Required
		UsageQuantity:  aws.Int64(1),                 // Required
	}
	resp, err := svc.MeterUsage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
