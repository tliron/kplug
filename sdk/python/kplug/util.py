
from io import StringIO
from ruamel.yaml import YAML


yaml=YAML(typ='rt')


def to_yaml(o):
  s = StringIO()
  yaml.dump(o, s)
  return s.getvalue()
