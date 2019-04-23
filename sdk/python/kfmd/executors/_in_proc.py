from ._executor import Executor, Execution
import hashlib
import os

class InProcExecutor(Executor):
    def execute(self, execution):
        from .._config import default_config
        execution_path = os.path.join(
            default_config.artifact_store.base_dir, 'executions', execution.id + '.json')
        if os.path.isfile(execution_path):
            # Cached
            print('using cache')
            with open(execution_path, 'r') as f:
                execution_json = f.read()
            execution.load_json(execution_json)
            return
        args = { **execution.inputs, 'outputs': execution.outputs}
        execution.fn(**args)
        # Cache
        print('caching to ' + execution_path)
        os.makedirs(os.path.dirname(execution_path), exist_ok=True)
        with open(execution_path, 'w') as f:
            f.write(execution.to_json())
