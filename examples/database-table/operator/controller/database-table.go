package controller

import (
	resources "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	"github.com/tliron/kutil/ard"
)

func (self *Controller) processDatabaseTable(databaseTable *resources.DatabaseTable) (bool, error) {
	if plugin := self.Plugins.Get(); plugin != nil {
		if status, err := plugin.CreateWithReferences(databaseTable, databaseTable.Spec.References, databaseTable.Namespace); err == nil {
			if len(status) > 0 {
				if implementation, ok := ard.NewNode(status).Get("implementation").String(); ok {
					databaseTable.Status.Implementation = implementation
				}

				self.Log.Infof("updating status for database table: %+v", databaseTable.Status)
				if _, err := self.Client.UpdateDatabaseTableStatus(databaseTable); err != nil {
					return false, err
				}
			}
		} else {
			return false, err
		}

		return true, nil
	} else {
		return false, nil
	}
}
