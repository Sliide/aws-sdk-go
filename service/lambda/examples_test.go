// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sliide/aws-sdk-go/aws"
	"github.com/sliide/aws-sdk-go/aws/session"
	"github.com/sliide/aws-sdk-go/service/lambda"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLambda_AddPermission() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.AddPermissionInput{
		Action:           aws.String("Action"),       // Required
		FunctionName:     aws.String("FunctionName"), // Required
		Principal:        aws.String("Principal"),    // Required
		StatementId:      aws.String("StatementId"),  // Required
		EventSourceToken: aws.String("EventSourceToken"),
		Qualifier:        aws.String("Qualifier"),
		SourceAccount:    aws.String("SourceOwner"),
		SourceArn:        aws.String("Arn"),
	}
	resp, err := svc.AddPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.CreateAliasInput{
		FunctionName:    aws.String("FunctionName"), // Required
		FunctionVersion: aws.String("Version"),      // Required
		Name:            aws.String("Alias"),        // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateEventSourceMapping() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.CreateEventSourceMappingInput{
		EventSourceArn:   aws.String("Arn"),                 // Required
		FunctionName:     aws.String("FunctionName"),        // Required
		StartingPosition: aws.String("EventSourcePosition"), // Required
		BatchSize:        aws.Int64(1),
		Enabled:          aws.Bool(true),
	}
	resp, err := svc.CreateEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateFunction() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{ // Required
			S3Bucket:        aws.String("S3Bucket"),
			S3Key:           aws.String("S3Key"),
			S3ObjectVersion: aws.String("S3ObjectVersion"),
			ZipFile:         []byte("PAYLOAD"),
		},
		FunctionName: aws.String("FunctionName"), // Required
		Handler:      aws.String("Handler"),      // Required
		Role:         aws.String("RoleArn"),      // Required
		Runtime:      aws.String("Runtime"),      // Required
		Description:  aws.String("Description"),
		MemorySize:   aws.Int64(1),
		Publish:      aws.Bool(true),
		Timeout:      aws.Int64(1),
		VpcConfig: &lambda.VpcConfig{
			SecurityGroupIds: []*string{
				aws.String("SecurityGroupId"), // Required
				// More values...
			},
			SubnetIds: []*string{
				aws.String("SubnetId"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.CreateFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.DeleteAliasInput{
		FunctionName: aws.String("FunctionName"), // Required
		Name:         aws.String("Alias"),        // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteEventSourceMapping() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.DeleteEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteFunction() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.DeleteFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.DeleteFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.GetAliasInput{
		FunctionName: aws.String("FunctionName"), // Required
		Name:         aws.String("Alias"),        // Required
	}
	resp, err := svc.GetAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetEventSourceMapping() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.GetEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.GetEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetFunction() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.GetFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetFunctionConfiguration() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetFunctionConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetPolicy() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.GetPolicyInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_Invoke() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.InvokeInput{
		FunctionName:   aws.String("FunctionName"), // Required
		ClientContext:  aws.String("String"),
		InvocationType: aws.String("InvocationType"),
		LogType:        aws.String("LogType"),
		Payload:        []byte("PAYLOAD"),
		Qualifier:      aws.String("Qualifier"),
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_InvokeAsync() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.InvokeAsyncInput{
		FunctionName: aws.String("FunctionName"),         // Required
		InvokeArgs:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.InvokeAsync(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListAliases() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.ListAliasesInput{
		FunctionName:    aws.String("FunctionName"), // Required
		FunctionVersion: aws.String("Version"),
		Marker:          aws.String("String"),
		MaxItems:        aws.Int64(1),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListEventSourceMappings() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.ListEventSourceMappingsInput{
		EventSourceArn: aws.String("Arn"),
		FunctionName:   aws.String("FunctionName"),
		Marker:         aws.String("String"),
		MaxItems:       aws.Int64(1),
	}
	resp, err := svc.ListEventSourceMappings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListFunctions() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.ListFunctionsInput{
		Marker:   aws.String("String"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListFunctions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListVersionsByFunction() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.ListVersionsByFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Marker:       aws.String("String"),
		MaxItems:     aws.Int64(1),
	}
	resp, err := svc.ListVersionsByFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_PublishVersion() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.PublishVersionInput{
		FunctionName: aws.String("FunctionName"), // Required
		CodeSha256:   aws.String("String"),
		Description:  aws.String("Description"),
	}
	resp, err := svc.PublishVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_RemovePermission() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.RemovePermissionInput{
		FunctionName: aws.String("FunctionName"), // Required
		StatementId:  aws.String("StatementId"),  // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.RemovePermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateAlias() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.UpdateAliasInput{
		FunctionName:    aws.String("FunctionName"), // Required
		Name:            aws.String("Alias"),        // Required
		Description:     aws.String("Description"),
		FunctionVersion: aws.String("Version"),
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateEventSourceMapping() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.UpdateEventSourceMappingInput{
		UUID:         aws.String("String"), // Required
		BatchSize:    aws.Int64(1),
		Enabled:      aws.Bool(true),
		FunctionName: aws.String("FunctionName"),
	}
	resp, err := svc.UpdateEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateFunctionCode() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.UpdateFunctionCodeInput{
		FunctionName:    aws.String("FunctionName"), // Required
		Publish:         aws.Bool(true),
		S3Bucket:        aws.String("S3Bucket"),
		S3Key:           aws.String("S3Key"),
		S3ObjectVersion: aws.String("S3ObjectVersion"),
		ZipFile:         []byte("PAYLOAD"),
	}
	resp, err := svc.UpdateFunctionCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateFunctionConfiguration() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := lambda.New(sess)

	params := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
		Description:  aws.String("Description"),
		Handler:      aws.String("Handler"),
		MemorySize:   aws.Int64(1),
		Role:         aws.String("RoleArn"),
		Runtime:      aws.String("Runtime"),
		Timeout:      aws.Int64(1),
		VpcConfig: &lambda.VpcConfig{
			SecurityGroupIds: []*string{
				aws.String("SecurityGroupId"), // Required
				// More values...
			},
			SubnetIds: []*string{
				aws.String("SubnetId"), // Required
				// More values...
			},
		},
	}
	resp, err := svc.UpdateFunctionConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
