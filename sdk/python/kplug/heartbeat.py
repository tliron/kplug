
import grpc, threading, time, logging
from kplug import kplug_pb2_grpc, kplug_pb2


class Heartbeat:
  def __init__(self, plugin_information, controller_address, interval=5, timeout=3):
    self.plugin_information = plugin_information
    self.controller_address = controller_address
    self.interval = interval # seconds
    self.timeout = timeout # seconds

  def start(self):
    thread = threading.Thread(target=self.thread, daemon=True)
    logging.info("heartbeat.start to " + self.controller_address)
    thread.start()

  def thread(self):
    while True:
        self.call()
        time.sleep(self.interval)

  def call(self):
    try:
      with grpc.insecure_channel(self.controller_address) as channel:
        stub = kplug_pb2_grpc.ControllerStub(channel)
        logging.debug('heartbeat.call')
        plugin_information = kplug_pb2.PluginInformation(address=self.plugin_information.address, name=self.plugin_information.name, version=self.plugin_information.version, api=self.plugin_information.api, apiVersion=self.plugin_information.api_version)
        response = stub.pluginHeartbeat(plugin_information, timeout=self.timeout)
      logging.debug('heartbeat.call accepted: ' + str(response.accepted) + (', reason: ' + response.notAcceptedReason if response.notAcceptedReason else ''))
    except Exception:
        logging.exception('heartbeat.call')
