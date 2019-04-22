from kfmd import executable
import unittest

def add(a, b, outputs):
    outputs['sum'] = a + b

class ExecutableTest(unittest.TestCase):

    def test_executable_calls_fn(self):
        add_op = executable(add)(a = 1, b = 2)

        self.assertEqual(3, add_op.outputs['sum'])