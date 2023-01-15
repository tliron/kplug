package kplug

import (
	contextpkg "context"
	"net"

	api "github.com/tliron/kplug/kplug/grpc"
	"github.com/tliron/kutil/logging"
	"github.com/tliron/kutil/util"
	"google.golang.org/grpc"
)

//
// GRPCServer
//

type GRPCServer struct {
	api.UnimplementedControllerServer

	Protocol string
	Port     int
	Plugins  *Plugins
	Log      logging.Logger

	grpcServer *grpc.Server
}

func NewGRPCServer(protocol string, port int, plugins *Plugins, log logging.Logger) *GRPCServer {
	return &GRPCServer{
		Protocol: protocol,
		Port:     port,
		Plugins:  plugins,
		Log:      log,
	}
}

func (self *GRPCServer) Start() error {
	self.grpcServer = grpc.NewServer()
	api.RegisterControllerServer(self.grpcServer, self)

	if address, err := util.ToReachableIPAddress("0.0.0.0"); err == nil {
		if listener, err := net.Listen(self.Protocol, util.JoinIPAddressPort(address, self.Port)); err == nil {
			self.Log.Noticef("starting gRPC server on %s", listener.Addr().String())
			go func() {
				if err := self.grpcServer.Serve(listener); err != nil {
					self.Log.Errorf("%s", err.Error())
				}
			}()
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func (self *GRPCServer) Stop() {
	if self.grpcServer != nil {
		self.grpcServer.Stop()
	}
}

// api.ControllerServer interface
func (self *GRPCServer) PluginHeartbeat(context contextpkg.Context, pluginInformation *api.PluginInformation) (*api.PluginHeartbeatResponse, error) {
	self.Log.Infof("received heartbeat: %s", pluginInformation)
	if err := self.Plugins.Register(pluginInformation); err == nil {
		return &api.PluginHeartbeatResponse{Accepted: true}, nil
	} else {
		return &api.PluginHeartbeatResponse{Accepted: false, NotAcceptedReason: err.Error()}, nil
	}
}
