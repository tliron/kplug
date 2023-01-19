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

func toGrpcResources(base runtime.Object, references map[string]*unstructured.Unstructured) (*api.Resources, error) {
	base_, err := toYaml(base)
	if err != nil {
		return nil, err
	}
	base__ := api.Resource{Yaml: base_}

	references_ := make(map[string]*api.Resource)
	for key, resource := range references {
		if resource_, err := toYaml(resource); err == nil {
			references_[key] = &api.Resource{Yaml: resource_}
		} else {
			return nil, err
		}
	}

	return &api.Resources{Base: &base__, References: references_}, nil
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

	referenceStatuses := make(map[string]ard.StringMap)
	for key, referenceStatus := range resources.References {
		if referenceStatus_, _, err := ard.DecodeYAML(referenceStatus.Yaml, false); err == nil {
			referenceStatus_, _ = ard.NormalizeStringMaps(referenceStatus_)
			if referenceStatuses[key], ok = referenceStatus_.(ard.StringMap); !ok {
				return nil, nil, fmt.Errorf("reference status not a map: %T", referenceStatus_)
			}
		} else {
			return nil, nil, err
		}
	}

	return baseStatus_, referenceStatuses, nil
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
