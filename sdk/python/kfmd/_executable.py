from ._config import default_config
from .executors import Execution, Outputs

def executable(executable_fn, config = None):
    if not config:
        config = default_config
    else:
        config.merge(default_config)
    def execution_factory(**kwargs):
        outputs = Outputs()
        execution = Execution(executable_fn, kwargs, outputs, config)
        config.executor.execute(execution)
        return execution

    return execution_factory