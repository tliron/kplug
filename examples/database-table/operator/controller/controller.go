package controller

import (
	contextpkg "context"
	"fmt"
	"time"

	"github.com/tliron/commonlog"
	myorgpkg "github.com/tliron/kplug/examples/database-table/operator/apis/clientset/versioned"
	myorginformers "github.com/tliron/kplug/examples/database-table/operator/apis/informers/externalversions"
	myorglisters "github.com/tliron/kplug/examples/database-table/operator/apis/listers/myorg.org/v1alpha1"
	clientpkg "github.com/tliron/kplug/examples/database-table/operator/client"
	myorgresources "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	"github.com/tliron/kplug/kplug"
	kplugpkg "github.com/tliron/kplug/kplug"
	kubernetesutil "github.com/tliron/kutil/kubernetes"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	dynamicpkg "k8s.io/client-go/dynamic"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
)

//
// Controller
//

type Controller struct {
	Dynamic      *kubernetesutil.Dynamic
	Kubernetes   kubernetes.Interface
	MyOrg        myorgpkg.Interface
	Client       *clientpkg.Client
	StopChannel  <-chan struct{}
	GRPCProtocol string
	GRPCPort     int

	Processors *kubernetesutil.Processors
	Events     record.EventRecorder

	KubernetesInformerFactory    informers.SharedInformerFactory
	DatabaseTableInformerFactory myorginformers.SharedInformerFactory

	DatabaseTables myorglisters.DatabaseTableLister

	Context contextpkg.Context
	Log     commonlog.Logger

	Plugins *kplugpkg.Plugins
}

func NewController(context contextpkg.Context, toolName string, namespace string, dynamic dynamicpkg.Interface, kubernetes kubernetes.Interface, kplug myorgpkg.Interface, informerResyncPeriod time.Duration, stopChannel <-chan struct{}, grpcProtocol string, grpcPort int) *Controller {
	log := commonlog.GetLoggerf("%s.controller", toolName)

	self := Controller{
		Dynamic:      kubernetesutil.NewDynamic(toolName, dynamic, kubernetes.Discovery(), namespace, context),
		Kubernetes:   kubernetes,
		MyOrg:        kplug,
		StopChannel:  stopChannel,
		GRPCPort:     grpcPort,
		GRPCProtocol: grpcProtocol,
		Processors:   kubernetesutil.NewProcessors(toolName),
		Events:       kubernetesutil.CreateEventRecorder(kubernetes, "DatabaseTable", log),
		Context:      context,
		Log:          log,
	}

	self.Plugins = kplugpkg.NewPlugins(self.Dynamic, context, log)

	self.Client = clientpkg.NewClient(
		kplug,
		context,
		namespace,
		fmt.Sprintf("%s.client", toolName),
	)

	self.KubernetesInformerFactory = informers.NewSharedInformerFactoryWithOptions(kubernetes, informerResyncPeriod, informers.WithNamespace(namespace))
	self.DatabaseTableInformerFactory = myorginformers.NewSharedInformerFactoryWithOptions(kplug, informerResyncPeriod, myorginformers.WithNamespace(namespace))

	// Informers
	databaseTableInformer := self.DatabaseTableInformerFactory.Myorg().V1alpha1().DatabaseTables()

	// Listers
	self.DatabaseTables = databaseTableInformer.Lister()

	// Processors

	processorPeriod := 5 * time.Second

	self.Processors.Add(myorgresources.DatabaseTableGVK, kubernetesutil.NewProcessor(
		toolName,
		"database-tables",
		databaseTableInformer.Informer(),
		processorPeriod,
		func(name string, namespace string) (any, error) {
			return self.Client.GetDatabaseTable(namespace, name)
		},
		func(object any) (bool, error) {
			return self.processDatabaseTable(object.(*myorgresources.DatabaseTable))
		},
	))

	return &self
}

func (self *Controller) Run(concurrency uint, startup func()) error {
	defer utilruntime.HandleCrash()

	grpc := kplug.NewGRPCServer(self.GRPCProtocol, self.GRPCPort, self.Plugins, self.Log)
	if err := grpc.Start(); err == nil {
		defer grpc.Stop()
	} else {
		return err
	}

	self.Log.Info("starting informer factories")
	self.KubernetesInformerFactory.Start(self.StopChannel)
	self.DatabaseTableInformerFactory.Start(self.StopChannel)

	self.Log.Info("waiting for processor informer caches to sync")
	utilruntime.HandleError(self.Processors.WaitForCacheSync(self.StopChannel))

	self.Log.Infof("starting processors (concurrency=%d)", concurrency)
	self.Processors.Start(concurrency, self.StopChannel)
	defer self.Processors.ShutDown()

	if startup != nil {
		go startup()
	}

	<-self.StopChannel

	self.Log.Info("shutting down")

	return nil
}
