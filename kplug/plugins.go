package kplug

import (
	contextpkg "context"
	"sync"

	api "github.com/tliron/kplug/kplug/grpc"
	"github.com/tliron/kutil/kubernetes"
	"github.com/tliron/kutil/logging"
)

//
// Plugins
//

type Plugins struct {
	Plugins map[PluginKey]*Plugin
	Dynamic *kubernetes.Dynamic
	Context contextpkg.Context
	Log     logging.Logger

	lock sync.Mutex
}

func NewPlugins(dynamic *kubernetes.Dynamic, context contextpkg.Context, log logging.Logger) *Plugins {
	return &Plugins{
		Plugins: make(map[PluginKey]*Plugin),
		Dynamic: dynamic,
		Context: context,
		Log:     log,
	}
}

func (self *Plugins) Get() *Plugin {
	self.lock.Lock()
	defer self.lock.Unlock()

	for _, plugin := range self.Plugins {
		return plugin
	}

	return nil
}

func (self *Plugins) Register(pluginInformation *api.PluginInformation) error {
	if plugin, err := self.NewPlugin(pluginInformation); err == nil {
		self.lock.Lock()
		defer self.lock.Unlock()

		self.Plugins[NewPluginKey(pluginInformation)] = plugin
		return nil
	} else {
		return err
	}
}

//
// PluginKey
//

type PluginKey struct {
	Name    string
	Version string
}

func NewPluginKey(pluginInformation *api.PluginInformation) PluginKey {
	return PluginKey{pluginInformation.Name, pluginInformation.Version}
}
