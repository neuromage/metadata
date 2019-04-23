from ._config import default_config
from .executors import Execution, Outputs

def executable(executable_fn):
    def execution_factory(**kwargs):
        outputs = Outputs()
        execution = Execution(executable_fn, kwargs, outputs, default_config)
        default_config.executor.execute(execution)
        return execution

    return execution_factory

