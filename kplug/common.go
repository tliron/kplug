package kplug

import (
	"fmt"
	"strings"

	api "github.com/tliron/kplug/kplug/grpc"
	"github.com/tliron/kutil/ard"
	"github.com/tliron/kutil/transcribe"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
)

func toGrpcResources(base runtime.Object, extensions map[string]*unstructured.Unstructured) (*api.Resources, error) {
	base_, err := toYaml(base)
	if err != nil {
		return nil, err
	}
	base__ := api.Resource{Yaml: base_}

	extensions_ := make(map[string]*api.Resource)
	for key, resource := range extensions {
		if resource_, err := toYaml(resource); err == nil {
			extensions_[key] = &api.Resource{Yaml: resource_}
		} else {
			return nil, err
		}
	}

	return &api.Resources{Base: &base__, Extensions: extensions_}, nil
}

func fromGrpcResources(resources *api.Resources) (ard.StringMap, map[string]ard.StringMap, error) {
	baseStatus, _, err := ard.DecodeYAML(resources.Base.Yaml, false)
	if err != nil {
		return nil, nil, err
	}
	baseStatus, _ = ard.NormalizeStringMaps(baseStatus)
	var baseStatus_ ard.StringMap
	var ok bool
	if baseStatus_, ok = baseStatus.(ard.StringMap); !ok {
		return nil, nil, fmt.Errorf("base status not a map: %T", baseStatus)
	}

	extensionsStatuses := make(map[string]ard.StringMap)
	for key, extensionStatus := range resources.Extensions {
		if extensionStatus_, _, err := ard.DecodeYAML(extensionStatus.Yaml, false); err == nil {
			extensionStatus_, _ = ard.NormalizeStringMaps(extensionStatus_)
			if extensionsStatuses[key], ok = extensionStatus_.(ard.StringMap); !ok {
				return nil, nil, fmt.Errorf("extension status not a map: %T", extensionStatus_)
			}
		} else {
			return nil, nil, err
		}
	}

	return baseStatus_, extensionsStatuses, nil
}

func toYaml(object any) (string, error) {
	if object_, ok := object.(runtime.Object); ok {
		serializer := json.NewYAMLSerializer(json.DefaultMetaFactory, nil, nil)
		var builder strings.Builder
		if err := serializer.Encode(object_, &builder); err == nil {
			return builder.String(), nil
		} else {
			return "", err
		}
	} else {
		return transcribe.EncodeYAML(object, "  ", false)
	}
}
