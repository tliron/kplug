
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
    base, extensions = from_grpc_resources(resources)
    base_status, extension_statuses = self.implementation.create(base, extensions)
    return to_grpc_resources(base_status, extension_statuses)

  def update(self, resources, context):
    logging.debug('plugin.update')
    base, extensions = from_grpc_resources(resources)
    base_status, extension_statuses = self.implementation.update(base, extensions)
    return to_grpc_resources(base_status, extension_statuses)

  def delete(self, resources, context):
    logging.debug('plugin.delete')
    base, extensions = from_grpc_resources(resources)
    self.implementation.delete(base, extensions)


def from_grpc_resources(resources):
  base = yaml.load(resources.base.yaml)
  extensions = {}
  for key, extension in resources.extensions.items():
    extensions[key] = yaml.load(extension.yaml)
  return base, extensions


def to_grpc_resources(base_status, extension_statuses):
  base_status = kplug_pb2.Resource(yaml=to_yaml(base_status))
  extension_statuses_ = {}
  if extension_statuses:
    for key, extension_status in extension_statuses.items():
      extension_statuses_[key] = kplug_pb2.Resource(yaml=to_yaml(extension_status))
  return kplug_pb2.Resources(base=base_status, extensions=extension_statuses_)
