class Execution:
    def __init__(self, fn, inputs):
        self.fn = fn
        self.inputs = inputs
        self.outputs = {}

class Executor:
    def execute(self, execution):
        pass