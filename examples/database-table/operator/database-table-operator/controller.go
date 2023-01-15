package main

import (
	"fmt"
	"net/http"

	"github.com/heptiolabs/healthcheck"
	myorgpkg "github.com/tliron/kplug/examples/database-table/operator/apis/clientset/versioned"
	controllerpkg "github.com/tliron/kplug/examples/database-table/operator/controller"
	"github.com/tliron/kutil/kubernetes"
	"github.com/tliron/kutil/util"
	versionpkg "github.com/tliron/kutil/version"
	"k8s.io/client-go/dynamic"
	kubernetespkg "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Controller() {
	if version {
		versionpkg.Print()
		util.Exit(0)
		return
	}

	log.Noticef("%s version=%s revision=%s", toolName, versionpkg.GitVersion, versionpkg.GitRevision)

	// Config

	config, err := clientcmd.BuildConfigFromFlags(masterUrl, kubeconfigPath)
	util.FailOnError(err)

	config.QPS = 100
	config.Burst = 10

	if namespace == "" {
		if namespace_, ok := kubernetes.GetConfiguredNamespace(kubeconfigPath, kubeconfigContext); ok {
			namespace = namespace_
		}
		if namespace == "" {
			namespace = kubernetes.GetServiceAccountNamespace()
		}
		if namespace == "" {
			util.Fail("could not discover namespace and namespace not provided")
		}
	}

	// Clients

	dynamicClient, err := dynamic.NewForConfig(config)
	util.FailOnError(err)

	kubernetesClient, err := kubernetespkg.NewForConfig(config)
	util.FailOnError(err)

	myOrgClient, err := myorgpkg.NewForConfig(config)
	util.FailOnError(err)

	// Controller

	controller := controllerpkg.NewController(
		context,
		toolName,
		namespace,
		dynamicClient,
		kubernetesClient,
		myOrgClient,
		resyncPeriod,
		util.SetupSignalHandler(),
		grpcProtocol,
		int(grpcPort),
	)

	// Run

	err = controller.Run(concurrency, func() {
		log.Notice("starting health monitor")
		health := healthcheck.NewHandler()
		err := http.ListenAndServe(fmt.Sprintf(":%d", healthPort), health)
		util.FailOnError(err)
	})
	util.FailOnError(err)
}
