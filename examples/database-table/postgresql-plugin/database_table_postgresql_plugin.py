#!/usr/bin/env python3

import logging
from kplug.server import Server
from kplug.plugin import PluginInformation, Plugin
from kplug.heartbeat import Heartbeat
from implementation import Implementation


logging.basicConfig(encoding='utf-8', level=logging.DEBUG)

server = Server(Plugin(Implementation()), '0.0.0.0:50050')
server.start()

plugin_information = PluginInformation('database-table-postgresql-plugin:50050', 'PostgreSQL', '1.0.0', 'DatabaseTable', '1.0.0')
heartbeat = Heartbeat(plugin_information, 'database-table-operator:50050')
heartbeat.start()

server.block()
