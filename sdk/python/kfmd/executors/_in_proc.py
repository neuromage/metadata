
from ._executor import Executor

class InProcExecutor(Executor):
    def execute(self, execution):
        args = { **execution.inputs, 'outputs': execution.outputs}
        execution.fn(**args)
