from ._config import default_config
from .executors import Execution

def executable(executable_fn):
    def execution_factory(**kwargs):
        execution = Execution(executable_fn, kwargs, outputs={})
        default_config.executor.execute(execution)
        return execution

    return execution_factory

