// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package budgets_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sliide/aws-sdk-go/aws"
	"github.com/sliide/aws-sdk-go/aws/session"
	"github.com/sliide/aws-sdk-go/service/budgets"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleBudgets_CreateBudget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.CreateBudgetInput{
		AccountId: aws.String("AccountId"), // Required
		Budget: &budgets.Budget{ // Required
			BudgetLimit: &budgets.Spend{ // Required
				Amount: aws.String("NumericValue"),  // Required
				Unit:   aws.String("GenericString"), // Required
			},
			BudgetName: aws.String("BudgetName"), // Required
			BudgetType: aws.String("BudgetType"), // Required
			CostTypes: &budgets.CostTypes{ // Required
				IncludeSubscription: aws.Bool(true), // Required
				IncludeTax:          aws.Bool(true), // Required
				UseBlended:          aws.Bool(true), // Required
			},
			TimePeriod: &budgets.TimePeriod{ // Required
				End:   aws.Time(time.Now()), // Required
				Start: aws.Time(time.Now()), // Required
			},
			TimeUnit: aws.String("TimeUnit"), // Required
			CalculatedSpend: &budgets.CalculatedSpend{
				ActualSpend: &budgets.Spend{ // Required
					Amount: aws.String("NumericValue"),  // Required
					Unit:   aws.String("GenericString"), // Required
				},
				ForecastedSpend: &budgets.Spend{
					Amount: aws.String("NumericValue"),  // Required
					Unit:   aws.String("GenericString"), // Required
				},
			},
			CostFilters: map[string][]*string{
				"Key": { // Required
					aws.String("GenericString"), // Required
					// More values...
				},
				// More values...
			},
		},
		NotificationsWithSubscribers: []*budgets.NotificationWithSubscribers{
			{ // Required
				Notification: &budgets.Notification{ // Required
					ComparisonOperator: aws.String("ComparisonOperator"), // Required
					NotificationType:   aws.String("NotificationType"),   // Required
					Threshold:          aws.Float64(1.0),                 // Required
				},
				Subscribers: []*budgets.Subscriber{ // Required
					{ // Required
						Address:          aws.String("GenericString"),    // Required
						SubscriptionType: aws.String("SubscriptionType"), // Required
					},
					// More values...
				},
			},
			// More values...
		},
	}
	resp, err := svc.CreateBudget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_CreateNotification() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.CreateNotificationInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		Subscribers: []*budgets.Subscriber{ // Required
			{ // Required
				Address:          aws.String("GenericString"),    // Required
				SubscriptionType: aws.String("SubscriptionType"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.CreateNotification(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_CreateSubscriber() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.CreateSubscriberInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		Subscriber: &budgets.Subscriber{ // Required
			Address:          aws.String("GenericString"),    // Required
			SubscriptionType: aws.String("SubscriptionType"), // Required
		},
	}
	resp, err := svc.CreateSubscriber(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DeleteBudget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DeleteBudgetInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
	}
	resp, err := svc.DeleteBudget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DeleteNotification() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DeleteNotificationInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
	}
	resp, err := svc.DeleteNotification(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DeleteSubscriber() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DeleteSubscriberInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		Subscriber: &budgets.Subscriber{ // Required
			Address:          aws.String("GenericString"),    // Required
			SubscriptionType: aws.String("SubscriptionType"), // Required
		},
	}
	resp, err := svc.DeleteSubscriber(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DescribeBudget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DescribeBudgetInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
	}
	resp, err := svc.DescribeBudget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DescribeBudgets() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DescribeBudgetsInput{
		AccountId:  aws.String("AccountId"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("GenericString"),
	}
	resp, err := svc.DescribeBudgets(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DescribeNotificationsForBudget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DescribeNotificationsForBudgetInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("GenericString"),
	}
	resp, err := svc.DescribeNotificationsForBudget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_DescribeSubscribersForNotification() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.DescribeSubscribersForNotificationInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("GenericString"),
	}
	resp, err := svc.DescribeSubscribersForNotification(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_UpdateBudget() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.UpdateBudgetInput{
		AccountId: aws.String("AccountId"), // Required
		NewBudget: &budgets.Budget{ // Required
			BudgetLimit: &budgets.Spend{ // Required
				Amount: aws.String("NumericValue"),  // Required
				Unit:   aws.String("GenericString"), // Required
			},
			BudgetName: aws.String("BudgetName"), // Required
			BudgetType: aws.String("BudgetType"), // Required
			CostTypes: &budgets.CostTypes{ // Required
				IncludeSubscription: aws.Bool(true), // Required
				IncludeTax:          aws.Bool(true), // Required
				UseBlended:          aws.Bool(true), // Required
			},
			TimePeriod: &budgets.TimePeriod{ // Required
				End:   aws.Time(time.Now()), // Required
				Start: aws.Time(time.Now()), // Required
			},
			TimeUnit: aws.String("TimeUnit"), // Required
			CalculatedSpend: &budgets.CalculatedSpend{
				ActualSpend: &budgets.Spend{ // Required
					Amount: aws.String("NumericValue"),  // Required
					Unit:   aws.String("GenericString"), // Required
				},
				ForecastedSpend: &budgets.Spend{
					Amount: aws.String("NumericValue"),  // Required
					Unit:   aws.String("GenericString"), // Required
				},
			},
			CostFilters: map[string][]*string{
				"Key": { // Required
					aws.String("GenericString"), // Required
					// More values...
				},
				// More values...
			},
		},
	}
	resp, err := svc.UpdateBudget(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_UpdateNotification() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.UpdateNotificationInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		NewNotification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		OldNotification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
	}
	resp, err := svc.UpdateNotification(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleBudgets_UpdateSubscriber() {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	svc := budgets.New(sess)

	params := &budgets.UpdateSubscriberInput{
		AccountId:  aws.String("AccountId"),  // Required
		BudgetName: aws.String("BudgetName"), // Required
		NewSubscriber: &budgets.Subscriber{ // Required
			Address:          aws.String("GenericString"),    // Required
			SubscriptionType: aws.String("SubscriptionType"), // Required
		},
		Notification: &budgets.Notification{ // Required
			ComparisonOperator: aws.String("ComparisonOperator"), // Required
			NotificationType:   aws.String("NotificationType"),   // Required
			Threshold:          aws.Float64(1.0),                 // Required
		},
		OldSubscriber: &budgets.Subscriber{ // Required
			Address:          aws.String("GenericString"),    // Required
			SubscriptionType: aws.String("SubscriptionType"), // Required
		},
	}
	resp, err := svc.UpdateSubscriber(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
