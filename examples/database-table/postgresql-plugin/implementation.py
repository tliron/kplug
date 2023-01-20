
import logging
from kplug.util import is_reference_gvk, to_yaml


class Implementation:
  def create(self, base, references):
    #logging.debug('create:\nbase:\n' + to_yaml(base) + '\nreferences:\n' + to_yaml(references))

    # Extract properties we need from the base resource
    server_name = base['spec']['serverName']
    columns = base['spec']['columns']

    # Extract properties we need from the references that we support
    partition_by = ''
    postgresql_reference_id = ''
    for id_, reference in references.items():
      # Check for our supported reference (there could be more than one, the semantics are up to us)
      if is_reference_gvk(reference, 'myorg.org', 'v1alpha1', 'PostgreSqlDatabaseTable'):
        postgresql_reference_id = id_
        partition_by = reference['spec']['partitionBy']

    logging.debug('create' +
                  '\nserver_name: ' + server_name +
                  '\npartition_by: ' + partition_by +
                  '\ncolumns:\n' + to_yaml(columns))

    #
    # database table creation logic goes here
    #

    # We return status updates: first for the base resource and then for references by their ID
    return {'implementation': 'PostgreSQL'}, {
             postgresql_reference_id:
               {'partitioned': True}}

  def update(self, base, references):
    return {}, {}

  def delete(self, base, references):
    pass
