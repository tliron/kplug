package kplug

import (
	"errors"

	api "github.com/tliron/kplug/kplug/grpc"
	"github.com/tliron/kutil/ard"
	"google.golang.org/grpc"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

//
// Plugin
//

type Plugin struct {
	Information *api.PluginInformation
	Plugins     *Plugins
}

func (self *Plugins) NewPlugin(pluginInformation *api.PluginInformation) (*Plugin, error) {
	if pluginInformation.Address == "" {
		return nil, errors.New("no address")
	}
	if pluginInformation.Name == "" {
		return nil, errors.New("no name")
	}
	if pluginInformation.Api == "" {
		return nil, errors.New("no api")
	}

	return &Plugin{
		Information: pluginInformation,
		Plugins:     self,
	}, nil
}

func (self *Plugin) CreateWithReferences(base runtime.Object, references []core.ObjectReference, defaultNamespace string) (ard.StringMap, error) {
	if resourceReferences, err := NewResourceReferences(self.Plugins.Dynamic, references, defaultNamespace, self.Plugins.Log); err == nil {
		if baseStatus, referenceStatuses, err := self.Create(base, resourceReferences.Resources); err == nil {
			if err := resourceReferences.UpdateStatuses(referenceStatuses); err == nil {
				return baseStatus, nil
			} else {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (self *Plugin) Create(base runtime.Object, references map[string]*unstructured.Unstructured) (ard.StringMap, map[string]ard.StringMap, error) {
	if client, err := self.NewClient(); err == nil {
		if resources, err := toGrpcResources(base, references); err == nil {
			if resources_, err := client.Create(self.Plugins.Context, resources); err == nil {
				return fromGrpcResources(resources_)
			} else {
				return nil, nil, err
			}
		} else {
			return nil, nil, err
		}
	} else {
		return nil, nil, err
	}
}

func (self *Plugin) NewClient() (api.PluginClient, error) {
	self.Plugins.Log.Infof("calling plugin at %s", self.Information.Address)
	if connection, err := grpc.Dial(self.Information.Address, grpc.WithInsecure()); err == nil {
		return api.NewPluginClient(connection), nil
	} else {
		return nil, err
	}
}
