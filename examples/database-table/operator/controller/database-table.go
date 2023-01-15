package controller

import (
	resources "github.com/tliron/kplug/examples/database-table/operator/resources/myorg.org/v1alpha1"
	"github.com/tliron/kutil/ard"
)

func (self *Controller) processDatabaseTable(databaseTable *resources.DatabaseTable) (bool, error) {
	if plugin := self.Plugins.Get(); plugin != nil {
		if status, err := plugin.CreateWithExtensions(databaseTable, databaseTable.Spec.Extensions, databaseTable.Namespace); err == nil {
			// Log
			//baseStatus_, _ := transcribe.EncodeYAML(baseStatus, "  ", false)
			//extensionStatuses_, _ := transcribe.EncodeYAML(extensionStatuses, "  ", false)
			//self.Log.Infof("%s.create\n%s\n%s", plugin.Information.Name, baseStatus_, extensionStatuses_)

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
