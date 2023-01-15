package client

import (
	resources "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (self *Client) GetDatabaseTable(namespace string, databaseTableName string) (*resources.DatabaseTable, error) {
	if namespace == "" {
		namespace = self.Namespace
	}

	if databaseTable, err := self.MyOrg.MyorgV1alpha1().DatabaseTables(namespace).Get(self.Context, databaseTableName, meta.GetOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if databaseTable.Kind == "" {
			databaseTable = databaseTable.DeepCopy()
			databaseTable.APIVersion, databaseTable.Kind = resources.DatabaseTableGVK.ToAPIVersionAndKind()
		}
		return databaseTable, nil
	} else {
		return nil, err
	}
}

func (self *Client) ListDatabaseTables() (*resources.DatabaseTableList, error) {
	return self.MyOrg.MyorgV1alpha1().DatabaseTables(self.Namespace).List(self.Context, meta.ListOptions{})
}

func (self *Client) UpdateDatabaseTableSpec(databaseTable *resources.DatabaseTable) (*resources.DatabaseTable, error) {
	if databaseTable_, err := self.MyOrg.MyorgV1alpha1().DatabaseTables(databaseTable.Namespace).Update(self.Context, databaseTable, meta.UpdateOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if databaseTable_.Kind == "" {
			databaseTable_ = databaseTable_.DeepCopy()
			databaseTable_.APIVersion, databaseTable_.Kind = resources.DatabaseTableGVK.ToAPIVersionAndKind()
		}
		return databaseTable_, nil
	} else {
		return databaseTable, err
	}
}

func (self *Client) UpdateDatabaseTableStatus(databaseTable *resources.DatabaseTable) (*resources.DatabaseTable, error) {
	if databaseTable_, err := self.MyOrg.MyorgV1alpha1().DatabaseTables(databaseTable.Namespace).UpdateStatus(self.Context, databaseTable, meta.UpdateOptions{}); err == nil {
		// When retrieved from cache the GVK may be empty
		if databaseTable_.Kind == "" {
			databaseTable_ = databaseTable_.DeepCopy()
			databaseTable_.APIVersion, databaseTable_.Kind = resources.DatabaseTableGVK.ToAPIVersionAndKind()
		}
		return databaseTable_, nil
	} else {
		return databaseTable, err
	}
}

func (self *Client) DeleteDatabaseTable(namespace string, databaseTableName string) error {
	if namespace == "" {
		namespace = self.Namespace
	}

	return self.MyOrg.MyorgV1alpha1().DatabaseTables(namespace).Delete(self.Context, databaseTableName, meta.DeleteOptions{})
}
