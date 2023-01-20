package kplug

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tliron/kutil/ard"
	"github.com/tliron/kutil/kubernetes"
	"github.com/tliron/kutil/logging"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

//
// ResourceReferences
//

type ResourceReferences struct {
	Resources map[string]*unstructured.Unstructured
	Dynamic   *kubernetes.Dynamic
	Log       logging.Logger
}

func NewResourceReferences(dynamic *kubernetes.Dynamic, objectReferences []core.ObjectReference, defaultNamespace string, log logging.Logger) (*ResourceReferences, error) {
	self := ResourceReferences{
		Resources: make(map[string]*unstructured.Unstructured),
		Dynamic:   dynamic,
		Log:       log,
	}

	for _, objectReference := range objectReferences {
		namespace := objectReference.Namespace
		if namespace == "" {
			namespace = defaultNamespace
		}

		gvk := objectReference.GroupVersionKind()
		if resource, err := dynamic.GetResource(gvk, objectReference.Name, namespace); err == nil {
			key := ToResourceReferenceKey(gvk, namespace, objectReference.Name)
			self.Resources[key] = resource
		} else {
			return nil, err
		}
	}

	return &self, nil
}

func (self *ResourceReferences) SetController(controller meta.Object) error {
	for _, resource := range self.Resources {
		if err := kubernetes.SetControllerOfUnstructured(resource, controller); err == nil {
			if _, err := self.Dynamic.UpdateResource(resource); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func (self *ResourceReferences) UpdateStatuses(statuses map[string]ard.StringMap) error {
	for key, status := range statuses {
		if err := self.UpdateStatus(key, status); err != nil {
			return err
		}
	}

	return nil
}

func (self *ResourceReferences) UpdateStatus(key string, status ard.StringMap) error {
	if resource, ok := self.Resources[key]; ok {
		if len(status) > 0 {
			var status_ map[string]any
			if status__, ok := resource.Object["status"]; ok {
				status_, _ = status__.(map[string]any)
			}
			if status_ == nil {
				status_ = make(map[string]any)
				resource.Object["status"] = status_
			}

			ard.MergeStringMaps(status_, status, false)

			self.Log.Infof("updating status for %s: %+v", key, status_)
			if _, err := self.Dynamic.UpdateResourceStatus(resource); err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("unsupported resource reference status key: %s", key)
	}

	return nil
}

// Utils

func ToResourceReferenceKey(gvk schema.GroupVersionKind, namespace string, name string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", gvk.Group, gvk.Version, gvk.Kind, namespace, name)
}

func FromResourceReferenceKey(key string) (schema.GroupVersionKind, string, string, error) {
	s := strings.Split(key, "/")

	if len(s) != 5 {
		return schema.GroupVersionKind{}, "", "", errors.New("malformed resource reference key")
	}

	group := s[0]
	version := s[1]
	kind := s[2]
	namespace := s[3]
	name := s[4]

	return schema.GroupVersionKind{Group: group, Version: version, Kind: kind}, namespace, name, nil
}
