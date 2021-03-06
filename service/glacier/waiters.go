// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package glacier

import (
	"github.com/sliide/aws-sdk-go/private/waiter"
)

// WaitUntilVaultExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *Glacier) WaitUntilVaultExists(input *DescribeVaultInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVault",
		Delay:       3,
		MaxAttempts: 15,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "success",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "retry",
				Matcher:  "error",
				Argument: "",
				Expected: "ResourceNotFoundException",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}

// WaitUntilVaultNotExists uses the Amazon Glacier API operation
// DescribeVault to wait for a condition to be met before returning.
// If the condition is not meet within the max attempt window an error will
// be returned.
func (c *Glacier) WaitUntilVaultNotExists(input *DescribeVaultInput) error {
	waiterCfg := waiter.Config{
		Operation:   "DescribeVault",
		Delay:       3,
		MaxAttempts: 15,
		Acceptors: []waiter.WaitAcceptor{
			{
				State:    "retry",
				Matcher:  "status",
				Argument: "",
				Expected: 200,
			},
			{
				State:    "success",
				Matcher:  "error",
				Argument: "",
				Expected: "ResourceNotFoundException",
			},
		},
	}

	w := waiter.Waiter{
		Client: c,
		Input:  input,
		Config: waiterCfg,
	}
	return w.Wait()
}
