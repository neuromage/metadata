import inspect
import json
import os
import hashlib

class Execution:
    def __init__(self, fn, inputs, outputs, config = None):
        self.fn = fn
        self.inputs = inputs
        self.outputs = outputs
        self.outputs.execution = self
        self.outputs.artifact_store = config.artifact_store
        self.config = config
        self.id = hashlib.md5(self.to_json().encode('utf-8')).hexdigest()

    def to_json(self):
        return json.dumps({
            'inputs': self.inputs,
            'outputs': self.outputs.dict,
            'fn': inspect.getsource(self.fn)
        })

    def load_json(self, data):
        states = json.loads(data)
        self.inputs = states['inputs']
        self.outputs.dict = states['outputs']
        return self

class Outputs:
    def __init__(self, **kwargs):
        self._outputs = {} if kwargs is None else kwargs
        self.artifact_store = None
        self.execution = None

    def __getitem__(self, key):
        return self._outputs[key]

    def __setitem__(self, key, value):
        self._outputs[key] = value

    def artifact(self, key):
        artifact = {
            'uri': self._generate_artifact_path(key)
        }
        self._outputs[key] = artifact
        return artifact

    @property
    def dict(self):
        return self._outputs

    @dict.setter
    def dict(self, value):
        self._outputs = value

    def _generate_artifact_path(self, key):
        if self.artifact_store is None or self.execution is None:
            return None
        return os.path.join(self.artifact_store.base_dir, self.execution.id, key)

class Executor:
    def execute(self, execution):
        pass