
import logging
from kplug import kplug_pb2_grpc, kplug_pb2
from kplug.util import yaml, to_yaml


class PluginInformation:
  def __init__(self, address, name, version, api, api_version):
    self.address = address
    self.name = name
    self.version = version
    self.api = api
    self.api_version = api_version


class Plugin(kplug_pb2_grpc.PluginServicer):
  def __init__(self, implementation):
    super().__init__()
    self.implementation = implementation

  def create(self, resources, context):
    logging.debug('plugin.create')
    base, references = from_grpc_resources(resources)
    base_status, reference_statuses = self.implementation.create(base, references)
    return to_grpc_resources(base_status, reference_statuses)

  def update(self, resources, context):
    logging.debug('plugin.update')
    base, references = from_grpc_resources(resources)
    base_status, reference_statuses = self.implementation.update(base, references)
    return to_grpc_resources(base_status, reference_statuses)

  def delete(self, resources, context):
    logging.debug('plugin.delete')
    base, references = from_grpc_resources(resources)
    self.implementation.delete(base, references)


def from_grpc_resources(resources):
  base = yaml.load(resources.base.yaml)
  references = {}
  for key, reference in resources.references.items():
    references[key] = yaml.load(reference.yaml)
  return base, references


def to_grpc_resources(base_status, reference_statuses):
  base_status = kplug_pb2.Resource(yaml=to_yaml(base_status))
  reference_statuses_ = {}
  if reference_statuses:
    for key, reference_status in reference_statuses.items():
      reference_statuses_[key] = kplug_pb2.Resource(yaml=to_yaml(reference_status))
  return kplug_pb2.Resources(base=base_status, references=reference_statuses_)
