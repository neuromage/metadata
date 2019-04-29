from ._executor import Executor, Execution
from ._in_proc import InProcExecutor
import cloudpickle
import base64
import os

class DockerExecutor(Executor):
    def __init__(self, base_image = 'python:3.6', requires=[]):
        self.base_image = base_image
        self.requires = requires + ['kfmd']

    def execute(self, execution):
        def entry_point():
            InProcExecutor().execute(execution)
        entry_point_fn = base64.b64encode(cloudpickle.dumps(entry_point))
        python_shim = "import cloudpickle;import base64;cloudpickle.loads(base64.b64decode({}))()".format(
            entry_point_fn)
        python_shim = 'python3 -c "{}"'.format(python_shim)
        pip_install_code = ';'.join(['pip install {}'.format(require) for require in self.requires])
        bash_shim = '{};{}'.format(pip_install_code, python_shim)
        bash_shim = '/bin/bash -c "{}"'.format(bash_shim.replace('"', '\\"'))

        import docker
        client = docker.from_env()
        base_dir = execution.config.artifact_store.base_dir
        print(client.containers.run(self.base_image, bash_shim, remove = True,
            volumes = {
                base_dir: {
                    'bind': base_dir,
                    'mode': 'rw'
                }
            }).decode('utf-8'))
        execution_path = os.path.join(
            base_dir, 'executions', execution.id + '.json')
        if os.path.isfile(execution_path):
            with open(execution_path, 'r') as f:
                execution_json = f.read()
            execution.load_json(execution_json)
            return
