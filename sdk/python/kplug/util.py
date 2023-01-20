
from io import StringIO
from ruamel.yaml import YAML


yaml=YAML()


def is_reference_gvk(reference, group, version, kind):
  return (reference['apiVersion'] == group + '/' + version) and (reference['kind'] == kind)


def to_yaml(o):
  s = StringIO()
  yaml.dump(o, s)
  return s.getvalue()
