import inspect
import json

class Execution:
    def __init__(self, fn, inputs, outputs):
        self.fn = fn
        self.inputs = inputs
        self.outputs = outputs

    def to_json(self):
        return json.dumps({
            'inputs': self.inputs,
            'outputs': self.outputs,
            'fn': inspect.getsource(self.fn)
        })

    def load_json(self, data):
        states = json.loads(data)
        self.inputs = states['inputs']
        self.outputs = states['outputs']
        return self

class Executor:
    def execute(self, execution):
        pass