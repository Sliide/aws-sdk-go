// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package redshift

import (
	"github.com/sliide/aws-sdk-go/aws"
	"github.com/sliide/aws-sdk-go/aws/client"
	"github.com/sliide/aws-sdk-go/aws/client/metadata"
	"github.com/sliide/aws-sdk-go/aws/request"
	"github.com/sliide/aws-sdk-go/aws/signer/v4"
	"github.com/sliide/aws-sdk-go/private/protocol/query"
)

// Overview
//
// This is an interface reference for Amazon Redshift. It contains documentation
// for one of the programming or command line interfaces you can use to manage
// Amazon Redshift clusters. Note that Amazon Redshift is asynchronous, which
// means that some interfaces may require techniques, such as polling or asynchronous
// callback handlers, to determine when a command has been applied. In this
// reference, the parameter descriptions indicate whether a change is applied
// immediately, on the next instance reboot, or during the next maintenance
// window. For a summary of the Amazon Redshift cluster management interfaces,
// go to Using the Amazon Redshift Management Interfaces (http://docs.aws.amazon.com/redshift/latest/mgmt/using-aws-sdk.html).
//
// Amazon Redshift manages all the work of setting up, operating, and scaling
// a data warehouse: provisioning capacity, monitoring and backing up the cluster,
// and applying patches and upgrades to the Amazon Redshift engine. You can
// focus on using your data to acquire new insights for your business and customers.
//
// If you are a first-time user of Amazon Redshift, we recommend that you begin
// by reading the Amazon Redshift Getting Started Guide (http://docs.aws.amazon.com/redshift/latest/gsg/getting-started.html).
//
// If you are a database developer, the Amazon Redshift Database Developer Guide
// (http://docs.aws.amazon.com/redshift/latest/dg/welcome.html) explains how
// to design, build, query, and maintain the databases that make up your data
// warehouse.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type Redshift struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "redshift"

// New creates a new instance of the Redshift client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a Redshift client from just a session.
//     svc := redshift.New(mySession)
//
//     // Create a Redshift client with additional configuration
//     svc := redshift.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *Redshift {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *Redshift {
	svc := &Redshift{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2012-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a Redshift operation and runs any
// custom request initialization.
func (c *Redshift) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
