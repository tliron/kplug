
import concurrent.futures, logging, grpc
from kplug import kplug_pb2_grpc


class Server:
  def __init__(self, plugin, address, max_workers=3):
    self.address = address
    self.server = grpc.server(concurrent.futures.ThreadPoolExecutor(max_workers=max_workers))
    self.server.add_insecure_port(address)
    kplug_pb2_grpc.add_PluginServicer_to_server(plugin, self.server)

  def start(self):
    logging.info('server.start at ' + self.address)
    self.server.start()

  def block(self):
    self.server.wait_for_termination()
