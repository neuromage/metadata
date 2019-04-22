
from .executors import InProcExecutor
from .artifacts import LocalStore

class Config:
    def __init__(self, executor, artifact_store):
        self.executor = executor
        self.artifact_store = artifact_store

default_config = Config(InProcExecutor(), LocalStore('/tmp/kfmd/artifacts/'))

def config(executor):
    default_config.executor = executor