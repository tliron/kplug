package client

import (
	contextpkg "context"

	"github.com/tliron/commonlog"
	myorgpkg "github.com/tliron/kplug/examples/database-table/operator/apis/clientset/versioned"
)

//
// Client
//

type Client struct {
	MyOrg myorgpkg.Interface

	Namespace string

	Context contextpkg.Context
	Log     commonlog.Logger
}

func NewClient(myOrg myorgpkg.Interface, context contextpkg.Context, namespace string, logName string) *Client {
	return &Client{
		MyOrg:     myOrg,
		Namespace: namespace,
		Context:   context,
		Log:       commonlog.GetLogger(logName),
	}
}
