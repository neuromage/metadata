
from .executors import InProcExecutor
from .artifacts import LocalStore

class Config:
    def __init__(self, executor, artifact_store):
        self.executor = executor
        self.artifact_store = artifact_store

    def merge(self, config):
        if not self.executor:
            self.executor = config.executor
        
        if not self.artifact_store:
            self.artifact_store = config.artifact_store

default_config = Config(InProcExecutor(), LocalStore('/tmp/kfmd/artifacts/'))

def config(executor = None, artifact_store = None):
    if executor:
        default_config.executor = executor

    if artifact_store:
        default_config.artifact_store = artifact_store