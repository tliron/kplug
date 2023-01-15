package client

import (
	contextpkg "context"

	myorgpkg "github.com/tliron/kplug/examples/database-table/operator/apis/clientset/versioned"
	"github.com/tliron/kutil/logging"
)

//
// Client
//

type Client struct {
	MyOrg myorgpkg.Interface

	Namespace string

	Context contextpkg.Context
	Log     logging.Logger
}

func NewClient(myOrg myorgpkg.Interface, context contextpkg.Context, namespace string, logName string) *Client {
	return &Client{
		MyOrg:     myOrg,
		Namespace: namespace,
		Context:   context,
		Log:       logging.GetLogger(logName),
	}
}
