
from .executors import InProcExecutor

class Config:
    def __init__(self, executor):
        self.executor = executor

default_config = Config(InProcExecutor())

def config(executor):
    default_config.executor = executor