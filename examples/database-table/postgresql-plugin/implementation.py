
import logging
from kplug.util import to_yaml


class Implementation:
  def create(self, base, extensions):
    #logging.debug('create:\nbase:\n' + to_yaml(base) + '\nextensions:\n' + to_yaml(extensions))

    # Extract properties we need from the base resource
    columns = base['spec']['columns']

    # Extract properties we need from the extensions that we support
    partition_by = ''
    postgresql_extension_id = ''
    for id_, extension in extensions.items():
      # Check for our supported extension (there could be more than one, the semantics are up to us)
      if (extension['apiVersion'] == 'myorg.org/v1alpha1') and (extension['kind'] == 'DatabaseTablePostgreSqlExtension'):
        postgresql_extension_id = id_
        partition_by = extension['spec']['partitionBy']

    logging.debug('create\ncolumns:\n' + to_yaml(columns) + '\npartition_by: ' + partition_by)

    #
    # database table creation logic goes here
    #

    # We return status updates: first for the base resource and then for extensions by their ID
    return {'implementation': 'PostgreSQL'}, {
             postgresql_extension_id: {'partitioned': True}
           }

  def update(self, base, extensions):
    return {}, {}

  def delete(self, base, extensions):
    pass
